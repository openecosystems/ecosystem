package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

type ApiExplorerKeyMap struct {
	Call      key.Binding
	Synthetic key.Binding
}

func (k ApiExplorerKeyMap) Name() KeyBindingType {
	return Api
}

var ApiKeys = ApiExplorerKeyMap{
	Call: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "call"),
	),
	Synthetic: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "synthetic"),
	),
}

func ApiFullHelp() []key.Binding {
	return []key.Binding{
		ApiKeys.Call,
		ApiKeys.Synthetic,
	}
}

func rebindApiKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding Api key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "call":
			k = &ApiKeys.Call
		case "synthetic":
			k = &ApiKeys.Synthetic
		default:
			return fmt.Errorf("unknown built-in api key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
