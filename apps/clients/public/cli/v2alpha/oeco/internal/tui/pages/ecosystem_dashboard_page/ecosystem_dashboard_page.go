package ecosystemdashboardpage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	ecosystemdashboardcontent "apps/clients/public/cli/v2alpha/oeco/internal/tui/content/ecosystem_dashboard_content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	ecosystemdashboardsidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar/ecosystem_dashboard_sidebar"
)

// ModelConfig represents the configuration structure for initializing and customizing a model instance.
type ModelConfig struct{}

// Model represents an aggregate model combining base functionality, form element, content, and a sidebar component.
type Model struct {
	*pages.BaseModel
}

// NewModel initializes and returns a new Model instance using a ProgramContext.
// It sets up a base model, form, main content, and sidebar with dependencies configured from the provided context.
func NewModel(pctx *context.ProgramContext) contract.Page {
	m := ecosystemdashboardcontent.NewModel(pctx)
	s := ecosystemdashboardsidebar.NewModel(pctx)

	baseModel := pages.NewBaseModel(
		pctx,
		&pages.NewBaseOptions{
			Default:            true,
			Keys:               keys.Keys,
			KeyBindings:        nil,
			CurrentMainContent: m,
			CurrentSidebar:     s,
			PageSettings: &contract.PageSettings{
				Title:     "Ecosystem Dashboard",
				IsDefault: true,
				// KeyBindings:   GetKeyBindings(),
				ContentHeight: 0,
				Type:          config.EcosystemDashboardPage,
			},
		},
	)

	p := &Model{
		BaseModel: baseModel,
	}

	pctx.Logger.Debug("Page - Ecosystem Dashboard Page: Configured initial model")

	return p
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds,
		m.InitBase(),
		m.CurrentMainContent.Init(),
		m.CurrentSidebar.Init(),
	)

	return tea.Batch(cmds...)
}

// Update processes the given message, updates the model's components, and returns the updated model and a batch of commands.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		baseCmd tea.Cmd
		// mainContentCmd tea.Cmd
		// sidebarCmd     tea.Cmd
		cmds []tea.Cmd
	)

	m.BaseModel, baseCmd = m.UpdateBase(msg)
	m.UpdateProgramContext(m.Ctx)
	//_, mainContentCmd = m.CurrentMainContent.Update(msg)
	//_, sidebarCmd = m.CurrentSidebar.Update(msg)

	cmds = append(
		cmds,
		baseCmd,
		// mainContentCmd,
		// sidebarCmd,
	)

	return m, tea.Batch(cmds...)
}

// View renders the combined view of the main content and sidebar by arranging them horizontally with a base style.
func (m *Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.CurrentMainContent.View(),
		m.CurrentSidebar.View(),
	))

	// s := m.ViewDebug()
	// s.WriteString(m.mainContent.View())
	// s.WriteString("\n")
	// s.WriteString(m.sidebar.View())
	// s.WriteString("\n")
	// return s.String()
}
