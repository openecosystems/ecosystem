package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type ConnectorKeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

func (k ConnectorKeyMap) Name() KeyBindingType {
	return Connector
}

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

func ConnectorFullHelp() []key.Binding {
	return []key.Binding{
		ConnectorKeys.Connect,
		ConnectorKeys.Disconnect,
	}
}
