package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

type ContextKeyMap struct {
	List key.Binding
	View key.Binding
	Edit key.Binding
}

func (k ContextKeyMap) Name() KeyBindingType {
	return Context
}

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

func ContextFullHelp() []key.Binding {
	return []key.Binding{
		ContextKeys.List,
		ContextKeys.View,
		ContextKeys.Edit,
	}
}

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
