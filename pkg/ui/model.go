package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	IncomeForm  = "IncomeForm"
	ExpenseForm = "ExpenseForm"
)

type model struct {
	helper   *keyHelp
	formType string

	lastKey  string
	quitting bool
}

func newModel() *model {
	return &model{
		helper: newKeyHelp(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.helper.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.helper.keys.Tab):

		case key.Matches(msg, m.helper.keys.ShiftTab):

		case key.Matches(msg, m.helper.keys.Up):
			m.lastKey = "↑"
		case key.Matches(msg, m.helper.keys.Down):
			m.lastKey = "↓"
		case key.Matches(msg, m.helper.keys.Left):
			m.lastKey = "←"
		case key.Matches(msg, m.helper.keys.Right):
			m.lastKey = "→"
		case key.Matches(msg, m.helper.keys.Help):
			m.helper.help.ShowAll = !m.helper.help.ShowAll
		case key.Matches(msg, m.helper.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		default:
			m.lastKey = msg.String()
		}
	}
	return m, tea.EnterAltScreen
}

func (m model) View() string {
	switch m.formType {
	case IncomeForm:
		return m.incomeFormView()
	case ExpenseForm:
		return m.expenseFormView()
	default:
	}
	if m.quitting {
		return "Bye!\n"
	}

	var status string
	if m.lastKey == "" {
		status = "Waiting for input..."
	} else {
		//status = "You choose:" + m.inputStyle.Render(m.lastKey)
		status = "You choose:" + m.lastKey
	}

	helpView := m.helper.help.View(m.helper.keys)
	//height := 2 - strings.Count(status, "\n") - strings.Count(helpView, "\n")

	//return m.menu() + status + strings.Repeat("\n", height) + helpView
	return status + "\n" + helpView
}

func (m model) incomeFormView() string {
	return "Income Form"
}

func (m model) expenseFormView() string {
	return "Expense Form"
}
