package ecosystemcreatecontent

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"

	ecosystemcreateform "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	markdown "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	content "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
)

var introduction = `
This will help you explore the different connectors we support using mock-data based prototyping:
These connectors are fake/mock versions that are identical in structure,
but whose actual values are randomized and synthetically generated.
This allows you to prototype and test before doing mesh execution.
`

// Model represents a UI component that combines markdown rendering, a form interface, and contextual data handling.
type Model struct {
	*content.BaseModel

	form             *ecosystemcreateform.Model
	formModel        *tea.Model
	markdownRenderer *glamour.TermRenderer
	introduction     string
}

// NewModel initializes and returns a new Model instance using the provided program context and connector form model.
// It sets up internal properties, including the base model, viewport, and a markdown renderer for introduction rendering.
func NewModel(pctx *context.ProgramContext, form *ecosystemcreateform.Model) *Model {
	viewportModel := viewport.New(pctx.MainContentBodyWidth, pctx.MainContentBodyHeight)
	markdownRendererModel := markdown.GetMarkdownRenderer(theme.MainContentMarkdownWidth)

	m := &Model{
		form:             form,
		markdownRenderer: &markdownRendererModel,
		BaseModel: content.NewBaseModel(
			pctx,
			&content.NewBaseOptions{
				Viewport: &viewportModel,
			},
		),
	}

	var err error
	m.introduction, err = m.markdownRenderer.Render(introduction)
	if err != nil {
		panic(err)
	}

	m.Ctx.Logger.Debug("Content - Ecosystem Create Content: Initial Configuration")

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

// Update handles incoming messages to update the model's state and returns the updated model along with a command batch.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		formCmd     tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	f, formCmd := m.form.Update(msg)
	m.formModel = &f

	m.Viewport.SetContent(m.introduction + m.form.View())
	v, c := m.Viewport.Update(msg)
	m.Viewport = &v
	viewportCmd = c

	cmds = append(
		cmds,
		cmd,
		formCmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m *Model) View() string {
	return m.ViewBase()
}
