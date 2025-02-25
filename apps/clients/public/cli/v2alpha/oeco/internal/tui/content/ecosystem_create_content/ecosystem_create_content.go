package ecosystemcreatecontent

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"

	ecosystemcreateform "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form/ecosystem_create_form"
	markdown "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/markdown"
	content "apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

var introduction = `
This will help you explore the different connectors we support using mock-data based prototyping:
These connectors are fake/mock versions that are identical in structure,
but whose actual values are randomized and synthetically generated.
This allows you to prototype and test before doing mesh execution.
`

// Model represents a UI component that combines markdown rendering, a form interface, and contextual data handling.
type Model struct {
	content.BaseModel

	form             *ecosystemcreateform.Model
	formModel        tea.Model
	markdownRenderer glamour.TermRenderer
	introduction     string
}

// NewModel initializes and returns a new Model instance using the provided program context and connector form model.
// It sets up internal properties, including the base model, viewport, and a markdown renderer for introduction rendering.
func NewModel(pctx *context.ProgramContext, form *ecosystemcreateform.Model) contract.MainContent {
	m := Model{
		form:             form,
		markdownRenderer: markdown.GetMarkdownRenderer(80),
		BaseModel: content.NewBaseModel(
			pctx,
			content.NewBaseOptions{
				Viewport: viewport.New(pctx.MainContentBodyWidth, pctx.MainContentBodyHeight),
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
func (m Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update handles incoming messages to update the model's state and returns the updated model along with a command batch.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		formCmd     tea.Cmd
		viewportCmd tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	m.formModel, formCmd = m.form.Update(msg)
	m.Viewport.SetContent(m.introduction + m.form.View())
	m.Viewport, viewportCmd = m.Viewport.Update(msg)

	m.Ctx.Logger.Debug("Content - Ecosystem Create Content - Update: Set model")

	cmds = append(
		cmds,
		cmd,
		formCmd,
		viewportCmd,
	)
	return m, tea.Batch(cmds...)
}

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m Model) View() string {
	return m.ViewBase()
}
