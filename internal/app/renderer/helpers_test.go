package renderer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyTableData struct {
	columns []string
	rows    [][]any
	caption string
	err     error
}

func newDummyTableData(columns []string, rows [][]any, caption string) Data {
	return dummyTableData{
		columns: columns,
		rows:    rows,
		caption: caption,
	}
}

func newDummyTableDataWithError(err error) Data {
	return dummyTableData{err: err}
}

func (d dummyTableData) Columns() []string {
	return d.columns
}

func (d dummyTableData) Rows() ([][]any, error) {
	return d.rows, d.err
}

func (d dummyTableData) Caption() string {
	return d.caption
}

func assertContainsFold(t *testing.T, got string, wants ...string) {
	t.Helper()
	lowered := strings.ToLower(got)
	for _, want := range wants {
		assert.Contains(t, lowered, strings.ToLower(want))
	}
}
