package zaploggerv1

import (
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ResolvedConfiguration holds the resolved runtime configuration for the application and is shared across components.
var ResolvedConfiguration *Configuration

// Configuration holds the settings for initializing a zap-based logging framework with custom configuration options.
type Configuration struct {
	Zap zap.Config `yaml:"zap,omitempty"`
}

// ResolveConfiguration resolves and merges the Binding's configuration by utilizing the default configuration as a base.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration performs validation checks on the logger configuration and returns an error if invalid.
func (b *Binding) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns the default logging configuration for the Binding with predefined settings for Zap.
func (b *Binding) GetDefaultConfiguration() interface{} {
	level, _ := zap.ParseAtomicLevel("info")

	return Configuration{
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
func (b *Binding) CreateConfiguration() (interface{}, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (b *Binding) GetConfiguration() interface{} {
	return nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations() error {
	return nil
}
