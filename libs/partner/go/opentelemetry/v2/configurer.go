package opentelemetryv2

import (
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the configuration for this binding
var ResolvedConfiguration *Configuration

// Configuration represents the configuration settings for Opentelemetry providers used in the binding.
type Configuration struct {
	Opentelemetry struct {
		TraceProviderEnabled  bool
		MeterProviderEnabled  bool
		LoggerProviderEnabled bool
	}
}

// ResolveConfiguration resolves and merges the default and user-provided configurations for the Binding instance.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
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
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{
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
