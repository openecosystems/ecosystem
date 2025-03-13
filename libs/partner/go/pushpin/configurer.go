package pushpinv1

import (
	"encoding/json"
	"fmt"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration is a pointer to a Configuration instance storing the currently resolved configuration settings.
var ResolvedConfiguration *Configuration

// Configuration represents a type for managing configurable settings or parameters.
type Configuration struct{}

// ResolveConfiguration populates the binding's configuration by merging the default and current configuration values.
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

// ValidateConfiguration validates the current configuration and ensures it meets the required criteria or constraints.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration instance used by the Binding.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	return &Configuration{}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (*Configuration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfiguration() *Configuration {
	return b.configuration
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := json.Marshal(*b.GetConfiguration())
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
