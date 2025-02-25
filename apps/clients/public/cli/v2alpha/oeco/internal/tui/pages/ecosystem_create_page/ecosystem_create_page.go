package ecosystemcreatepage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	ecosystemcreateform "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	ecosystemcreatecontent "apps/clients/public/cli/v2alpha/oeco/internal/tui/content/ecosystem_create_content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	ecosystemcreatesidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar/ecosystem_create_sidebar"
)

// ModelConfig represents the configuration structure for initializing and customizing a model instance.
type ModelConfig struct{}

// Model represents an aggregate model combining base functionality, form element, content, and a sidebar component.
type Model struct {
	pages.BaseModel[ModelConfig]

	form             contract.Component
	mainContent      contract.MainContent
	mainContentModel tea.Model
	sidebar          contract.Sidebar
	sidebarModel     tea.Model
}

// NewModel initializes and returns a new Model instance using a ProgramContext.
// It sets up a base model, form, main content, and sidebar with dependencies configured from the provided context.
func NewModel(pctx *context.ProgramContext) contract.Page {
	f := ecosystemcreateform.NewModel(pctx).(ecosystemcreateform.Model)
	m := ecosystemcreatecontent.NewModel(pctx, &f)
	s := ecosystemcreatesidebar.NewModel(pctx, &f)

	p := Model{
		form:        f,
		mainContent: m,
		sidebar:     s,
		BaseModel: pages.NewBaseModel[ModelConfig](
			pctx,
			pages.NewBaseOptions[ModelConfig]{
				Default:     true,
				PageConfig:  &ModelConfig{},
				Keys:        keys.Keys,
				KeyBindings: nil,
			},
		),
	}

	pctx.Logger.Debug("Page - Ecosystem Create Page: Configured initial model")

	return p
}

// GetPageSettings returns the settings for the "Create an Ecosystems" page, including title, default status, and page type.
func (m Model) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:     "Create an Ecosystem",
		IsDefault: true,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.EcosystemCreatePage,
	}
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the given message, updates the model's components, and returns the updated model and a batch of commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd            tea.Cmd
		mainContentCmd tea.Cmd
		sidebarCmd     tea.Cmd
		cmds           []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	m.UpdateProgramContext(m.Ctx)
	m.mainContentModel, mainContentCmd = m.mainContent.Update(msg)
	m.sidebarModel, sidebarCmd = m.sidebar.Update(msg)

	cmds = append(
		cmds,
		cmd,
		mainContentCmd,
		sidebarCmd,
	)

	return m, tea.Batch(cmds...)
}

// View renders the combined view of the main content and sidebar by arranging them horizontally with a base style.
func (m Model) View() string {
	return m.ViewBase(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.mainContent.View(),
		m.sidebar.View(),
	))

	// s := m.ViewDebug()
	// s.WriteString(m.mainContent.View())
	// s.WriteString("\n")
	// s.WriteString(m.sidebar.View())
	// s.WriteString("\n")
	// return s.String()
}
