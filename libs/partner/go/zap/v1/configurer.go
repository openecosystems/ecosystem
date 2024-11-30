package zaploggerv1

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

var (
	ResolvedConfiguration *Configuration
)

type Configuration struct {
	Zap zap.Config `yaml:"zap,omitempty"`
}

func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

func (b *Binding) ValidateConfiguration() error {
	return nil
}

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
