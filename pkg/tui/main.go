package tui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

func (v *View) Render() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Here will be the menu"
	p.Title = "Menu"
	p.SetRect(0, 0, 50, 5)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(2.0/3, p),
		),
	)

	ui.Render(grid)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
	return nil
}
