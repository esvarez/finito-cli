package tui

type storage struct {
	cash   int
	debit  int
	credit int
	total  int
}

func newStorage() *storage {
	return &storage{
		cash:   3500,
		debit:  500,
		credit: -500,
	}
}

func (s storage) getCash() float32 {
	return float32(s.cash / 100)
}

func (s storage) getDebit() float32 {
	return float32(s.debit / 100)
}

func (s storage) getCredit() float32 {
	return float32(s.credit / 100)
}

func (s storage) getTotal() float32 {
	total := s.cash + s.debit + s.credit
	return float32(total / 100)
}
