package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	println("Message: " + m.msg)

	switch tmsg := msg.(type) {

	case tea.WindowSizeMsg:

		m.msg = "tea window message"

	case string:
		m.msg = tmsg
	}

	return m, nil

}
