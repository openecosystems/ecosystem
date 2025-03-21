package ecosystemcreatesidebar

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	ecosystemcreateform "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	sidebar "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
)

// Model represents a user interface model combining a sidebar and a form within a program context.
type Model struct {
	*sidebar.BaseModel

	form *ecosystemcreateform.Model
}

// NewModel creates and initializes a new Model instance with a given program context and connector form configuration.
func NewModel(ctx *context.ProgramContext, form *ecosystemcreateform.Model) *Model {
	v := viewport.New(ctx.SidebarContentWidth, ctx.SidebarContentHeight)

	baseModel := sidebar.NewBaseModel(
		ctx,
		&sidebar.NewBaseOptions{
			Opened:   true,
			Viewport: &v,
		},
	)

	m := &Model{
		BaseModel: baseModel,
		form:      form,
	}

	return m
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds,
		m.InitBase(),
	)

	return tea.Batch(cmds...)
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with batched commands.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds        []tea.Cmd
		baseCmd     tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, baseCmd = m.UpdateBase(msg)
	m.Viewport.SetContent(m.form.SidebarView())
	v, c := m.Viewport.Update(msg)
	m.Viewport = &v
	viewportCmd = c

	cmds = append(
		cmds,
		baseCmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the string representation of the BaseModel's current view by delegating to the ViewBase method.
func (m *Model) View() string {
	return m.ViewBase()
}
