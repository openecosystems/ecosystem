package config

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type KeyBinding struct {
	Key     string `yaml:"key"`
	Command string `yaml:"command"`
	Builtin string `yaml:"builtin"`
}

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

type KeyBindings struct {
	Universal    []KeyBinding `yaml:"universal"`
	Enclave      []KeyBinding `yaml:"enclave"`
	Context      []KeyBinding `yaml:"context"`
	Organization []KeyBinding `yaml:"organization"`
	Package      []KeyBinding `yaml:"package"`
	Connector    []KeyBinding `yaml:"connector"`
	Api          []KeyBinding `yaml:"api"`
	Ecosystem    []KeyBinding `yaml:"ecosystem"`
}

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

	if section == ApiSection {
		for _, kb := range kbs.Api {
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

func (kbs *KeyBindings) ExecuteKeyBinding(key string) tea.Cmd {

	for _, kb := range kbs.Universal {
		if kb.Key != key {
			continue
		}

		log.Debug("executing keybind", "key", kb.Key, "command", kb.Command)
		//return m.runCustomUniversalCommand(kb.Command)
	}

	return nil
}
