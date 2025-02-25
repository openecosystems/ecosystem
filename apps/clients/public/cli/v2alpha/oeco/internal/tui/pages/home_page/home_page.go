package homepage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/content/home_content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	homesidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar/home_sidebar"
)

// ModelConfig represents the configuration settings for a model in the application.
type ModelConfig struct{}

// Model represents a page model that extends BaseModel with main content and sidebar integration functionality.
type Model struct {
	pages.BaseModel[ModelConfig]

	mainContent homecontent.Model
	sidebar     homesidebar.Model
}

// NewModel initializes and returns a new Model instance configured with a ProgramContext and default settings.
func NewModel(ctx *context.ProgramContext) Model {
	c := ModelConfig{}

	p := Model{
		mainContent: homecontent.NewModel(ctx),
		sidebar:     homesidebar.NewModel(ctx),
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
		Title:     "Home",
		IsDefault: false,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.ConnectorRequestsPage,
	}
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the received message to update the model's state and returns the updated model along with any commands.
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

// View renders the combined view of the main content and the sidebar using the base view logic.
func (m Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.mainContent.View(),
		m.sidebar.View(),
	))
}
