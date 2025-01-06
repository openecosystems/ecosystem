package infrastructurev2alphalib

import (
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the configuration for this binding
var ResolvedConfiguration *Configuration

type Configuration struct {
	sdkv2alphalib.App

	err error
}

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

func (c *Configuration) ValidateConfiguration() error {
	return nil
}

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
