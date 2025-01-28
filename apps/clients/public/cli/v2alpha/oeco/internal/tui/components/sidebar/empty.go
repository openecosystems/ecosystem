package sidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"

	tea "github.com/charmbracelet/bubbletea"
)

// EmptyModel represents a specialized type that embeds BaseModel with no additional functionality.
type EmptyModel struct {
	BaseModel
}

// NewEmptyModel creates and initializes a new EmptyModel with a ProgramContext, returning it as a Sidebar implementation.
func NewEmptyModel(ctx *context.ProgramContext) contract.Sidebar {
	m := &EmptyModel{}
	m.BaseModel = NewBaseModel(
		ctx,
		NewBaseOptions{},
	)

	return m
}

// Update handles incoming messages, updates the base model and viewport, and returns the updated model and batched commands.
func (m EmptyModel) Update(msg tea.Msg) (EmptyModel, tea.Cmd) {
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
