package connector_requests_sidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

type Model struct {
	sidebar.BaseModel
}

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
