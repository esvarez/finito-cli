package tui

type ui struct {
	tabs      []string
	activeTab int
	title     string
	body      map[string]func() string
}

func newUI() *ui {
	return &ui{
		tabs:  []string{"Resumen", "Presupuesto", "Transacciones"},
		title: "Resumen",
		body:  make(map[string]func() string),
	}
}

func (t *ui) nextTab() {
	t.activeTab++
	if t.activeTab > len(t.tabs)-1 {
		t.activeTab = 0
	}
	t.title = t.tabs[t.activeTab]
}

func (t *ui) previousTab() {
	t.activeTab--
	if t.activeTab < 0 {
		t.activeTab = len(t.tabs) - 1
	}
	t.title = t.tabs[t.activeTab]
}
