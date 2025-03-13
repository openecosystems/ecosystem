package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

// APIExplorerKeyMap represents the key bindings specific to API exploration actions in the application.
// Call corresponds to the binding for initiating API calls.
// Synthetic corresponds to the binding for handling synthetic key actions.
type APIExplorerKeyMap struct {
	Call      key.Binding
	Synthetic key.Binding
}

// Name returns the KeyBindingType corresponding to the Api constant.
func (k APIExplorerKeyMap) Name() KeyBindingType {
	return API
}

// APIKeys represents the key bindings for API operations, including 'call' and 'synthetic' functionalities.
var APIKeys = APIExplorerKeyMap{
	Call: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "call"),
	),
	Synthetic: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "synthetic"),
	),
}

// APIFullHelp returns a slice of key bindings representing available API operations and their corresponding shortcuts.
func APIFullHelp() []key.Binding {
	return []key.Binding{
		APIKeys.Call,
		APIKeys.Synthetic,
	}
}

// rebindApiKeys updates the key bindings for API actions based on the provided configuration.
// It assigns new keys for built-in API functionalities like "call" and "synthetic".
// Returns an error for unrecognized built-in keys.
//
//nolint:unused
func rebindAPIKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding Api key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "call":
			k = &APIKeys.Call
		case "synthetic":
			k = &APIKeys.Synthetic
		default:
			return fmt.Errorf("unknown built-in api key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
