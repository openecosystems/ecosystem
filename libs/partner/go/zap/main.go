package zaploggerv1

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"go.uber.org/zap"
)

// Binding represents a logging framework binding utilizing Uber Zap for structured logging.
// It includes both a standard logger and a sugared logger for flexible usage within applications.
// This type also manages the configuration for initializing the logger effectively.
type Binding struct {
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger

	LoggerWrapper        *ZapLoggerWrapper
	SugaredLoggerWrapper *ZapSugaredLoggerWrapper

	configuration *Configuration
}

// Bound is a global variable representing an instance of the Binding structure for logging purposes.
// BindingName is a constant string used as the identifier for the logging binding.
var (
	Bound       *Binding
	BindingName = "ZAP_LOGGING_BINDING"
)

// Name returns the unique name identifier for the Binding.
func (b *Binding) Name() string {
	return BindingName
}

// Validate performs validation checks on the Binding within the provided context and bindings.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any log requirements
	return nil
}

// Bind attaches the binding to the provided bindings structure, initializes logger if necessary, and returns updated bindings.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				var err error
				b.Logger, err = b.configuration.Zap.Build()
				if err != nil {
					fmt.Println(fmt.Errorf("could not build Zap logger: %v", err))
				}

				defer b.Logger.Sync() //nolint:errcheck

				Bound = &Binding{
					Logger:               b.Logger,
					SugaredLogger:        b.Logger.Sugar(),
					LoggerWrapper:        NewZapLoggerWrapper(b.Logger),
					SugaredLoggerWrapper: NewZapSugaredLoggerWrapper(b.Logger.Sugar()),

					configuration: b.configuration,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Zap Logging already bound")
	}

	return bindings
}

// GetBinding returns the current instance of the Bound binding.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close shuts down the Uber Zap Logger Binding and performs necessary cleanup operations.
func (b *Binding) Close() error {
	fmt.Println("Closing the Uber Zap Logger Binding")
	return nil
}

// ZapLoggerWrapper wraps a zap.Logger to implement log-compatible interfaces like Printf for structured logging.
// It provides a convenient way to integrate zap.Logger into libraries expecting standard logging interfaces.
type ZapLoggerWrapper struct {
	logger *zap.Logger
}

// NewZapLoggerWrapper creates and returns a new instance of ZapLoggerWrapper using the provided zap.Logger instance.
func NewZapLoggerWrapper(z *zap.Logger) *ZapLoggerWrapper {
	return &ZapLoggerWrapper{
		logger: z,
	}
}

// Printf logs a formatted message at the INFO level using the zap.Logger instance.
func (z *ZapLoggerWrapper) Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	z.logger.Info(msg)
}

// Debug logs a debug-level message with key-value pairs for structured logging using a sugared logger.
func (z *ZapLoggerWrapper) Debug(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Debugw(msg, keyvals...)
}

// Info logs an informational message with structured key-value pair fields using the SugaredLogger.
func (z *ZapLoggerWrapper) Info(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Infow(msg, keyvals...)
}

// Warn logs a warning-level message with optional structured key-value pairs using the sugared logger.
func (z *ZapLoggerWrapper) Warn(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Warnw(msg, keyvals...)
}

// Error logs an error message with optional key-value pairs using the SugaredLogger for structured logging.
func (z *ZapLoggerWrapper) Error(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Errorw(msg, keyvals...)
}

// ZapSugaredLoggerWrapper is a wrapper around zap.SugaredLogger to provide custom logging functionalities.
// It includes methods for structured and formatted logging leveraging the underlying zap.SugaredLogger instance.
type ZapSugaredLoggerWrapper struct {
	logger *zap.SugaredLogger
}

// NewZapSugaredLoggerWrapper initializes and returns a new ZapSugaredLoggerWrapper using the provided SugaredLogger instance.
func NewZapSugaredLoggerWrapper(z *zap.SugaredLogger) *ZapSugaredLoggerWrapper {
	return &ZapSugaredLoggerWrapper{
		logger: z,
	}
}

// Printf logs a formatted message at the Info level using the provided format and arguments.
func (z *ZapSugaredLoggerWrapper) Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	z.logger.Info(msg)
}

// Debug logs a debug-level message with key-value pairs for structured logging using a sugared logger.
func (z *ZapSugaredLoggerWrapper) Debug(msg string, keyvals ...interface{}) {
	z.logger.Debugw(msg, keyvals...)
}

// Info logs an informational message with structured key-value pair fields using the SugaredLogger.
func (z *ZapSugaredLoggerWrapper) Info(msg string, keyvals ...interface{}) {
	z.logger.Infow(msg, keyvals...)
}

// Warn logs a warning-level message with optional structured key-value pairs using the sugared logger.
func (z *ZapSugaredLoggerWrapper) Warn(msg string, keyvals ...interface{}) {
	z.logger.Warnw(msg, keyvals...)
}

// Error logs an error message with optional key-value pairs using the SugaredLogger for structured logging.
func (z *ZapSugaredLoggerWrapper) Error(msg string, keyvals ...interface{}) {
	z.logger.Errorw(msg, keyvals...)
}
