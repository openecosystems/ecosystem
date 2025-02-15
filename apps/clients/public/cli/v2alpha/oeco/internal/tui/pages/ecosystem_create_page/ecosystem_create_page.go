package ecosystemcreatepage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	ecosystemcreatecontent "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content/ecosystem_create_content"
	ecosystemcreateform "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	ecosystemcreatesidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar/ecosystem_create_sidebar"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
)

// ModelConfig represents the configuration structure for initializing and customizing a model instance.
type ModelConfig struct{}

// Model represents an aggregate model combining base functionality, form element, content, and a sidebar component.
type Model struct {
	pages.BaseModel[ModelConfig]

	form        *ecosystemcreateform.Model
	mainContent *ecosystemcreatecontent.Model
	sidebar     *ecosystemcreatesidebar.Model
}

// NewModel initializes and returns a new Model instance using a ProgramContext.
// It sets up a base model, form, main content, and sidebar with dependencies configured from the provided context.
func NewModel(ctx *context.ProgramContext) *Model {
	c := ModelConfig{}

	f := ecosystemcreateform.NewModel(ctx)
	p := Model{
		form: &f,
	}
	p.mainContent = ecosystemcreatecontent.NewModel(ctx, &f)
	p.sidebar = ecosystemcreatesidebar.NewModel(ctx, &f)

	p.BaseModel = pages.NewBaseModel[ModelConfig](
		ctx,
		pages.NewBaseOptions[ModelConfig]{
			Default:     true,
			PageConfig:  &c,
			Keys:        keys.Keys,
			KeyBindings: nil,
		},
	)

	return &p
}

// GetPageSettings returns the settings for the "Create an Ecosystems" page, including title, default status, and page type.
func (m *Model) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:     "Create an Ecosystem",
		IsDefault: true,
		// KeyBindings:   GetKeyBindings(),
		ContentHeight: 0,
		Type:          config.EcosystemCreatePage,
	}
}

// Update processes the given message, updates the model's components, and returns the updated model and a batch of commands.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
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

// View renders the combined view of the main content and sidebar by arranging them horizontally with a base style.
func (m *Model) View() string {
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
