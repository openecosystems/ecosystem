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
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	cmdv2alphapbcmd "libs/public/go/cli/v2alpha/gen/platform/cmd"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// DefaultVersion defines the fallback version identifier when no compile-time version is provided.
const (
	DefaultVersion = "0.0.0"
)

// context2 holds a secondary context string for the application.
// debug indicates whether debug mode is enabled.
// version determines if the application should display its version and exit.
// verboseLog toggles verbose logging mode.
// logFile specifies the file where logs should be written.
// quiet suppresses non-essential output.
var (
	context2   string
	debug      bool
	version    bool
	verboseLog bool
	logFile    string
	quiet      bool
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
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("debug logging enabled")
		}
	},
}

// Execute runs the main command-line interface (CLI) program logic, initializing settings, context, and commands.
func Execute(cli *cliv2alphalib.CLI) {
	defer cli.GracefulShutdown()

	err := validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	settingsProvider, err := sdkv2alphalib.NewCLISettingsProvider(&sdkv2alphalib.RuntimeConfigurationOverrides{
		Context:    &context2,
		Logging:    &debug,
		Verbose:    &version,
		VerboseLog: &verboseLog,
		LogFile:    &logFile,
		Quiet:      &quiet,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err = settingsProvider.WatchSettings(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	settings := settingsProvider.GetSettings()

	ctx := context.Background()
	ctx = context.WithValue(ctx, sdkv2alphalib.SettingsContextKey, settings)
	RootCmd.SetContext(ctx)
	AddCommands(settings)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// AddCommands registers and adds commands to the RootCmd based on the provided SpecSettings.
func AddCommands(settings *specv2pb.SpecSettings) {
	cmdv2alphapbcmd.CommandRegistry.RegisterCommands()

	if settings != nil && settings.Systems != nil {
		for _, system := range settings.Systems {
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

	RootCmd.PersistentFlags().StringVar(&context2, "context", "", "context to use for this call")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug level logging")
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
