package ecosystemv2alphapbint

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	data "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/data"
	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	ecosystemcreate "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/sections/ecosystem_create"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// Cmd defines a CLI command named "create" for interaction, supporting one optional argument and a short description.
var Cmd = &cobra.Command{
	Use:     "create",
	Short:   "Ecosystem Create",
	Version: "",
	Args:    cobra.MaximumNArgs(1),
}

// createModel initializes an ecosystem model and optionally sets up logging based on the provided command flags.
func createModel(cmd *cobra.Command) *ecosystemcreate.Model {
	settings := cmd.Context().Value(sdkv2alphalib.SettingsContextKey).(*sdkv2alphalib.CLIConfiguration)
	logger := cmd.Context().Value(sdkv2alphalib.LoggerContextKey).(*log.Logger)

	// c := config.Config{}
	c, err := config.ParseConfig()
	if err != nil {
		logger.Fatal("error parsing config: ", err, "")
	}

	t := theme.ParseTheme(&c)
	pctx := &context.ProgramContext{
		Config:   &c,
		Settings: settings,
		Logger:   logger,
		User:     data.GetUserName(),
		Theme:    t,
		Styles:   theme.InitStyles(t),
	}

	return ecosystemcreate.NewModel(pctx)
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
