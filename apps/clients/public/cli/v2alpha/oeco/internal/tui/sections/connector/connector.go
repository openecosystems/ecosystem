//nolint:typecheck,revive
package connector

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	details_page "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/details_page"
	homepage "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/home_page"
	logs_page "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/logs_page"
	requestspage "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/requests_page"
	sections "apps/clients/public/cli/v2alpha/oeco/internal/tui/sections"
)

// Model represents the main state containing a base model, key bindings, and tasks for the application.
type Model struct {
	sections.BaseModel
	keys *keys.KeyMap
}

// NewModel initializes and returns a new instance of the Model with the provided SpecSettings.
func NewModel(pctx *context.ProgramContext) Model {
	m := Model{
		keys: keys.Keys,
	}

	pctx.Section = config.ConnectorSection
	pctx.Page = config.ConnectorDetailsPage

	var pages []contract.Page
	pages = append(pages, homepage.NewModel(pctx))
	pages = append(pages, details_page.NewModel(pctx))
	pages = append(pages, requestspage.NewModel(pctx))
	pages = append(pages, logs_page.NewModel(pctx))
	m.Pages = pages
	m.BaseModel = sections.NewBaseModel(
		pctx,
		sections.NewBaseOptions{
			Singular:      "connector",
			Plural:        "connectors",
			CurrentPageID: m.GetDefaultPageID(),
			Pages:         pages,
		},
	)

	m.Spinner.Style = lipgloss.NewStyle().
		Background(m.Ctx.Theme.SelectedBackground)

	return m
}

// Init initializes the model by batching the BaseModel initialization and enabling the alternative screen mode.
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.InitBase())
}

// Update handles incoming messages, updates the model state, and returns the updated model and command batch.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd       tea.Cmd
		tabsCmd   tea.Cmd
		pageCmd   tea.Cmd
		footerCmd tea.Cmd
		cmds      []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.Up):
		}
	}

	m.Tabs, tabsCmd = m.Tabs.Update(msg)
	switch page := m.CurrentPage.(type) {
	case details_page.Model:
		m.Ctx.Page = config.ConnectorDetailsPage
		m.CurrentPageModel, pageCmd = page.Update(msg)
	case requestspage.Model:
		m.Ctx.Page = config.ConnectorRequestsPage
		m.CurrentPageModel, pageCmd = page.Update(msg)
	case logs_page.Model:
		m.Ctx.Page = config.ConnectorLogsPage
		m.CurrentPageModel, pageCmd = page.Update(msg)
	}
	m.Footer, footerCmd = m.Footer.Update(msg)
	m.UpdateProgramContext(m.Ctx)

	cmds = append(
		cmds,
		cmd,
		tabsCmd,
		pageCmd,
		footerCmd,
	)

	return m, tea.Batch(cmds...)
}

// View renders the current application's view by combining the base layout and the current page's view content.
func (m Model) View() string {
	return m.ViewBase(m.CurrentPage.View())

	// s := m.ViewDebug()
	// s.WriteString(m.CurrentPage.View())
	// return s.String()
}

// GetPages returns the collection of pages managed by the model.
func (m Model) GetPages() []contract.Page {
	return m.Pages
}
