package homepage

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap represents a set of key bindings for specific actions such as connecting and disconnecting.
type KeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

// Keys is a KeyMap containing key bindings for connect and disconnect operations.
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

// GetKeyBindings returns a slice of key.Binding used to define the key mappings for various user interactions.
func GetKeyBindings() []key.Binding {
	return []key.Binding{
		Keys.Connect,
		Keys.Disconnect,
	}
}

// FullHelp returns a slice of key bindings representing the complete set of available commands and their descriptions.
func FullHelp() []key.Binding {
	return GetKeyBindings()
}
