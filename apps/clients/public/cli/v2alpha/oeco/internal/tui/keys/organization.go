package keys

import (
	"fmt"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"
)

type OrganizationKeyMap struct {
	Create key.Binding
}

func (k OrganizationKeyMap) Name() KeyBindingType {
	return Organization
}

var OrganizationKeys = OrganizationKeyMap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "call"),
	),
}

func OrganizationFullHelp() []key.Binding {
	return []key.Binding{
		OrganizationKeys.Create,
	}
}

func rebindOrganizationKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding Organization key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "create":
			k = &OrganizationKeys.Create
		default:
			return fmt.Errorf("unknown built-in organization key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
