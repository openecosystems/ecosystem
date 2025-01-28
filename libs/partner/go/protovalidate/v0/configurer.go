package protovalidatev0

import (
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration is a globally accessible pointer to the current resolved Configuration instance.
var ResolvedConfiguration *Configuration

// Configuration is a placeholder type used for defining and resolving application-specific configurations.
type Configuration struct{}

// ResolveConfiguration resolves and applies the default configuration to the Binding instance.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration validates the current configuration associated with the binding and returns an error if invalid.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration instance for the binding.
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
