package connector

import (
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/details_page"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/logs_page"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages/requests_page"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/sections"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	sections.BaseModel
	keys  *keys.KeyMap
	tasks map[string]context.Task
}

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
		StartTask: func(task context.Task) tea.Cmd {
			return m.Spinner.Tick
		},
	}

	var pages []contract.Page
	pages = append(pages, details_page.NewModel(ctx))
	pages = append(pages, requests_page.NewModel(ctx))
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

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.InitBase(), tea.EnterAltScreen)
}

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
	case requests_page.Model:
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

func (m Model) View() string {
	return m.ViewBase(m.CurrentPage.View())

	// s := m.ViewDebug()
	// s.WriteString(m.CurrentPage.View())
	// return s.String()
}

func (m Model) GetPages() []contract.Page {
	return m.Pages
}
