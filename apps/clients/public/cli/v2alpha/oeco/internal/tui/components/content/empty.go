package content

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"

	tea "github.com/charmbracelet/bubbletea"
)

// EmptyModel is a concrete model that embeds BaseModel and implements minimal functionality as a placeholder or default model.
type EmptyModel struct {
	BaseModel
}

// NewEmptyModel initializes and returns a new EmptyModel as a MainContent using the provided ProgramContext.
func NewEmptyModel(ctx *context.ProgramContext) contract.MainContent {
	m := &EmptyModel{}
	m.BaseModel = NewBaseModel(
		ctx,
		NewBaseOptions{},
	)

	return m
}

// Update processes the given message, updates the EmptyModel, and returns the updated model along with a batch of commands.
func (m EmptyModel) Update(msg tea.Msg) (EmptyModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	cmds = append(
		cmds,
		cmd,
	)
	return m, tea.Batch(cmds...)
}
