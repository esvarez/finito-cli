package entity

type Transaction struct {
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
