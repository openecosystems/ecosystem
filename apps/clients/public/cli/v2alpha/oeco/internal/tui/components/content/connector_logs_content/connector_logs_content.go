package connector_logs_content

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents a container structure embedding BaseModel for managing viewport updates and contextual program state.
type Model struct {
	content.BaseModel
}

// NewModel creates a new instance of Model with a base model initialized using the provided ProgramContext.
func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.BaseModel = content.NewBaseModel(
		ctx,
		content.NewBaseOptions{},
	)

	return m
}

// Update processes a given message, updates the model's state, and returns the updated model along with a command batch.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent("Logs")
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	cmds = append(
		cmds,
		cmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}
