package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

// DetailsPageKeyType represents a key binding type specifically for details page actions in the system.
const (
	DetailsPageKeyType KeyBindingType = iota
)

// ConnectorDetailsPageKeyMap defines key bindings for actions on the connector details page.
type ConnectorDetailsPageKeyMap struct {
	Connect    key.Binding
	Disconnect key.Binding
}

// Name returns the KeyBindingType associated with the ConnectorDetailsPageKeyMap, specifically DetailsPageKeyType.
func (k ConnectorDetailsPageKeyMap) Name() KeyBindingType {
	return DetailsPageKeyType
}

// ConnectorDetailsPageKeys defines the key bindings for actions on the Connector Details page, such as connect and disconnect.
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

// GetKeyBindings returns a slice of key bindings for the connector details page functions, such as connect and disconnect.
func GetKeyBindings() []key.Binding {
	return []key.Binding{
		ConnectorDetailsPageKeys.Connect,
		ConnectorDetailsPageKeys.Disconnect,
	}
}

// ConnectorDetailsPageFullHelp returns key bindings for actions specific to the connector details page, such as connect and disconnect.
func ConnectorDetailsPageFullHelp() []key.Binding {
	return GetKeyBindings()
}
