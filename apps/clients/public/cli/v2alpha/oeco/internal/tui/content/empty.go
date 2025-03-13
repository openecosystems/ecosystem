package content

import (
	tea "github.com/charmbracelet/bubbletea"

	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// EmptyModel is a concrete model that embeds BaseModel and implements minimal functionality as a placeholder or default model.
type EmptyModel struct {
	*BaseModel
}

// NewEmptyModel initializes and returns a new EmptyModel as a MainContent using the provided ProgramContext.
func NewEmptyModel(pctx *context.ProgramContext) *EmptyModel {
	baseModel := NewBaseModel(
		pctx,
		&NewBaseOptions{},
	)

	m := &EmptyModel{
		BaseModel: baseModel,
	}

	return m
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *EmptyModel) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the given message, updates the EmptyModel, and returns the updated model along with a batch of commands.
func (m *EmptyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m *EmptyModel) View() string {
	return m.ViewBase()
}
