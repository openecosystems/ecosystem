package connectorv2alphatui

import (
	"fmt"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	slog "log"
	"os"
	"strconv"
	"time"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/sections/connector"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "connector",
	Short:   "Connector Interactions",
	Version: "",
	Args:    cobra.MaximumNArgs(1),
}

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

func init() {
	Cmd.Run = func(cmd *cobra.Command, args []string) {
		settings := cmd.Root().Context().Value("settings").(*specv2pb.SpecSettings)

		model, logger := createModel(cmd, settings)

		if logger != nil {
			defer func(logger *os.File) {
				err := logger.Close()
				if err != nil {
				}
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

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
	tea.Quit()
}
