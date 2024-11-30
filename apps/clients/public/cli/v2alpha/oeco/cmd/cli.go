package cmd

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"

	"apps/clients/public/cli/v2alpha/oeco/internal/configuration/v2alpha"
	"apps/clients/public/cli/v2alpha/oeco/internal/connector/v2alpha"
	"apps/clients/public/cli/v2alpha/oeco/internal/cryptography/v2alpha"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/public/go/cli/v2alpha"
	"libs/public/go/cli/v2alpha/gen/platform/cmd"
	"libs/public/go/sdk/v2alpha"
)

const (
	DefaultVersion = "0.0.0"
)

// Global Flags that can override configuration at runtime
var (
	context2   string
	debug      bool
	version    bool
	verboseLog bool
	logFile    string
	quiet      bool
)

var (
	compileTimeVersion string
	Version            = bestVersion()
	Commit             string
	Date               string
	BuiltBy            string
)

var manuallyImplementedSystems = map[string]bool{
	"cryptography":  true,
	"configuration": true,
}

type AboutPlatformCLI struct {
	Version string
	Commit  string
	Date    string
	BuiltBy string
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:          "oeco",
	Short:        "Connect to Open Ecosystems",
	Long:         `Allows you to securely and efficiently interact with Open Economic Systems`,
	Version:      Version,
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("debug logging enabled")
		}
	},
}

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

	if err2 := settingsProvider.WatchSettings(); err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	settings := settingsProvider.GetSettings()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "settings", settings)
	RootCmd.SetContext(ctx)

	AddCommands(settings)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCommands(settings *specv2pb.SpecSettings) {

	cmdv2alphapbcmd.CommandRegistry.RegisterCommands()

	if settings != nil && settings.Systems2 != nil {
		for _, system := range settings.Systems2 {
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

			RootCmd.AddCommand(command)
		}
	}

	// Manually add certain system commands
	RootCmd.AddCommand(cryptographyv2alphapbint.SystemCmd)
	RootCmd.AddCommand(configurationv2alphapbint.SystemCmd)
	RootCmd.AddCommand(connectorv2alphatui.Cmd)

}

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

func bestVersion() string {

	if compileTimeVersion != "" {
		return compileTimeVersion
	}
	return DefaultVersion

}
