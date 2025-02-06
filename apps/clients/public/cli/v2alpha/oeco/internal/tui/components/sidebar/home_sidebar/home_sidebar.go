package homesidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents a sidebar model that extends BaseModel with additional update and view capabilities.
type Model struct {
	sidebar.BaseModel
}

// NewModel initializes and returns a new Model instance with a sidebar base model configured using the provided context.
func NewModel(ctx *context.ProgramContext) Model {
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

// Update processes a message, updates the model's state, and returns the updated model along with a batch of commands.
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
