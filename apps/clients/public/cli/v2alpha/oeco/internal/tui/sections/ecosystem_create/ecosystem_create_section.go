package ecosystemcreate

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	ecosystemcreatepage "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/ecosystem_create_page"
	ecosystemdashboardpage "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/ecosystem_dashboard_page"
	sections "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/sections"
)

// Model represents the main state containing a base model, key bindings, and tasks for the application.
type Model struct {
	*sections.BaseModel

	keys *keys.KeyMap
}

// NewModel initializes and returns a new instance of the Model with the provided SpecSettings.
func NewModel(pctx *context.ProgramContext) *Model {
	ecosystemCreatePageModel := ecosystemcreatepage.NewModel(pctx)
	ecosystemDashboardPageModel := ecosystemdashboardpage.NewModel(pctx)

	var p []contract.Page
	p = append(p,
		ecosystemCreatePageModel,
		ecosystemDashboardPageModel,
	)

	defaultPageID := 0
	for i, page := range p {
		if page.GetPageSettings().IsDefault {
			defaultPageID = i
			break
		}
	}

	pctx.Section = config.EcosystemSection
	pctx.Page = config.EcosystemCreatePage

	m := &Model{
		keys: keys.Keys,
		BaseModel: sections.NewBaseModel(
			pctx,
			&sections.NewBaseOptions{
				Singular:      "ecosystem",
				Plural:        "ecosystems",
				CurrentPageID: defaultPageID,
				Pages:         p,
			},
		),
	}

	m.Spinner.Style = lipgloss.NewStyle().Background(m.Ctx.Theme.SelectedBackground)

	// m.Ctx.Logger.Debug("Section: Base Model Initial Configuration \n\n", m.ViewDebug().String())
	m.Ctx.Logger.Debug("Section: Ecosystem Create: Initial Configuration", "Current Page ID", m.CurrentPageID, "Current Page", m.CurrentPage.GetPageSettings().Title)

	return m
}

// Init initializes the model by batching the BaseModel initialization and enabling the alternative screen mode.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, page := range m.Pages {
		cmds = append(cmds, page.Init())
	}

	cmds = append(cmds, m.InitBase())

	return tea.Batch(cmds...)
}

// Update handles incoming messages, updates the model state, and returns the updated model and command batch.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		baseCmd tea.Cmd
		cmds    []tea.Cmd
	)

	m.BaseModel, baseCmd = m.UpdateBase(msg)

	switch page := m.CurrentPage.(type) {
	case *ecosystemcreatepage.Model:
		m.Ctx.Logger.Debug("Section: Ecosystem Create: Page Update")
		m.Ctx.Page = config.EcosystemCreatePage
		_ = page
	case *ecosystemdashboardpage.Model:
		m.Ctx.Logger.Debug("Section: Ecosystem Dashboard: Page Update")
		m.Ctx.Page = config.EcosystemDashboardPage
	}
	m.UpdateProgramContext(m.Ctx)

	cmds = append(
		cmds,
		baseCmd,
	)

	return m, tea.Batch(cmds...)
}

// View renders the current application's view by combining the base layout and the current page's view content.
func (m *Model) View() string {
	switch page := m.CurrentPage.(type) {
	case *ecosystemcreatepage.Model:
		return m.ViewBase(page.View())
	case *ecosystemdashboardpage.Model:
		return m.ViewBase(page.View())
	default:
		return ""
	}

	// s := m.ViewDebug()
	// s.WriteString("")
	// return s.String()
}

// GetPages returns the collection of pages managed by the model.
func (m *Model) GetPages() []contract.Page {
	return m.Pages
}
