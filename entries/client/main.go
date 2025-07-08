package client

import (
	tea "github.com/charmbracelet/bubbletea"
	"gssh/tui"
)

func main() {
	tea.NewProgram(model tea.Model, opts ...tea.ProgramOption)
}
