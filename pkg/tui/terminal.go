package tui

type terminal struct {
	tabs      []string
	activeTab int
	title     string
}

func newTerminal() *terminal {
	return &terminal{
		tabs:  []string{"Resumen", "Presupuesto", "Transacciones"},
		title: "Resumen",
	}
}

func (t *terminal) nextTab() {
	t.activeTab++
	if t.activeTab > len(t.tabs)-1 {
		t.activeTab = 0
	}
	t.title = t.tabs[t.activeTab]
}

func (t *terminal) previousTab() {
	t.activeTab--
	if t.activeTab < 0 {
		t.activeTab = len(t.tabs) - 1
	}
	t.title = t.tabs[t.activeTab]
}
