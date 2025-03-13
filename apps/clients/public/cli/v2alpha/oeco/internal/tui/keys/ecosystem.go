package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

// EcosystemKeyMap defines key bindings for navigating ecosystem-related commands or actions.
type EcosystemKeyMap struct {
	List  key.Binding
	View  key.Binding
	Graph key.Binding
}

// Name returns the KeyBindingType associated with the EcosystemKeyMap, which is Ecosystem.
func (k EcosystemKeyMap) Name() KeyBindingType {
	return Ecosystem
}

// EcosystemKeys represents the key bindings for the ecosystem commands: list, view, and graph functionality.
var EcosystemKeys = EcosystemKeyMap{
	List: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", list),
	),
	View: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", view),
	),
	Graph: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "graph"),
	),
}

// EcosystemFullHelp returns a list of key bindings used for navigating and interacting with the ecosystem functionality.
func EcosystemFullHelp() []key.Binding {
	return []key.Binding{
		EcosystemKeys.List,
		EcosystemKeys.View,
		EcosystemKeys.Graph,
	}
}

// rebindEcosystemKeys updates the key bindings for ecosystem functionalities based on the provided key bindings configuration.
// Each key in the provided configuration is mapped to a corresponding built-in functionality or an error is returned.
// Passing a key binding with an unknown "builtin" value results in an error.
//
//nolint:unused
func rebindEcosystemKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding ecosystem key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case list:
			k = &EcosystemKeys.List
		case view:
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
