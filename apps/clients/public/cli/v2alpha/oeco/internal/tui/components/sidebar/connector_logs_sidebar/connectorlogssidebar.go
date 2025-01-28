package connector_logs_sidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents a structured type embedding sidebar.BaseModel, providing extended functionality for UI components.
type Model struct {
	sidebar.BaseModel
}

// NewModel initializes a new instance of Model using the given ProgramContext.
// It also sets up the BaseModel with default sidebar options.
func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.BaseModel = sidebar.NewBaseModel(
		ctx,
		sidebar.NewBaseOptions{},
	)

	return m
}

// Update updates the Model state based on the given message and returns the updated Model and a command batch.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
