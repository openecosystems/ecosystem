package requests_page

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

var Keys = KeyMap{
	Connect: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "connect"),
	),
	Disconnect: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "disconnect"),
	),
}

func GetKeyBindings() []key.Binding {
	return []key.Binding{
		Keys.Connect,
		Keys.Disconnect,
	}
}

func FullHelp() []key.Binding {
	return GetKeyBindings()
}
