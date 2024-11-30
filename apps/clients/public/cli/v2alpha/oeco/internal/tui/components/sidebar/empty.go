package sidebar

import (
	"github.com/charmbracelet/bubbletea"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

type EmptyModel struct {
	BaseModel
}

func NewEmptyModel(ctx *context.ProgramContext) contract.Sidebar {

	m := &EmptyModel{}
	m.BaseModel = NewBaseModel(
		ctx,
		NewBaseOptions{},
	)

	return m

}

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
