package renderer

import (
	"strings"
	"testing"

	"github.com/balajz/pgxcli/internal/config"
	"github.com/stretchr/testify/require"
)

func TestTableFormattingAdvanced(t *testing.T) {
	// These test cases are adapted from the CockroachDB reference files
	// in cli/clisqlexec/ to ensure robust multi-line and special character handling.

	testCases := []struct {
		name         string
		cols         []string
		rows         [][]string
		caption      string
		wantContains []string
	}{
		{
			name: "special characters and quotes",
			cols: []string{`f"oo`, `f'oo`, `f\oo`, `κόσμε`},
			rows: [][]string{
				{`0`, `0`, `0`, `0`},
			},
			wantContains: []string{
				`F " OO`, `F ' OO`, `F \ OO`, `ΚΌΣΜΕ`,
			},
		},
		{
			name: "multi-line content wrapping",
			cols: []string{"id", "description"},
			rows: [][]string{
				{"1", "short\nvery very long\nnot much"},
				{"2", "just fine"},
			},
			wantContains: []string{
				"ID", "DESCRIPTION",
				"short", "very very long", "not much",
				"just fine",
			},
		},
		{
			name: "unicode and symbols",
			cols: []string{"a|b", "܈85"},
			rows: [][]string{
				{"val1", "val2"},
			},
			wantContains: []string{
				"A | B", "܈ 85", "val1", "val2",
			},
		},
		{
			name: "empty results with header",
			cols: []string{"col1", "col2"},
			rows: [][]string{},
			wantContains: []string{
				"COL 1", "COL 2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out strings.Builder
			rowIter := NewRowSliceIter(tc.rows)

			// Use default config for styling
			cfg := &config.Config{}

			err := TableRender(tc.cols, rowIter, tc.caption, &out, &out, cfg)
			require.NoError(t, err)

			rendered := out.String()
			for _, want := range tc.wantContains {
				require.Contains(t, rendered, want, "Output should contain: %s", want)
			}
		})
	}
}

func TestTableRenderAlignment(t *testing.T) {
	// Verify that the table structure is maintained with different lengths
	cols := []string{"name", "value"}
	rows := [][]string{
		{"Alice", "100"},
		{"Bob", "1000000"},
	}

	var out strings.Builder
	rowIter := NewRowSliceIter(rows)
	err := TableRender(cols, rowIter, "", &out, &out, &config.Config{})
	require.NoError(t, err)

	rendered := out.String()
	// Basic check for table structure (borders/separators)
	require.Contains(t, rendered, "Alice")
	require.Contains(t, rendered, "1000000")
	// Ensure separators exist (StyleRounded uses Unicode characters like ┌ ┬ ┐)
	require.Contains(t, rendered, "┌")
	require.Contains(t, rendered, "│")
}
