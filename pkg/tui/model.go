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

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 96

	columnWidth = 30
)

var (
	// General
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Tabs
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Copy().Border(activeTabBorder, true)

	tabGap = tab.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Title

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Dialog.

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	// Status bar
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.Copy().
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))

	// Page
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

type model struct {
	ui      *ui
	storage *storage

	keys       keyMap
	help       help.Model
	inputStyle lipgloss.Style
	lastKey    string
	quitting   bool
}

func newModel() model {
	return model{
		ui:      newUI(),
		storage: newStorage(),

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
			m.ui.nextTab()
		case key.Matches(msg, m.keys.ShiftTab):
			m.ui.previousTab()
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
	return m, tea.EnterAltScreen
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
	//height := 2 - strings.Count(status, "\n") - strings.Count(helpView, "\n")

	//return m.menu() + status + strings.Repeat("\n", height) + helpView
	return m.menu() + "\n" + status + "\n" + helpView
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

	//// Body
	{
		doc.WriteString(m.getResume() + "\n\n")
	}

	// Status bar
	{
		doc.WriteString(m.getStatusBar() + "\n\n")
	}

	if physicalWidht > 0 {
		docStyle = docStyle.MaxWidth(physicalWidht)
	}

	return docStyle.Render(doc.String())
}

func (m model) getTabs() string {
	var tabs []string
	for i, t := range m.ui.tabs {
		if i == m.ui.activeTab {
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
		SetString(m.ui.title)

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

func (m model) getResume() string {
	var (
		colors = colorGrid(1, 5)
		title  strings.Builder
	)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		MarginBottom(1).
		Foreground(lipgloss.Color("#FFF7DB")).
		SetString("Balance")

	fmt.Fprint(&title, titleStyle.Copy().
		Background(lipgloss.Color(colors[m.ui.activeTab][0])))

	style := lipgloss.NewStyle().Padding(0, 1).Width(50).Align(lipgloss.Left)
	cash := style.Render(fmt.Sprintf("Efectivo  $%.2f", m.storage.getCash()))
	debit := style.Render(fmt.Sprintf("Debito    $%.2f", m.storage.getDebit()))
	credit := style.Render(fmt.Sprintf("Credito   $%.2f", m.storage.getCredit()))
	total := style.Render(fmt.Sprintf("Total     $%.2f", m.storage.getTotal()))
	balance := lipgloss.JoinVertical(lipgloss.Top, cash, debit, credit, total)
	ui := lipgloss.JoinVertical(lipgloss.Center, title.String(), balance)

	dialog := lipgloss.Place(width, 12,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	return dialog
}

func (m model) getStatusBar() string {
	w := lipgloss.Width

	statusKey := statusStyle.Render("STATUS")
	encoding := encodingStyle.Render("UTF-8")
	fishCake := fishCakeStyle.Render("🍥 Fish Cake")
	statusVal := statusText.Copy().
		Width(width - w(statusKey) - w(encoding) - w(fishCake)).
		Render("Ravishing")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
		fishCake,
	)

	return statusBarStyle.Width(width).Render(bar)
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
