package connector

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/details_page"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/logs_page"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/sections"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"

	requestspage "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/requests_page"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the main state containing a base model, key bindings, and tasks for the application.
type Model struct {
	sections.BaseModel
	keys  *keys.KeyMap
	tasks map[string]context.Task
}

// NewModel initializes and returns a new instance of the Model with the provided SpecSettings.
func NewModel(settings *specv2pb.SpecSettings) Model {
	m := Model{
		keys:  keys.Keys,
		tasks: map[string]context.Task{},
	}

	ctx := &context.ProgramContext{
		Config:   &config.Config{},
		Settings: settings,
		Section:  config.ConnectorSection,
		Page:     config.ConnectorDetailsPage,
		User:     "dimyjeannot",
		StartTask: func(_ context.Task) tea.Cmd {
			return m.Spinner.Tick
		},
	}

	var pages []contract.Page
	pages = append(pages, details_page.NewModel(ctx))
	pages = append(pages, requestspage.NewModel(ctx))
	pages = append(pages, logs_page.NewModel(ctx))
	m.Pages = pages
	m.BaseModel = sections.NewBaseModel(
		ctx,
		sections.NewBaseOptions{
			Singular:      "connector",
			Plural:        "connectors",
			CurrentPageId: m.GetDefaultPageId(),
			Pages:         pages,
		},
	)

	m.Spinner.Style = lipgloss.NewStyle().
		Background(m.Ctx.Theme.SelectedBackground)

	return m
}

// Init initializes the model by batching the BaseModel initialization and enabling the alternative screen mode.
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.InitBase(), tea.EnterAltScreen)
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
		m.CurrentPage, pageCmd = page.Update(msg)
	case requestspage.Model:
		m.Ctx.Page = config.ConnectorRequestsPage
		m.CurrentPage, pageCmd = page.Update(msg)
	case logs_page.Model:
		m.Ctx.Page = config.ConnectorLogsPage
		m.CurrentPage, pageCmd = page.Update(msg)
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
