package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

// ConnectorKeyMap defines key bindings for connecting and disconnecting operations within a connector scope.
type ConnectorKeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

// Name returns the KeyBindingType constant representing the connector category.
func (k ConnectorKeyMap) Name() KeyBindingType {
	return Connector
}

// ConnectorKeys defines key bindings for connecting and disconnecting actions within the ConnectorKeyMap structure.
var ConnectorKeys = ConnectorKeyMap{
	Connect: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "connect"),
	),
	Disconnect: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "disconnect"),
	),
}

// ConnectorFullHelp returns a list of key bindings specific to connector actions such as connect and disconnect.
func ConnectorFullHelp() []key.Binding {
	return []key.Binding{
		ConnectorKeys.Connect,
		ConnectorKeys.Disconnect,
	}
}
