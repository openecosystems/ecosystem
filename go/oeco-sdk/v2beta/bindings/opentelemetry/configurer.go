package opentelemetryv1

import (
	"encoding/json"
	"fmt"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

// Configuration represents the configuration settings for Opentelemetry providers used in the binding.
type Configuration struct {
	Opentelemetry Opentelemetry
}

type Opentelemetry struct {
	TraceProviderEnabled  bool
	MeterProviderEnabled  bool
	LoggerProviderEnabled bool
}

// ResolveConfiguration resolves and merges the default and user-provided configurations for the Binding instance.
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

// ValidateConfiguration checks the configuration for binding and logs warnings for disabled Opentelemetry providers.
func (b *Binding) ValidateConfiguration() error {
	if !b.configuration.Opentelemetry.TraceProviderEnabled {
		fmt.Println("warn: opentelemtry trace is disabled. This may cause errors if you have other bindings that depend on it. Binding dependency management is on the roadmap.")
	}

	if b.configuration.Opentelemetry.MeterProviderEnabled {
		fmt.Println("Enabled Open Telemetry MeterProvider")
	}

	if b.configuration.Opentelemetry.LoggerProviderEnabled {
		fmt.Println("Enabled Open Telemetry LoggerProvider")
	}
	return nil
}

// GetDefaultConfiguration returns the default configuration for the Binding, including Opentelemetry provider settings.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	return &Configuration{
		Opentelemetry: struct {
			TraceProviderEnabled  bool
			MeterProviderEnabled  bool
			LoggerProviderEnabled bool
		}{
			TraceProviderEnabled:  false,
			MeterProviderEnabled:  false,
			LoggerProviderEnabled: false,
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
