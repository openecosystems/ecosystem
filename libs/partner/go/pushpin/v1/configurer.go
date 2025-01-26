package pushpinv1

import sdkv2alphalib "libs/public/go/sdk/v2alpha"

// ResolvedConfiguration is a pointer to a Configuration instance storing the currently resolved configuration settings.
var ResolvedConfiguration *Configuration

// Configuration represents a type for managing configurable settings or parameters.
type Configuration struct{}

// ResolveConfiguration populates the binding's configuration by merging the default and current configuration values.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration validates the current configuration and ensures it meets the required criteria or constraints.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration instance used by the Binding.
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
