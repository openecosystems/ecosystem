package sdkv2alphalib

// ResolvedConfiguration holds the final resolved and unified configuration settings for the application at runtime.
var ResolvedConfiguration *Configuration

// Configurable represents an interface for components capable of resolving, validating, and providing default configurations.
type Configurable interface {
	ResolveConfiguration()
	GetDefaultConfiguration() interface{}
	ValidateConfiguration() error
}

// App represents the configuration for an application, including its name, version, environment, and debugging options.
type App struct {
	Name            string `yaml:"name,omitempty" env:"SERVICE_NAME"`
	Version         string `yaml:"version,omitempty" env:"VERSION_NUMBER"`
	EnvironmentName string `yaml:"environmentName,omitempty" env:"ENV_NAME"`
	EnvironmentType string `yaml:"environmentType,omitempty" env:"ENV_TYPE"`
	Trace           bool   `yaml:"trace,omitempty" env:"TRACE"`
	Debug           bool   `yaml:"debug,omitempty" env:"DEBUG"`
}

// Configuration represents the main structure for application-specific configuration settings.
type Configuration struct {
	App App `yaml:"app,omitempty"`
}
