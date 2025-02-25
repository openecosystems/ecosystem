package connectorlogssidebar

import (
	tea "github.com/charmbracelet/bubbletea"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	sidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
)

// Model represents a structured type embedding sidebar.BaseModel, providing extended functionality for UI components.
type Model struct {
	sidebar.BaseModel
}

// NewModel initializes a new instance of Model using the given ProgramContext.
// It also sets up the BaseModel with default sidebar options.
func NewModel(ctx *context.ProgramContext) contract.Sidebar {
	m := Model{}
	m.BaseModel = sidebar.NewBaseModel(
		ctx,
		sidebar.NewBaseOptions{},
	)

	return m
}

// Init initializes the EmptyModel by returning a batched command with no specific functionality.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update updates the Model state based on the given message and returns the updated Model and a command batch.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent(m.ViewDebug().String())
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

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
