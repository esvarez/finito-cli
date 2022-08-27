package entity

type Transaction struct {
	Date        string `json:"date"`
	Month       string `json:"month"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
