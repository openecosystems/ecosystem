package homecontent

import (
	tea "github.com/charmbracelet/bubbletea"

	content "apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

// Model represents a UI component that extends BaseModel for handling viewport rendering and program context updates.
type Model struct {
	content.BaseModel
}

// NewModel initializes and returns a new Model instance using the provided ProgramContext.
func NewModel(ctx *context.ProgramContext) contract.MainContent {
	m := Model{}
	m.BaseModel = content.NewBaseModel(
		ctx,
		content.NewBaseOptions{},
	)

	return m
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the given message, updates the model state, and returns a new model along with a batch of commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent("Home")
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	cmds = append(
		cmds,
		cmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m Model) View() string {
	return m.ViewBase()
}
