package keys

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/log"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
)

type PackageKeyMap struct {
	List     key.Binding
	Generate key.Binding
}

func (k PackageKeyMap) Name() KeyBindingType {
	return Package
}

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

func PackageFullHelp() []key.Binding {
	return []key.Binding{
		PackageKeys.List,
		PackageKeys.Generate,
	}
}

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
