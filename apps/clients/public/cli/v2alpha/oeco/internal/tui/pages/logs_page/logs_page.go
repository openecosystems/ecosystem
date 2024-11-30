package logs_page

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content/connector_logs_content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar/connector_logs_sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
)

type ModelConfig struct{}

type Model struct {
	pages.BaseModel[ModelConfig]

	mainContent connector_logs_content.Model
	sidebar     connector_logs_sidebar.Model
}

func NewModel(ctx *context.ProgramContext) Model {
	c := ModelConfig{}

	p := Model{
		mainContent: connector_logs_content.NewModel(ctx),
		sidebar:     connector_logs_sidebar.NewModel(ctx),
	}

	p.BaseModel = pages.NewBaseModel[ModelConfig](
		ctx,
		pages.NewBaseOptions[ModelConfig]{
			Default:     false,
			PageConfig:  &c,
			Keys:        keys.Keys,
			KeyBindings: nil,
		},
	)

	return p
}

func (m Model) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:     "Connector Logs",
		IsDefault: false,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.ConnectorLogsPage,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd            tea.Cmd
		mainContentCmd tea.Cmd
		sidebarCmd     tea.Cmd
		cmds           []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	m.UpdateProgramContext(m.Ctx)
	m.mainContent, mainContentCmd = m.mainContent.Update(msg)
	m.sidebar, sidebarCmd = m.sidebar.Update(msg)

	cmds = append(
		cmds,
		cmd,
		mainContentCmd,
		sidebarCmd,
	)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.mainContent.View(),
		m.sidebar.View(),
	))
}
