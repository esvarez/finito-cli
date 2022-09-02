package cmd

type uiController interface {
	// RenderAddExpense() error
	Render() error
	SetIncomeForm()
	SetExpenseForm()
}
