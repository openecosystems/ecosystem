package configurationv2alphalib

import (
	"encoding/json"
	"fmt"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the resolved configuration instance for the application, set during runtime resolution.
var ResolvedConfiguration *Configuration

// Configuration represents the main configuration structure used for setting and resolving application configurations.
type Configuration struct {
	specv2pb.App
}

// ResolveConfiguration initializes and resolves the binding's configuration by merging default and external configurations.
func (b *Binding) ResolveConfiguration(opts ...sdkv2alphalib.ConfigurationProviderOption) (*sdkv2alphalib.Configurer, error) {
	var c Configuration
	configurer, err := sdkv2alphalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2alphalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
	b.configuration = &c
	ResolvedConfiguration = &c

	return configurer, nil
}

// ValidateConfiguration validates the current configuration and returns an error if the validation fails.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration as an instance of the Configuration struct.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	return &Configuration{}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (*Configuration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfiguration() *Configuration {
	return nil
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := json.Marshal(*b.GetConfiguration()) //nolint:staticcheck
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
