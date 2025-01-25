package config

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// KeyBinding represents a configuration for mapping a key to a command or a built-in functionality.
type KeyBinding struct {
	Key     string `yaml:"key"`
	Command string `yaml:"command"`
	Builtin string `yaml:"builtin"`
}

// NewBinding creates a new key binding with keys and help derived from the current KeyBinding and an optional previous binding.
func (kb KeyBinding) NewBinding(previous *key.Binding) key.Binding {
	helpDesc := ""
	if previous != nil {
		helpDesc = previous.Help().Desc
	}

	return key.NewBinding(
		key.WithKeys(kb.Key),
		key.WithHelp(kb.Key, helpDesc),
	)
}

// KeyBindings defines a structure for organizing key bindings across various application sections.
// It categorizes bindings into universal, enclave, context, organization, package, connector, api, and ecosystem sections.
type KeyBindings struct {
	Universal    []KeyBinding `yaml:"universal"`
	Enclave      []KeyBinding `yaml:"enclave"`
	Context      []KeyBinding `yaml:"context"`
	Organization []KeyBinding `yaml:"organization"`
	Package      []KeyBinding `yaml:"package"`
	Connector    []KeyBinding `yaml:"connector"`
	API          []KeyBinding `yaml:"api"`
	Ecosystem    []KeyBinding `yaml:"ecosystem"`
}

// IsUserDefinedKeyBinding checks if a specified key message corresponds to a user-defined key binding in a given section.
func (kbs *KeyBindings) IsUserDefinedKeyBinding(msg tea.KeyMsg, section SectionType) bool {
	for _, kb := range kbs.Universal {
		if kb.Builtin == "" && kb.Key == msg.String() {
			return true
		}
	}

	if section == EnclaveSection {
		for _, kb := range kbs.Enclave {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	if section == ContextSection {
		for _, kb := range kbs.Context {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	if section == OrganizationSection {
		for _, kb := range kbs.Organization {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	if section == ConnectorSection {
		for _, kb := range kbs.Connector {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	if section == APISection {
		for _, kb := range kbs.API {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	if section == EcosystemSection {
		for _, kb := range kbs.Ecosystem {
			if kb.Builtin == "" && kb.Key == msg.String() {
				return true
			}
		}
	}

	return false
}

// ExecuteKeyBinding executes the command associated with the given key using the universal key bindings. Returns a tea.Cmd.
func (kbs *KeyBindings) ExecuteKeyBinding(key string) tea.Cmd {
	for _, kb := range kbs.Universal {
		if kb.Key != key {
			continue
		}

		log.Debug("executing keybind", "key", kb.Key, "command", kb.Command)
		// return m.runCustomUniversalCommand(kb.Command)
	}

	return nil
}
