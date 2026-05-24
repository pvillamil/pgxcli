package renderer

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/balajz/pgxcli/internal/database"
)

func GetColumnStrings(rows database.Rows, showMoreChars bool) []string {
	srcCols := rows.Columns()
	cols := make([]string, len(srcCols))
	for i, c := range srcCols {
		cols[i] = FormatVal(c, showMoreChars, showMoreChars)
	}
	return cols
}

func getAllRowStrings(rows database.Rows, showMoreChars bool) ([][]string, error) {
	var allRows [][]string

	for {
		rowStrings, err := getNextRowStrings(rows, showMoreChars)
		if err != nil {
			return nil, err
		}
		if rowStrings == nil {
			break
		}
		allRows = append(allRows, rowStrings)
	}

	return allRows, nil
}

func getNextRowStrings(rows database.Rows, showMoreChars bool) ([]string, error) {
	cols := rows.Columns()
	var vals []driver.Value
	if len(cols) > 0 {
		vals = make([]driver.Value, len(cols))
	}

	err := rows.Next(vals)
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// These are expected to already be formatted strings, but if not, try to
	// print them anyway.
	rowStrings := make([]string, len(cols))
	for i, v := range vals {
		rowStrings[i] = FormatVal(v, showMoreChars, showMoreChars)
	}
	return rowStrings, nil
}

func isNotPrintableASCII(r rune) bool { return r < 0x20 || r > 0x7e || r == '"' || r == '\\' }
func isNotGraphicUnicode(r rune) bool { return !unicode.IsGraphic(r) }
func isNotGraphicUnicodeOrTabOrNewline(r rune) bool {
	return r != '\t' && r != '\n' && !unicode.IsGraphic(r)
}

// FormatVal formats a value retrieved by a SQL driver into a string
// suitable for displaying to the user.
func FormatVal(val driver.Value, showPrintableUnicode bool, showNewLinesAndTabs bool) string {
	switch t := val.(type) {
	case nil:
		return "NULL"

	case string:
		if showPrintableUnicode {
			pred := isNotGraphicUnicode
			if showNewLinesAndTabs {
				pred = isNotGraphicUnicodeOrTabOrNewline
			}
			if utf8.ValidString(t) && strings.IndexFunc(t, pred) == -1 {
				return t
			}
		} else {
			if strings.IndexFunc(t, isNotPrintableASCII) == -1 {
				return t
			}
		}
	}

	return fmt.Sprint(val)
}
