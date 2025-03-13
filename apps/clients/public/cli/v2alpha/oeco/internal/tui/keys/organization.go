package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

// OrganizationKeyMap defines key bindings specific to organizational operations or actions in the application.
type OrganizationKeyMap struct {
	Create key.Binding
}

// Name returns the KeyBindingType associated with the OrganizationKeyMap, which is Organization.
func (k OrganizationKeyMap) Name() KeyBindingType {
	return Organization
}

// OrganizationKeys provides key bindings specific to organization-related operations, such as creating an organization.
var OrganizationKeys = OrganizationKeyMap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "call"),
	),
}

// OrganizationFullHelp returns a list of key bindings related to organization actions for display or processing.
func OrganizationFullHelp() []key.Binding {
	return []key.Binding{
		OrganizationKeys.Create,
	}
}

// rebindOrganizationKeys updates the key bindings for organizational actions based on provided configurations.
//
//nolint:unused
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
