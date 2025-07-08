package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	msg string
}

func (m model) Init() tea.Cmd {

	return nil
}
