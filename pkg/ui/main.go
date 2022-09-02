package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Controller struct {
	model *model
}

func NewController() *Controller {
	return &Controller{
		model: newModel(),
	}
}

func (c *Controller) Render() error {
	if os.Getenv("HELP_DEBUG") != "" {
		if f, err := tea.LogToFile("debug.log", "help"); err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		} else {
			defer f.Close()
		}
	}

	if err := tea.NewProgram(c.model).Start(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}

	return nil
}

func (c *Controller) SetIncomeForm() {
	c.model.formType = IncomeForm
}

func (c *Controller) SetExpenseForm() {
	c.model.formType = ExpenseForm
}
