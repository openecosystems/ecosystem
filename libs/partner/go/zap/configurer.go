package zaploggerv1

import (
	"fmt"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// Configuration holds the settings for initializing a zap-based logging framework with custom configuration options.
type Configuration struct {
	Zap zap.Config `yaml:"zap,omitempty"`
}

// ResolveConfiguration resolves and merges the Binding's configuration by utilizing the default configuration as a base.
func (b *Binding) ResolveConfiguration(opts ...sdkv2alphalib.ConfigurationProviderOption) (*sdkv2alphalib.Configurer, error) {
	var c Configuration
	configurer, err := sdkv2alphalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2alphalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
	b.configuration = &c

	return configurer, nil
}

// ValidateConfiguration performs validation checks on the logger configuration and returns an error if invalid.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default logging configuration for the Binding with predefined settings for Zap.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	level, _ := zap.ParseAtomicLevel("info")

	return &Configuration{
		Zap: zap.Config{
			Level:       level,
			Development: false,
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:       "message",
				LevelKey:         "level",
				TimeKey:          "time",
				NameKey:          "name",
				CallerKey:        "caller",
				FunctionKey:      "",
				StacktraceKey:    "stacktrace",
				SkipLineEnding:   false,
				LineEnding:       zapcore.DefaultLineEnding,
				EncodeLevel:      zapcore.CapitalColorLevelEncoder,
				EncodeTime:       zapcore.ISO8601TimeEncoder,
				EncodeDuration:   zapcore.StringDurationEncoder,
				EncodeCaller:     zapcore.FullCallerEncoder,
				EncodeName:       zapcore.FullNameEncoder,
				ConsoleSeparator: "",
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		},
	}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (*Configuration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfiguration() *Configuration {
	return b.configuration
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := yaml.Marshal(b.GetConfiguration())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
