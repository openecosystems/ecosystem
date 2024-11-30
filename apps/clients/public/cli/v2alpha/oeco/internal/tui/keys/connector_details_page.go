package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

const (
	DetailsPageKeyType KeyBindingType = iota
)

type ConnectorDetailsPageKeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

func (k ConnectorDetailsPageKeyMap) Name() KeyBindingType {
	return DetailsPageKeyType
}

var ConnectorDetailsPageKeys = ConnectorDetailsPageKeyMap{
	Connect: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "connect"),
	),
	Disconnect: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "disconnect"),
	),
}

func GetKeyBindings() []key.Binding {
	return []key.Binding{
		ConnectorDetailsPageKeys.Connect,
		ConnectorDetailsPageKeys.Disconnect,
	}
}

func ConnectorDetailsPageFullHelp() []key.Binding {
	return GetKeyBindings()
}
