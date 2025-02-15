package ecosystemcreatesidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	ecosystemcreateform "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	sidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// Model represents a user interface model combining a sidebar and a form within a program context.
type Model struct {
	sidebar.BaseModel

	form *ecosystemcreateform.Model
}

// NewModel creates and initializes a new Model instance with a given program context and connector form configuration.
func NewModel(ctx *context.ProgramContext, form *ecosystemcreateform.Model) *Model {
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

	return &m
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with batched commands.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
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
