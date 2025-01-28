package connectorv2alphatui

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/sections/connector"
	"fmt"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	slog "log"
	"os"
	"strconv"
	"time"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// Cmd defines a CLI command named "connector" for interaction, supporting one optional argument and a short description.
var Cmd = &cobra.Command{
	Use:     "connector",
	Short:   "Connector Interactions",
	Version: "",
	Args:    cobra.MaximumNArgs(1),
}

// createModel initializes a connector model and optionally sets up logging based on the provided command flags.
func createModel(cmd *cobra.Command, settings *specv2pb.SpecSettings) (connector.Model, *os.File) {
	var loggerFile *os.File

	d := cmd.Flag("debug").Value.String()
	debug, _ := strconv.ParseBool(d)

	if debug {
		var fileErr error
		newConfigFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if fileErr == nil {
			log.SetOutput(newConfigFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging to debug.log")
		} else {
			loggerFile, _ = tea.LogToFile("debug.log", "debug")
			slog.Print("Failed setting up logging", fileErr)
		}
	} else {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.FatalLevel)
	}

	return connector.NewModel(settings), loggerFile
}

// init initializes the `Cmd` execution logic by setting up the model, logging, cleanup, and running the TUI program.
func init() {
	Cmd.Run = func(cmd *cobra.Command, _ []string) {
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*specv2pb.SpecSettings)

		model, logger := createModel(cmd, settings)

		if logger != nil {
			defer func(logger *os.File) {
				_ = logger.Close()
			}(logger)
		}

		defer cleanup()

		p := tea.NewProgram(
			model,
			tea.WithAltScreen(),
			tea.WithReportFocus(),
			// tea.WithMouseCellMotion(),
		)

		if _, err := p.Run(); err != nil {
			log.Fatal("Failed starting the TUI", err)
		}
	}
}

// cleanup recovers from any panic that occurred and logs the recovery message before quitting the tea program.
func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
	_ = tea.Quit()
}
