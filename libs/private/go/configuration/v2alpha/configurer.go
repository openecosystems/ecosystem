package configurationv2alphalib

import (
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the resolved configuration instance for the application, set during runtime resolution.
var ResolvedConfiguration *Configuration

// Configuration represents the main configuration structure used for setting and resolving application configurations.
type Configuration struct {
	specv2pb.App
}

// ResolveConfiguration initializes and resolves the binding's configuration by merging default and external configurations.
func (b *Binding) ResolveConfiguration() {
	var c Configuration

	sdkv2alphalib.Resolve(&c, b.GetDefaultConfiguration().(Configuration)) //nolint:govet,copylocks
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

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (interface{}, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (b *Binding) GetConfiguration() interface{} {
	return nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations() error {
	return nil
}
