package ecosystemdashboardpage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/key"
	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	ecosystemdashboardcontent "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/content/ecosystem_dashboard_content"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	packetssidebar "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar/packets_sidebar"
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
	s := packetssidebar.NewModel(pctx)

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
	)

	return tea.Batch(cmds...)
}

// Update processes the given message, updates the model's components, and returns the updated model and a batch of commands.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		baseCmd tea.Cmd
		cmds    []tea.Cmd
		cmd     tea.Cmd
	)

	m.BaseModel, baseCmd = m.UpdateBase(msg)
	m.UpdateProgramContext(m.Ctx)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Logger.Debug("Key pressed", "key", message.String())
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.Down):
			// prevRow := m.CurrentMainContent.CurrRow()
			// nextRow := m.Table.NextRow()
			cmd = m.OnViewedRowChanged()
		}
	}

	cmds = append(
		cmds,
		baseCmd,
		cmd,
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

// OnViewedRowChanged synchronizes the sidebar and scrolls it to the top whenever the viewed row in the UI changes.
func (m *Model) OnViewedRowChanged() tea.Cmd {
	m.CurrentSidebar.SyncSidebar()
	m.CurrentSidebar.ScrollToTop()
	return tea.Batch()
}
