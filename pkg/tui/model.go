package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/term"
)

type model struct {
	terminal *terminal

	keys       keyMap
	help       help.Model
	inputStyle lipgloss.Style
	lastKey    string
	quitting   bool
}

func newModel() model {
	return model{
		terminal: newTerminal(),

		keys:       keys,
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Tab):
			m.terminal.nextTab()
		case key.Matches(msg, m.keys.ShiftTab):
			m.terminal.previousTab()
		case key.Matches(msg, m.keys.Up):
			m.lastKey = "↑"
		case key.Matches(msg, m.keys.Down):
			m.lastKey = "↓"
		case key.Matches(msg, m.keys.Left):
			m.lastKey = "←"
		case key.Matches(msg, m.keys.Right):
			m.lastKey = "→"
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		default:
			m.lastKey = msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	var status string
	if m.lastKey == "" {
		status = "Waiting for input..."
	} else {
		status = "You choose:" + m.inputStyle.Render(m.lastKey)
	}

	helpView := m.help.View(m.keys)
	height := 8 - strings.Count(status, "\n") - strings.Count(helpView, "\n")

	return m.menu() + status + strings.Repeat("\n", height) + helpView
}

func (m model) menu() string {
	physicalWidht, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Tabs
	{
		doc.WriteString(m.getTabs() + "\n\n")
	}
	// Title
	{
		doc.WriteString(m.getTitle() + "\n\n")
	}

	if physicalWidht > 0 {
		docStyle = docStyle.MaxWidth(physicalWidht)
	}

	return docStyle.Render(doc.String())
}

func (m model) getTabs() string {
	var tabs []string
	for i, t := range m.terminal.tabs {
		if i == m.terminal.activeTab {
			tabs = append(tabs, activeTab.Render(t))
		} else {
			tabs = append(tabs, tab.Render(t))
		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabs...,
	)
	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	return row
}

func (m model) getTitle() string {
	var (
		colors = colorGrid(1, 5)
		title  strings.Builder
	)

	titleStyle := lipgloss.NewStyle().
		MarginLeft(1).
		MarginRight(5).
		Padding(0, 1).
		Italic(true).
		Foreground(lipgloss.Color("#FFF7DB")).
		SetString(m.terminal.title)

	for i, v := range colors {
		const offset = 2
		c := lipgloss.Color(v[0])
		fmt.Fprint(&title, titleStyle.Copy().MarginLeft(i*offset).Background(c))
		if i < len(colors)-1 {
			title.WriteString("\n")
		}
	}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("Resumen de transacciones"),
		// infoStyle.Render("From Charm"+divider+url("https://github.com/charmbracelet/lipgloss")),
		infoStyle.Render("Erick Suarez"+divider+url("https://github.com/esvarez")),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	return row
}

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
