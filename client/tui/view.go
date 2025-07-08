package tui

func (m Model) View() string {

	s := "Hello World\n\n"

	s += "the message is:\n"
	s += m.msg

	s += "\nhope you enjoyed the message"
	return s

}
