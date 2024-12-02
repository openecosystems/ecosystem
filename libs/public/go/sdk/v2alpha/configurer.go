package sdkv2alphalib

// ResolvedConfiguration holds the configuration for this binding
var ResolvedConfiguration *Configuration

type Configurable interface {
	ResolveConfiguration()
	GetDefaultConfiguration() interface{}
	ValidateConfiguration() error
}

type App struct {
	Name            string `yaml:"name,omitempty" env:"SERVICE_NAME"`
	Version         string `yaml:"version,omitempty" env:"VERSION_NUMBER"`
	EnvironmentName string `yaml:"environmentName,omitempty" env:"ENV_NAME"`
	EnvironmentType string `yaml:"environmentType,omitempty" env:"ENV_TYPE"`
	Trace           bool   `yaml:"trace,omitempty" env:"TRACE"`
	Debug           bool   `yaml:"debug,omitempty" env:"DEBUG"`
}

type Configuration struct {
	App App `yaml:"app,omitempty"`
}
