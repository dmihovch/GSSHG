package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	println("Message: " + m.msg)

	m.msg = msg.(string)

	return m, nil

}
