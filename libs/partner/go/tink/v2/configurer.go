package tinkv2

import sdkv2alphalib "libs/public/go/sdk/v2alpha"

// ResolvedConfiguration holds the current resolved configuration instance derived from default and provided settings.
var ResolvedConfiguration *Configuration

// Configuration is a struct used to hold configuration settings in the application.
type Configuration struct{}

// ResolveConfiguration updates the binding's configuration by merging default and custom configurations.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration checks the current binding configuration for validity and returns an error if invalid.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default instance of the Configuration struct.
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
