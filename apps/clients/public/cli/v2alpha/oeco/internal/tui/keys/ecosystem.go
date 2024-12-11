package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

type EcosystemKeyMap struct {
	List  key.Binding
	View  key.Binding
	Graph key.Binding
}

func (k EcosystemKeyMap) Name() KeyBindingType {
	return Ecosystem
}

var EcosystemKeys = EcosystemKeyMap{
	List: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "list"),
	),
	View: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "view"),
	),
	Graph: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "graph"),
	),
}

func EcosystemFullHelp() []key.Binding {
	return []key.Binding{
		EcosystemKeys.List,
		EcosystemKeys.View,
		EcosystemKeys.Graph,
	}
}

func rebindEcosystemKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding ecosystem key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "list":
			k = &EcosystemKeys.List
		case "view":
			k = &EcosystemKeys.View
		case "graph":
			k = &EcosystemKeys.Graph
		default:
			return fmt.Errorf("unknown built-in ecosystem key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
