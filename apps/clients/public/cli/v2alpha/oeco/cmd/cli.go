package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	connectorv2alphatui "apps/clients/public/cli/v2alpha/oeco/internal/connector/v2alpha"
	ecosystemv2alphapbint "apps/clients/public/cli/v2alpha/oeco/internal/ecosytem/v2alpha"
	iamv2alphapbint "apps/clients/public/cli/v2alpha/oeco/internal/iam/v2alpha"
	markdown "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	charmbraceletloggerv0 "libs/partner/go/charmbracelet/v0"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	cmdv2alphapbcmd "libs/public/go/cli/v2alpha/gen/platform/cmd"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
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
	logFile string
	quiet   bool

	configuration *cliv2alphalib.Configuration
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
		override := cliv2alphalib.Configuration{
			App: specv2pb.App{
				Debug:   debug,
				Verbose: verbose,
				Quiet:   quiet,
			},
		}
		_ = charmbraceletloggerv0.Bound.Override(&charmbraceletloggerv0.Configuration{
			App: specv2pb.App{
				Debug:   debug,
				Verbose: verbose,
				Quiet:   quiet,
			},
		})
		sdkv2alphalib.Merge(&override, configuration)
		cmd.SetContext(context.WithValue(cmd.Root().Context(), sdkv2alphalib.SettingsContextKey, &override))
		cmd.SetContext(context.WithValue(cmd.Context(), sdkv2alphalib.LoggerContextKey, charmbraceletloggerv0.Bound.Logger))
	},
}

// Execute runs the main command-line interface (CLI) program logic, initializing settings, context, and commands.
func Execute() {
	bounds := []sdkv2alphalib.Binding{
		&charmbraceletloggerv0.Binding{},
		//&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
		//
		//}},
		&nebulav1ca.Binding{},
	}

	c := cliv2alphalib.NewCLI(
		context.Background(),
		cliv2alphalib.WithBounds(bounds),
		cliv2alphalib.WithConfigurationProvider(&cliv2alphalib.Configuration{}),
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
func AddCommands(settings *cliv2alphalib.Configuration) {
	cmdv2alphapbcmd.CommandRegistry.RegisterCommands()

	if settings != nil && settings.Systems != nil {
		for _, system := range settings.Systems { //nolint:copylocks,govet
			command, err := cmdv2alphapbcmd.CommandRegistry.GetCommandByFullCommandName(cmdv2alphapbcmd.FullCommandName{
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

			RootCmd.AddCommand(command.Commands()...)
		}
	}

	// Manually add certain system commands
	RootCmd.AddCommand(ecosystemv2alphapbint.EcosystemServiceServiceCmd)
	RootCmd.AddCommand(iamv2alphapbint.AccountAuthorityServiceServiceCmd)
	RootCmd.AddCommand(iamv2alphapbint.AccountAuthorityServiceServiceCmd)
	RootCmd.AddCommand(iamv2alphapbint.AccountServiceServiceCmd)
	RootCmd.AddCommand(connectorv2alphatui.Cmd)
}

// init initializes the logging handler, persistent flags, and markdown styling based on terminal background settings.
func init() {
	log.SetHandler(cli.Default)

	RootCmd.PersistentFlags().StringVar(&ecosystem, "context", "", "context to use for this call")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug level logging")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable additional logging")
	RootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "reduces logging output to only essential messages")
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
