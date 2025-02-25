package ecosystemcreatesidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	ecosystemcreateform "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	sidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
)

// Model represents a user interface model combining a sidebar and a form within a program context.
type Model struct {
	sidebar.BaseModel

	form *ecosystemcreateform.Model
}

// NewModel creates and initializes a new Model instance with a given program context and connector form configuration.
func NewModel(ctx *context.ProgramContext, form *ecosystemcreateform.Model) contract.Sidebar {
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

// Init initializes the EmptyModel by returning a batched command with no specific functionality.
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with batched commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

// View returns the string representation of the BaseModel's current view by delegating to the ViewBase method.
func (m Model) View() string {
	return m.ViewBase()
}
