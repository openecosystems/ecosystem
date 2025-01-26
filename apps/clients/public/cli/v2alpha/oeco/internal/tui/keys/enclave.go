package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

// EnclaveKeyMap defines a set of key bindings for interacting with an enclave through list, view, and edit actions.
type EnclaveKeyMap struct {
	List key.Binding
	View key.Binding
	Edit key.Binding
}

// Name returns the KeyBindingType associated with the EnclaveKeyMap, which is Enclave.
func (k EnclaveKeyMap) Name() KeyBindingType {
	return Enclave
}

// EnclaveKeys provides predefined key bindings for enclave actions such as list, view, and edit.
var EnclaveKeys = EnclaveKeyMap{
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

// EnclaveFullHelp returns a slice of key bindings for enclave operations, including list, view, and edit actions.
func EnclaveFullHelp() []key.Binding {
	return []key.Binding{
		EnclaveKeys.List,
		EnclaveKeys.View,
		EnclaveKeys.Edit,
	}
}

// rebindEnclaveKeys updates the key bindings for enclave-specific actions based on provided configuration.
// Returns an error if a provided built-in key is unknown.
//
//nolint:unused
func rebindEnclaveKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding enclave ", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "list":
			k = &EnclaveKeys.List
		case "view":
			k = &EnclaveKeys.View
		case "edit":
			k = &EnclaveKeys.Edit
		default:
			return fmt.Errorf("unknown built-in enclave key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
