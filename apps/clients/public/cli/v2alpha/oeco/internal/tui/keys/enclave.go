package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

type EnclaveKeyMap struct {
	List key.Binding
	View key.Binding
	Edit key.Binding
}

func (k EnclaveKeyMap) Name() KeyBindingType {
	return Enclave
}

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

func EnclaveFullHelp() []key.Binding {
	return []key.Binding{
		EnclaveKeys.List,
		EnclaveKeys.View,
		EnclaveKeys.Edit,
	}
}

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
