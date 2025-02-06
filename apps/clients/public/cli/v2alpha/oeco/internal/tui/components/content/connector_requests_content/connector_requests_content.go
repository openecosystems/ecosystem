package connectorrequestscontent

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents a UI component that extends BaseModel for handling viewport rendering and program context updates.
type Model struct {
	content.BaseModel
}

// NewModel initializes and returns a new Model instance using the provided ProgramContext.
func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.BaseModel = content.NewBaseModel(
		ctx,
		content.NewBaseOptions{},
	)

	return m
}

// Update processes the given message, updates the model state, and returns a new model along with a batch of commands.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent("Requests")
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	cmds = append(
		cmds,
		cmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}
