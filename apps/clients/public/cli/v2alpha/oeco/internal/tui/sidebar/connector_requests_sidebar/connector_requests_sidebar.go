package connectorrequestssidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	sidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
)

// Model represents a sidebar model that extends BaseModel with additional update and view capabilities.
type Model struct {
	sidebar.BaseModel
}

// NewModel initializes and returns a new Model instance with a sidebar base model configured using the provided context.
func NewModel(ctx *context.ProgramContext) contract.Sidebar {
	m := Model{}
	m.BaseModel = sidebar.NewBaseModel(
		ctx,
		sidebar.NewBaseOptions{
			Viewport: viewport.New(ctx.MainContentWidth, ctx.PageContentHeight),
			Opened:   true,
		},
	)

	return m
}

// Init initializes the EmptyModel by returning a batched command with no specific functionality.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes a message, updates the model's state, and returns the updated model along with a batch of commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent(m.ViewDebug().String())
	v, c := m.Viewport.Update(msg)
	m.Viewport = &v
	viewportCmd = c

	cmds = append(
		cmds,
		cmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the string representation of the BaseModel's current view by delegating to the ViewBase method.
func (m Model) View() string {
	return m.ViewBase()
}
