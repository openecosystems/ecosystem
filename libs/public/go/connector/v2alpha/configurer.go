package connectorv2alphalib

import (
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration is a pointer to the resolved configuration settings used by the application at runtime.
var ResolvedConfiguration *Configuration

// Configuration represents a primary structure for handling application configuration, extending sdkv2alphalib.App.
type Configuration struct {
	sdkv2alphalib.App

	err error
}

// ResolveConfiguration resolves and merges configuration data from multiple sources into a unified structure.
func (c *Configuration) ResolveConfiguration() {
	_, err := sdkv2alphalib.NewSpecYamlSettingsProvider()
	if err != nil {
		fmt.Println("resolve connector configuration error: ", err)
		c.err = err
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

// ValidateConfiguration checks the integrity and validity of the Configuration object and returns an error if invalid.
func (c *Configuration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default configuration for the application, including app name, version, and environment.
func (c *Configuration) GetDefaultConfiguration() interface{} {
	return Configuration{
		App: sdkv2alphalib.App{
			Name:            "server",
			Version:         "0.0.0",
			EnvironmentName: "local-1",
			EnvironmentType: "local",
		},
	}
}
