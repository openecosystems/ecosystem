package serverv2alphalib

import (
	"errors"
	"fmt"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the configuration for this binding
var ResolvedConfiguration *Configuration

type Grpc struct {
	Port string `yaml:"port,omitempty"`
}

type Http struct {
	Port string `yaml:"port,omitempty"`
}

type Configuration struct {
	App  sdkv2alphalib.App `yaml:"app,omitempty"`
	Grpc Grpc              `yaml:"grpc,omitempty"`
	Http Http              `yaml:"http,omitempty"`

	err error
}

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
		Grpc: Grpc{
			Port: "6510",
		},
		Http: Http{
			Port: "6410",
		},
	}
}
