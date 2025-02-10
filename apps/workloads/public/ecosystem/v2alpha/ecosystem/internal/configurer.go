package internal

import (
	"errors"
	"fmt"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration stores the resolved and finalized configuration for the application.
var ResolvedConfiguration *EcosystemConfiguration

// EcosystemConfiguration represents a structure for application configuration settings, including app, GRPC, and HTTP details.
type EcosystemConfiguration struct {
	App      specv2pb.App      `yaml:"app,omitempty"`
	Platform specv2pb.Platform `yaml:"platform,omitempty"`
	Context  specv2pb.Context  `yaml:"context,omitempty"`

	err error
}

// ResolveConfiguration merges and resolves the environment and default configuration settings into a unified structure.
func (c *EcosystemConfiguration) ResolveConfiguration() {
	_, err := sdkv2alphalib.NewSpecYamlSettingsProvider()
	if err != nil {
		fmt.Println("spec yaml configuration error: ", err)
		c.err = errors.Join(err)
	}

	var config EcosystemConfiguration
	sdkv2alphalib.Resolve(&config, c.GetDefaultConfiguration().(EcosystemConfiguration)) //nolint:govet,copylocks
	name, version, err := sdkv2alphalib.ImportPackageJson()
	if err != nil {
		return
	}

	if config.App.Name == "" {
		config.App.Name = name
	}

	if config.App.Version == "" {
		config.App.Version = version
	}

	ResolvedConfiguration = &config
}

// ValidateConfiguration checks if the configuration instance is valid and returns an error if validation fails.
func (c *EcosystemConfiguration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default `EcosystemConfiguration` instance with preset values for App, Grpc, and Http fields.
func (c *EcosystemConfiguration) GetDefaultConfiguration() interface{} {
	return EcosystemConfiguration{
		App: specv2pb.App{
			Name:            "server",
			Version:         "0.0.0",
			EnvironmentName: "local-1",
			EnvironmentType: "local",
		},
		Platform: specv2pb.Platform{
			Endpoint:            "http://localhost:6577",
			Insecure:            true,
			DnsEndpoints:        []string{"45.63.49.173:4242"},
			DynamicConfigReload: false,
			Mesh: &specv2pb.Mesh{
				Enabled:     true,
				Endpoint:    "http://192.168.100.5:6477",
				Insecure:    true,
				DnsEndpoint: "192.168.100.1",
				Punchy:      true,
			},
		},
	}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (c *EcosystemConfiguration) CreateConfiguration() (interface{}, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (c *EcosystemConfiguration) GetConfiguration() interface{} {
	return nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (c *EcosystemConfiguration) WatchConfigurations() error {
	return nil
}
