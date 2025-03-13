package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

// PackageKeyMap defines key bindings for operations related to packages such as listing and generating them.
type PackageKeyMap struct {
	List     key.Binding
	Generate key.Binding
}

// Name returns the KeyBindingType associated with the PackageKeyMap, which is typically of type 'Package'.
func (k PackageKeyMap) Name() KeyBindingType {
	return Package
}

// PackageKeys defines bindings for the list and generate commands using the PackageKeyMap structure.
var PackageKeys = PackageKeyMap{
	List: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "list"),
	),
	Generate: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "generate"),
	),
}

// PackageFullHelp returns a slice of key bindings for the package, providing help for available actions such as list and generate.
func PackageFullHelp() []key.Binding {
	return []key.Binding{
		PackageKeys.List,
		PackageKeys.Generate,
	}
}

// rebindPackageKeys updates key bindings for predefined package keys based on the provided configuration.
//
//nolint:unused
func rebindPackageKeys(keys []config.KeyBinding) error {
	for _, kb := range keys {
		if kb.Builtin == "" {
			continue
		}

		log.Debug("Rebinding Package key", "builtin", kb.Builtin, "key", kb.Key)

		var k *key.Binding

		switch kb.Builtin {
		case "list":
			k = &PackageKeys.List
		case "generate":
			k = &PackageKeys.Generate
		default:
			return fmt.Errorf("unknown built-in package key: '%s'", kb.Builtin)
		}

		k.SetKeys(kb.Key)
		k.SetHelp(kb.Key, k.Help().Desc)
	}

	return nil
}
