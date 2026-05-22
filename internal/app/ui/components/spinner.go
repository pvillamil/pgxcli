package components

import (
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SpinnerModel struct {
	spinner spinner.Model
	caption string

	spinnerStyle lipgloss.Style
	captionStyle lipgloss.Style
}

func NewSpinnerModel(spinnerStyle, captionStyle lipgloss.Style) SpinnerModel {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = spinnerStyle

	return SpinnerModel{
		spinner:      sp,
		caption:      "Postgresing...",
		spinnerStyle: spinnerStyle,
		captionStyle: captionStyle,
	}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	}
	return m, cmd
}

func (m SpinnerModel) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.spinner.View(),
		m.captionStyle.Render(m.caption),
	)
}
