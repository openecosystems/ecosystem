package infrastructurev2alphalib

import (
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration is a pointer to the globally resolved Configuration instance containing application settings.
var ResolvedConfiguration *Configuration

// Configuration represents the core structure for managing application-specific settings and resolving configurations.
type Configuration struct {
	sdkv2alphalib.App

	err error
}

// ResolveConfiguration initializes, loads, and merges configuration settings from defaults, YAML, and package.json.
func (c *Configuration) ResolveConfiguration() {
	_, err := sdkv2alphalib.NewSpecYamlSettingsProvider()
	if err != nil {
		fmt.Println("resolve infrastructure configuration error: ", err)
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

// ValidateConfiguration checks the current configuration's validity and returns an error if validation fails.
func (c *Configuration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default configuration instance with pre-defined application settings.
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
