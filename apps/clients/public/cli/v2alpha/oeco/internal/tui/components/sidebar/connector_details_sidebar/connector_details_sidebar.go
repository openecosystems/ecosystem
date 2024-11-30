package connector_details_sidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/connector_form"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

type Model struct {
	sidebar.BaseModel

	form *connector_form.Model
}

func NewModel(ctx *context.ProgramContext, form *connector_form.Model) Model {
	m := Model{
		form: form,
	}
	m.BaseModel = sidebar.NewBaseModel(
		ctx,
		sidebar.NewBaseOptions{
			Opened: true,
		},
	)
	m.Viewport = viewport.New(m.Ctx.SidebarContentWidth, m.Ctx.SidebarContentHeight)

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.Viewport.SetContent(m.form.SidebarView())
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	cmds = append(
		cmds,
		cmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}
