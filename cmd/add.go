package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type addCmd struct {
	ui uiController
}

func newAddCmd(ui uiController) *addCmd {
	return &addCmd{
		ui: ui,
	}
}

func (a addCmd) command() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "add",
		Short:        "Add an expense",
		Long:         `Add an expense`,
		SilenceUsage: false,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.UsageString())
		},
	}

	cmd.AddCommand(a.addExpense())
	cmd.AddCommand(a.addIncome())

	return cmd
}

func (a addCmd) addExpense() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expense",
		Short: "Add an expense",
		Long:  `Add an expense`,
		Run: func(cmd *cobra.Command, args []string) {
			a.ui.SetExpenseForm()
			a.ui.Render()
		},
	}

	return cmd
}

func (a addCmd) addIncome() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "income",
		Short: "Add an income",
		Long:  `Add an income`,
		Run: func(cmd *cobra.Command, args []string) {
			a.ui.SetIncomeForm()
			a.ui.Render()
		},
	}

	return cmd
}
