package cliv2alphalib

import (
	"encoding/json"
	"fmt"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// Configuration represents a structure for application configuration settings, including app, GRPC, and HTTP details.
type Configuration struct {
	App      specv2pb.App      `yaml:"app,omitempty"`
	Platform specv2pb.Platform `yaml:"platform,omitempty"`
	Context  specv2pb.Context  `yaml:"context,omitempty"`
	Systems  []specv2pb.SpecSystem

	configuration *Configuration
	// err error
}

// ResolveConfiguration merges and resolves the environment and default configuration settings into a unified structure.
func (c *Configuration) ResolveConfiguration(opts ...sdkv2alphalib.ConfigurationProviderOption) (*sdkv2alphalib.Configurer, error) {
	var config Configuration

	opts = append(opts, sdkv2alphalib.WithConfigPath(sdkv2alphalib.ContextDirectory))
	configurer, err := sdkv2alphalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2alphalib.Resolve(configurer, &config, c.GetDefaultConfiguration())

	config.configuration = &config
	c.configuration = &config

	return configurer, nil
}

// ValidateConfiguration checks if the configuration instance is valid and returns an error if validation fails.
func (c *Configuration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default `Configuration` instance with preset values for App, Grpc, and Http fields.
func (c *Configuration) GetDefaultConfiguration() *Configuration {
	return &Configuration{
		App: specv2pb.App{
			Name:            "oeco",
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
func (c *Configuration) CreateConfiguration() (*Configuration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (c *Configuration) GetConfiguration() *Configuration {
	return c.configuration
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (c *Configuration) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := json.Marshal(c.GetConfiguration())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (c *Configuration) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
