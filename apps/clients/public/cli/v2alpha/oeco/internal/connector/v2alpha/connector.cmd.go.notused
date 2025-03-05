package connectorv2alphatui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	data "apps/clients/public/cli/v2alpha/oeco/internal/data"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	connector "apps/clients/public/cli/v2alpha/oeco/internal/tui/sections/connector"
	tasks "apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
	theme "apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// Cmd defines a CLI command named "connector" for interaction, supporting one optional argument and a short description.
var Cmd = &cobra.Command{
	Use:     "dash",
	Short:   "Ecosystem Dashboard",
	Version: "",
	Args:    cobra.MaximumNArgs(1),
}

// createModel initializes a connector model and optionally sets up logging based on the provided command flags.
func createModel(cmd *cobra.Command) connector.Model {
	settings := cmd.Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)
	logger := cmd.Context().Value(sdkv2alphalib.LoggerContextKey).(*log.Logger)

	c := config.Config{}
	t := theme.ParseTheme(&c)
	pctx := &context.ProgramContext{
		Config:   &c,
		Settings: settings,
		Logger:   logger,
		User:     data.GetUserName(),
		Theme:    t,
		Styles:   theme.InitStyles(t),
	}

	return connector.NewModel(pctx)
}

// init initializes the `Cmd` execution logic by setting up the model, logging, cleanup, and running the TUI program.
func init() {
	Cmd.Run = func(cmd *cobra.Command, _ []string) {
		model := createModel(cmd)

		defer gracefulShutdown()

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
// gracefulShutdown recovers from any panic that occurred and logs the recovery message before quitting the tea program.
func gracefulShutdown() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}

	tasks.Close()
	_ = tea.Quit()
}
