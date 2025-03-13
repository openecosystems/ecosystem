package charmbraceletloggerv0

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/log"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// Binding represents a logging framework binding utilizing Uber Zap for structured logging.
// It includes both a standard logger and a sugared logger for flexible usage within applications.
// This type also manages the configuration for initializing the logger effectively.
type Binding struct {
	Logger *log.Logger

	configuration *Configuration
}

// Bound is a global variable representing an instance of the Binding structure for logging purposes.
// BindingName is a constant string used as the identifier for the logging binding.
var (
	Bound       *Binding
	BindingName = "CHARMBRACELET_LOGGING_BINDING"

	logFile *os.File
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
				logger := log.NewWithOptions(os.Stderr, log.Options{
					ReportCaller:    false,
					ReportTimestamp: true,
					TimeFormat:      time.Kitchen,
					Prefix:          "",
				})
				logger.SetStyles(GetDefaultStyles())

				if b.configuration.App.Debug {
					logger.SetLevel(log.DebugLevel)
				}

				if b.configuration.App.Verbose {
					logger.SetReportCaller(true)
				}

				if b.configuration.App.Quiet {
					logger.SetLevel(log.ErrorLevel)
				}

				b.Logger = logger

				if err != nil {
					fmt.Println(fmt.Errorf("could not build Charmbracelet logger: %v", err))
				}

				Bound = &Binding{
					Logger: b.Logger,

					configuration: b.configuration,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Charmbracelet Logging already bound")
	}

	return bindings
}

// GetBinding returns the current instance of the Bound binding.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close shuts down the Charbracelet Logger Binding and performs necessary cleanup operations.
func (b *Binding) Close() error {
	logFile.Close() // nolint:errcheck
	return nil
}

// Override applies configuration changes to the Binding's logger based on debug, verbose, and quiet flags.
func (b *Binding) Override(conf *Configuration) error {
	if conf.App.GetDebug() {
		b.Logger.SetLevel(log.DebugLevel)
	}

	if conf.App.GetVerbose() {
		b.Logger.SetReportCaller(true)
	}

	if conf.App.GetQuiet() {
		b.Logger.SetLevel(log.ErrorLevel)
	}

	if conf.App.GetLogToFile() {
		var fileErr error

		logFile, fileErr = os.OpenFile(sdkv2alphalib.OecoLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if fileErr == nil {
			b.Logger.SetOutput(logFile)
			b.Logger.SetTimeFormat(time.RFC3339)
			b.Logger.Debug("\n\nLogging to " + sdkv2alphalib.OecoLogFile)
		}
	}

	return nil
}
