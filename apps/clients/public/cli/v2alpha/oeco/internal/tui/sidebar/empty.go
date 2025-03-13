package sidebar

import (
	tea "github.com/charmbracelet/bubbletea"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// EmptyModel represents a specialized type that embeds BaseModel with no additional functionality.
type EmptyModel struct {
	*BaseModel
}

// NewEmptyModel creates and initializes a new EmptyModel with a ProgramContext, returning it as a Sidebar implementation.
func NewEmptyModel(ctx *context.ProgramContext) *EmptyModel {
	baseModel := NewBaseModel(
		ctx,
		&NewBaseOptions{},
	)

	m := &EmptyModel{
		BaseModel: baseModel,
	}

	return m
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *EmptyModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds,
		m.InitBase(),
	)

	return tea.Batch(cmds...)
}

// Update handles incoming messages, updates the base model and viewport, and returns the updated model and batched commands.
func (m *EmptyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds        []tea.Cmd
		baseCmd     tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, baseCmd = m.UpdateBase(msg)
	m.Viewport.SetContent(m.ViewDebug().String())
	v, c := m.Viewport.Update(msg)
	m.Viewport = &v
	viewportCmd = c

	cmds = append(
		cmds,
		baseCmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the string representation of the BaseModel's current view by delegating to the ViewBase method.
func (m *EmptyModel) View() string {
	return m.ViewBase()
}
