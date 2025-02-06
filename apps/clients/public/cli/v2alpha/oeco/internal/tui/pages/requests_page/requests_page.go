package requestspage

import (
	connectorrequestscontent "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content/connector_requests_content"
	connectorrequestssidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar/connector_requests_sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ModelConfig represents the configuration settings for a model in the application.
type ModelConfig struct{}

// Model represents a page model that extends BaseModel with main content and sidebar integration functionality.
type Model struct {
	pages.BaseModel[ModelConfig]

	mainContent connectorrequestscontent.Model
	sidebar     connectorrequestssidebar.Model
}

// NewModel initializes and returns a new Model instance configured with a ProgramContext and default settings.
func NewModel(ctx *context.ProgramContext) Model {
	c := ModelConfig{}

	p := Model{
		mainContent: connectorrequestscontent.NewModel(ctx),
		sidebar:     connectorrequestssidebar.NewModel(ctx),
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

// GetPageSettings returns the configuration and metadata for the page, including title, type, default status, and dimensions.
func (m Model) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:     "Connector Requests",
		IsDefault: false,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.ConnectorRequestsPage,
	}
}

// Update processes the received message to update the model's state and returns the updated model along with any commands.
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

// View renders the combined view of the main content and the sidebar using the base view logic.
func (m Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.mainContent.View(),
		m.sidebar.View(),
	))
}
