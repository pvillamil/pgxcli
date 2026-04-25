package renderer

import (
	"io"

	"github.com/balaji01-4d/pgxcli/internal/config"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

type Data interface {
	Columns() []string
	Rows() ([][]any, error)
	Caption() string
}

func Table(data Data, w io.Writer, c *config.Config) error {
	t := tablewriter.NewTable(w, tablewriter.WithRenderer(renderer.NewColorized(GetTableStyle(c))))
	rows, err := data.Rows()
	if err != nil {
		return err
	}

	t.Header(data.Columns())
	if err := t.Bulk(rows); err != nil {
		return err
	}

	captionColor := getCaptionColor(c.Table.Color.Caption)
	colorizedCaption := color.New(captionColor).Sprint(data.Caption())

	caption := tw.Caption{
		Text: colorizedCaption,
		Spot: tw.SpotBottomLeft,
	}
	t.Caption(caption)
	return t.Render()
}
