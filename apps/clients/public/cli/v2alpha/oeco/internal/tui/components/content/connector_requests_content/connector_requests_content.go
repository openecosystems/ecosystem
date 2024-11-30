package connector_requests_content

import (
	"github.com/charmbracelet/bubbletea"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

type Model struct {
	content.BaseModel
}

func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.BaseModel = content.NewBaseModel(
		ctx,
		content.NewBaseOptions{},
	)

	return m
}

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
