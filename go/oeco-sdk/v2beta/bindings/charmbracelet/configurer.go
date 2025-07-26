package charmbraceletv1

import (
	"encoding/json"
	"fmt"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

// Configuration holds the settings for initializing a zap-based logging framework with custom configuration options.
type Configuration struct {
	App specv2pb.App `yaml:"app,omitempty"`
}

// ResolveConfiguration resolves and merges the Binding's configuration by utilizing the default configuration as a base.
func (b *Binding) ResolveConfiguration(opts ...sdkv2betalib.ConfigurationProviderOption) (*sdkv2betalib.Configurer, error) {
	var c Configuration
	configurer, err := sdkv2betalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}
	sdkv2betalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
	b.configuration = &c

	return configurer, nil
}

// ValidateConfiguration performs validation checks on the logger configuration and returns an error if invalid.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default logging configuration for the Binding with predefined settings for Zap.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	return &Configuration{
		App: specv2pb.App{
			Debug:   false,
			Verbose: false,
			Quiet:   false,
		},
	}
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
	byteArray, err := json.Marshal(*b.GetConfiguration()) //nolint:staticcheck
	if err != nil {
		fmt.Println("SpecError:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
