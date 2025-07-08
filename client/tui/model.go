package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	msg string
}

func CreateModel() tea.Model {
	return Model{msg: "Hello"}
}

func (m Model) Init() tea.Cmd {

	return nil
}
