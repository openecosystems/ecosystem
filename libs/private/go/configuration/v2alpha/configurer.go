package configurationv2alphalib

import (
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the resolved configuration instance for the application, set during runtime resolution.
var ResolvedConfiguration *Configuration

// Configuration represents the main configuration structure used for setting and resolving application configurations.
type Configuration struct{}

// ResolveConfiguration initializes and resolves the binding's configuration by merging default and external configurations.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration validates the current configuration and returns an error if the validation fails.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration as an instance of the Configuration struct.
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
