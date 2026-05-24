package renderer

import (
	"io"

	"github.com/balajz/pgxcli/internal/database"
)

// RowStrIter is an iterator interface for the PrintQueryOutput function. It is
// used so that results can be streamed to the row formatters as they arrive
// to the CLI.
type RowStrIter interface {
	Next() (row []string, err error)
	ToSlice() (allRows [][]string, err error)
}

// rowSliceIter is an implementation of the rowStrIter interface and it is used
// to wrap a slice of rows that have already been completely buffered into
// memory.
type rowSliceIter struct {
	allRows [][]string
	index   int
}

// Next returns next row of rowSliceIter.
func (iter *rowSliceIter) Next() (row []string, err error) {
	if iter.index >= len(iter.allRows) {
		return nil, io.EOF
	}
	row = iter.allRows[iter.index]
	iter.index = iter.index + 1
	return row, nil
}

// ToSlice returns all rows of rowSliceIter.
func (iter *rowSliceIter) ToSlice() ([][]string, error) {
	return iter.allRows, nil
}

// NewRowSliceIter is an implementation of the rowStrIter interface and it is
// used when the rows have not been buffered into memory yet and we want to
// stream them to the row formatters as they arrive over the network.
func NewRowSliceIter(allRows [][]string) RowStrIter {
	return &rowSliceIter{
		allRows: allRows,
		index:   0,
	}
}

type RowIter struct {
	rows          database.Rows
	showMoreChars bool
}

func (iter *RowIter) Next() (row []string, err error) {
	nextRowString, err := getNextRowStrings(iter.rows, iter.showMoreChars)
	if err != nil {
		return nil, err
	}
	if nextRowString == nil {
		return nil, io.EOF
	}
	return nextRowString, nil
}

func (iter *RowIter) ToSlice() ([][]string, error) {
	return getAllRowStrings(iter.rows, iter.showMoreChars)
}

func NewRowIter(rows database.Rows, showMoreChars bool) *RowIter {
	return &RowIter{
		rows:          rows,
		showMoreChars: showMoreChars,
	}
}
