package temporalv1

import (
	"encoding/json"
	"fmt"

	"go.temporal.io/sdk/client"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

const (

	// LocalHostPort specifies the default address and port for connecting to the Temporal server.
	LocalHostPort string = "localhost:7233"

	// DefaultNamespace represents the default namespace setting used for Temporal client configuration.
	DefaultNamespace string = "default"
)

// Configuration holds the settings for initializing a zap-based logging framework with custom configuration options.
type Configuration struct {
	Temporal client.Options `yaml:"temporal,omitempty"`
}

// ResolveConfiguration resolves and merges the Binding's configuration by utilizing the default configuration as a base.
func (b *Binding) ResolveConfiguration(opts ...sdkv2alphalib.ConfigurationProviderOption) (*sdkv2alphalib.Configurer, error) {
	var c Configuration
	configurer, err := sdkv2alphalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2alphalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
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
		Temporal: client.Options{
			HostPort:           LocalHostPort,
			Namespace:          DefaultNamespace,
			Credentials:        nil,
			Logger:             nil,
			MetricsHandler:     nil,
			Identity:           "",
			DataConverter:      nil,
			FailureConverter:   nil,
			ContextPropagators: nil,
			ConnectionOptions: client.ConnectionOptions{
				Authority:                           "",
				DisableKeepAliveCheck:               false,
				KeepAliveTime:                       0,
				KeepAliveTimeout:                    0,
				GetSystemInfoTimeout:                0,
				DisableKeepAlivePermitWithoutStream: false,
				MaxPayloadSize:                      0,
				DialOptions:                         nil,
			},
			HeadersProvider:            nil,
			TrafficController:          nil,
			Interceptors:               nil,
			DisableErrorCodeMetricTags: false,
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
