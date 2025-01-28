package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

// ContextKeyMap represents a structure to define key bindings for different context actions such as List, View, and Edit.
type ContextKeyMap struct {
	List key.Binding
	View key.Binding
	Edit key.Binding
}

// Name returns the KeyBindingType associated with the ContextKeyMap, specifically the Context value.
func (k ContextKeyMap) Name() KeyBindingType {
	return Context
}

// ContextKeys defines key bindings for actions within a context, including list, view, and edit functionalities.
var ContextKeys = ContextKeyMap{
	List: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "list"),
	),
	View: key.NewBinding(
		key.WithKeys("V", " "),
		key.WithHelp("V/space", "view"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
}

// ContextFullHelp returns a slice of key bindings that represent available actions in a given context (list, view, edit).
func ContextFullHelp() []key.Binding {
	return []key.Binding{
		ContextKeys.List,
		ContextKeys.View,
		ContextKeys.Edit,
	}
}

// rebindContextKeys updates the key bindings in ContextKeys based on the provided slice of KeyBinding configurations.
// Any unknown built-in context key in the provided slice will result in an error being returned.
//
//nolint:unused
func rebindContextKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding context ", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "list":
			k = &ContextKeys.List
		case "view":
			k = &ContextKeys.View
		case "edit":
			k = &ContextKeys.Edit
		default:
			return fmt.Errorf("unknown built-in context key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
