package logs_page

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content/connector_logs_content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar/connector_logs_sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ModelConfig represents the configuration type used to customize the behavior and properties of a specific model.
type ModelConfig struct{}

// Model represents a page structure, embedding BaseModel with specific main content and sidebar models for UI rendering.
type Model struct {
	pages.BaseModel[ModelConfig]

	mainContent connector_logs_content.Model
	sidebar     connector_logs_sidebar.Model
}

// NewModel initializes and returns a new Model instance configured with the given ProgramContext.
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

// GetPageSettings returns the page configuration settings, including title, default status, content height, and page type.
func (m Model) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:     "Connector Logs",
		IsDefault: false,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.ConnectorLogsPage,
	}
}

// Update processes a given message, updates the model state, and returns the updated model along with any commands.
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

// View generates and returns the composed string representation of the model, combining main content and sidebar views.
func (m Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.mainContent.View(),
		m.sidebar.View(),
	))
}
