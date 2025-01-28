package serverv2alphalib

import (
	"errors"
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration stores the resolved and finalized configuration for the application.
var ResolvedConfiguration *Configuration

// PublicHTTP represents the configuration for a public HTTP server, including its port.
type PublicHTTP struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
}

// MeshHTTP represents the configuration for the HTTP server, including its port.
type MeshHTTP struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
}

// Configuration represents a structure for application configuration settings, including app, GRPC, and HTTP details.
type Configuration struct {
	App        sdkv2alphalib.App `yaml:"app,omitempty"`
	PublicHTTP PublicHTTP        `yaml:"publicHTTP,omitempty"`
	MeshHTTP   MeshHTTP          `yaml:"meshHTTP,omitempty"`

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
	dc := c.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&config, dc)
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
	sdkConfig.App.Trace = config.App.Trace

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
		App: sdkv2alphalib.App{
			Name:            "server",
			Version:         "0.0.0",
			EnvironmentName: "local-1",
			EnvironmentType: "local",
		},
		PublicHTTP: PublicHTTP{
			Host: "0.0.0.0",
			Port: "6577",
		},
		MeshHTTP: MeshHTTP{
			Host: "0.0.0.0",
			Port: "6477",
		},
	}
}
