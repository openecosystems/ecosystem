package serverv2alphalib

import (
	"errors"
	"fmt"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration stores the resolved and finalized configuration for the application.
var ResolvedConfiguration *Configuration

// Configuration represents a structure for application configuration settings, including app, GRPC, and HTTP details.
type Configuration struct {
	App      specv2pb.App      `yaml:"app,omitempty"`
	Platform specv2pb.Platform `yaml:"platform,omitempty"`
	Context  specv2pb.Context  `yaml:"context,omitempty"`

	err error
}

// ResolveConfiguration merges and resolves the environment and default configuration settings into a unified structure.
func (c *Configuration) ResolveConfiguration() {
	_, err := sdkv2alphalib.NewSpecYamlSettingsProvider()
	if err != nil {
		fmt.Println("spec yaml configuration error: ", err)
		c.err = errors.Join(err)
	}

	var config Configuration
	sdkv2alphalib.Resolve(&config, c.GetDefaultConfiguration().(Configuration)) //nolint:govet,copylocks
	var sdkConfig sdkv2alphalib.Configuration
	sdkv2alphalib.ImportPackageJson(&sdkConfig)

	if sdkConfig.App.Name != "" {
		config.App.Name = sdkConfig.App.Name
	}

	if sdkConfig.App.Version != "" {
		config.App.Version = sdkConfig.App.Version
	}

	if config.App.EnvironmentName != "" {
		sdkConfig.App.EnvironmentName = config.App.EnvironmentName
	}

	if config.App.EnvironmentType != "" {
		sdkConfig.App.EnvironmentType = config.App.EnvironmentType
	}

	sdkConfig.App.Debug = config.App.Debug
	sdkConfig.App.Verbose = config.App.Verbose

	ResolvedConfiguration = &config
	sdkv2alphalib.ResolvedConfiguration = &sdkConfig
}

// ValidateConfiguration checks if the configuration instance is valid and returns an error if validation fails.
func (c *Configuration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default `Configuration` instance with preset values for App, Grpc, and Http fields.
func (c *Configuration) GetDefaultConfiguration() interface{} {
	return Configuration{
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
