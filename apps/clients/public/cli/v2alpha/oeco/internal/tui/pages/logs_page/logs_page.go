package logspage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	connectorlogscontent "apps/clients/public/cli/v2alpha/oeco/internal/tui/content/connector_logs_content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	connectorlogssidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar/connector_logs_sidebar"
)

// ModelConfig represents the configuration type used to customize the behavior and properties of a specific model.
type ModelConfig struct{}

// Model represents a page structure, embedding BaseModel with specific main content and sidebar models for UI rendering.
type Model struct {
	pages.BaseModel[ModelConfig]

	mainContent connectorlogscontent.Model
	sidebar     connectorlogssidebar.Model
}

// NewModel initializes and returns a new Model instance configured with the given ProgramContext.
func NewModel(ctx *context.ProgramContext) Model {
	c := ModelConfig{}

	p := Model{
		mainContent: connectorlogscontent.NewModel(ctx),
		sidebar:     connectorlogssidebar.NewModel(ctx),
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

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes a given message, updates the model state, and returns the updated model along with any commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
