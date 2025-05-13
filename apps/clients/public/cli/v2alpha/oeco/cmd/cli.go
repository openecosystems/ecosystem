package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	apiv2alphapbint "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/api"
	ecosystemv2alphapbint "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/ecosytem/v2alpha"
	enclavev2alphapbint "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/enclave"
	markdown "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	charmbraceletloggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/charmbracelet"
	nebulav1ca "github.com/openecosystems/ecosystem/libs/partner/go/nebula/ca"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/communication/v1alpha/communicationv1alphapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/communication/v1beta/communicationv1betapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/configuration/v2alpha/configurationv2alphapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/cryptography/v2alpha/cryptographyv2alphapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ecosystem/v2alpha/ecosystemv2alphapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/iam/v2alpha/iamv2alphapbcli"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/system/v2alpha/systemv2alphapbcli"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// DefaultVersion defines the fallback version identifier when no compile-time version is provided.
const (
	DefaultVersion = "0.0.0"
)

// ecosystem holds a secondary context string for the application.
// debug indicates whether debug mode is enabled.
// version determines if the application should display its version and exit.
// verboseLog toggles verbose logging mode.
// logFile specifies the file where logs should be written.
// quiet suppresses non-essential output.
var (
	ecosystem string
	debug     bool
	// version   bool
	verbose bool
	// verboseLog bool
	logFile   string
	quiet     bool
	logToFile bool

	configuration *sdkv2alphalib.CLIConfiguration
)

// compileTimeVersion stores the version set at the time of compilation.
// Version holds the current version of the application determined at runtime.
// Commit contains the git commit hash corresponding to this build.
// Date indicates the build date of the application.
// BuiltBy specifies the entity or user who built the application.
var (
	compileTimeVersion string
	Version            = bestVersion()
	Commit             string
	Date               string
	BuiltBy            string
)

// manuallyImplementedSystems is a map of system names to a boolean indicating if the system requires manual command handling.
var manuallyImplementedSystems = map[string]bool{
	"iam":       true,
	"ecosystem": true,
}

// AboutPlatformCLI represents metadata about the platform's CLI, including version, commit hash, build date, and builder info.
type AboutPlatformCLI struct {
	Version string
	Commit  string
	Date    string
	BuiltBy string
}

// RootCmd is the root command for the CLI, used to interact with Open Ecosystems securely and efficiently.
var RootCmd = &cobra.Command{
	Use:          "oeco",
	Short:        "Connect to Open Ecosystems",
	Long:         `Allows you to securely and efficiently interact with Open Economic Systems `,
	Example:      `oeco --context=my-ecosystem`,
	Version:      Version,
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		override := sdkv2alphalib.CLIConfiguration{
			App: specv2pb.App{
				Debug:     debug,
				Verbose:   verbose,
				Quiet:     quiet,
				LogToFile: logToFile,
			},
		}
		_ = charmbraceletloggerv1.Bound.Override(&charmbraceletloggerv1.Configuration{
			App: specv2pb.App{
				Debug:     debug,
				Verbose:   verbose,
				Quiet:     quiet,
				LogToFile: logToFile,
			},
		})
		sdkv2alphalib.Merge(&override, configuration)
		cmd.SetContext(context.WithValue(cmd.Root().Context(), sdkv2alphalib.SettingsContextKey, &override))
		cmd.SetContext(context.WithValue(cmd.Context(), sdkv2alphalib.LoggerContextKey, charmbraceletloggerv1.Bound.Logger))
		cmd.SetContext(context.WithValue(cmd.Context(), sdkv2alphalib.NebulaCAContextKey, nebulav1ca.Bound))
	},
}

// Execute runs the main command-line interface (CLI) program logic, initializing settings, context, and commands.
func Execute() {
	bounds := []sdkv2alphalib.Binding{
		&charmbraceletloggerv1.Binding{},
		//&natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{
		//
		//}},
		&nebulav1ca.Binding{},
	}

	c := sdkv2alphalib.NewCLI(
		context.Background(),
		sdkv2alphalib.WithCLIBounds(bounds),
		sdkv2alphalib.WithCLIConfigurationProvider(&sdkv2alphalib.CLIConfiguration{}),
	)

	defer c.GracefulShutdown()

	err := validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	configuration = c.GetConfiguration()

	ctx := context.Background()
	RootCmd.SetContext(ctx)
	AddCommands(c.GetConfiguration())

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// AddCommands registers and adds commands to the RootCmd based on the provided SpecSettings.
func AddCommands(settings *sdkv2alphalib.CLIConfiguration) {
	// TODO: Make this dynamic
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "communication", Version: "v1alpha"}, communicationv1alphapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "communication", Version: "v1beta"}, communicationv1betapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "configuration", Version: "v2alpha"}, configurationv2alphapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "cryptography", Version: "v2alpha"}, cryptographyv2alphapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "ecosystem", Version: "v2alpha"}, ecosystemv2alphapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "iam", Version: "v2alpha"}, iamv2alphapbcli.SystemCmd)
	sdkv2alphalib.CommandRegistry.RegisterCommand(sdkv2alphalib.FullCommandName{Name: "system", Version: "v2alpha"}, systemv2alphapbcli.SystemCmd)

	sdkv2alphalib.CommandRegistry.RegisterCommands()

	if settings != nil && settings.Systems != nil {
		for _, system := range settings.Systems { //nolint:copylocks,govet
			command, err := sdkv2alphalib.CommandRegistry.GetCommandByFullCommandName(sdkv2alphalib.FullCommandName{
				Name:    system.Name,
				Version: system.Version,
			})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}

			if _, ok := manuallyImplementedSystems[system.Name]; ok {
				continue
			}

			apiv2alphapbint.APIServiceServiceCmd.AddCommand(command.Commands()...)
		}
	}

	// Manually add certain system commands
	RootCmd.AddCommand(enclavev2alphapbint.EnclaveServiceServiceCmd)
	RootCmd.AddCommand(apiv2alphapbint.APIServiceServiceCmd)
	RootCmd.AddCommand(ecosystemv2alphapbint.EcosystemServiceServiceCmd)
	// RootCmd.AddCommand(connectorv2alphatui.Cmd)
	// Dash
}

// init initializes the logging handler, persistent flags, and markdown styling based on terminal background settings.
func init() {
	RootCmd.PersistentFlags().StringVar(&ecosystem, "context", "", "context to use for this call")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug level logging")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable additional logging")
	RootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "reduces logging output to only essential messages")
	RootCmd.PersistentFlags().BoolVar(&logToFile, "logToFile", false, "log stdout and stderr to the default log file")
	RootCmd.PersistentFlags().StringVar(&logFile, "logFile", "", "log File path (if set, logging enabled automatically)")

	// Set bash-completion
	validConfigFilenames := []string{"yaml", "yml"}
	_ = RootCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)
	_ = RootCmd.PersistentFlags().SetAnnotation("logFile", cobra.BashCompFilenameExt, []string{})

	// see https://github.com/charmbracelet/lipgloss/issues/73
	lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())
	markdown.InitializeMarkdownStyle(termenv.HasDarkBackground())
}

// bestVersion returns the compile-time specified version if available; otherwise, it defaults to "0.0.0".
func bestVersion() string {
	if compileTimeVersion != "" {
		return compileTimeVersion
	}
	return DefaultVersion
}
