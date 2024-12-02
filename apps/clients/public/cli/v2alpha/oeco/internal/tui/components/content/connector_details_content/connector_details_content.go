package connector_details_content

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/connector_form"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

var introduction = `
This will help you explore the different connectors we support using mock-data based prototyping:
These connectors are fake/mock versions that are identical in structure,
but whose actual values are randomized and synthetically generated.
This allows you to prototype and test before doing mesh execution.
`

type Model struct {
	content.BaseModel

	form             *connector_form.Model
	markdownRenderer glamour.TermRenderer
	introduction     string
}

func NewModel(ctx *context.ProgramContext, form *connector_form.Model) Model {
	m := Model{
		form: form,
	}
	m.BaseModel = content.NewBaseModel(
		ctx,
		content.NewBaseOptions{},
	)
	m.Viewport = viewport.New(m.Ctx.MainContentBodyWidth, m.Ctx.MainContentBodyHeight)
	m.markdownRenderer = markdown.GetMarkdownRenderer(80)

	var err error
	m.introduction, err = m.markdownRenderer.Render(introduction)
	if err != nil {
		panic(err)
	}

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		formCmd     tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	f, formCmd := m.form.Update(msg)
	m.form = &f
	m.Viewport.SetContent(m.introduction + m.form.View())
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	cmds = append(
		cmds,
		cmd,
		formCmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}
