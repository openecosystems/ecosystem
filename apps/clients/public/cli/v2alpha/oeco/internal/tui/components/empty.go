package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

// EmptyModelConfig defines the configuration for an empty model, including basic metadata like the title of the page.
type EmptyModelConfig struct {
	Title string
}

// EmptyModel represents a UI page with minimal functionality, inheriting behavior and properties from BaseModel.
type EmptyModel struct {
	BaseModel[EmptyModelConfig]
}

// NewEmptyModel creates and initializes an EmptyModel as a Page using a given ProgramContext.
func NewEmptyModel(ctx *context.ProgramContext) contract.Component {
	m := EmptyModel{}

	m.BaseModel = NewBaseModel[EmptyModelConfig](
		ctx,
		NewBaseOptions[EmptyModelConfig]{
			Default: false,
			ComponentConfig: &EmptyModelConfig{
				Title: "Empty",
			},
		},
	)

	return m
}

// Init initializes the EmptyModel and returns a tea.Cmd batch for further processing or updates.
func (m EmptyModel) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes a given message and updates the EmptyModel's state and commands, returning the updated model and a batch of commands.
func (m EmptyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil
		_ = message
	}

	m.UpdateProgramContext(m.Ctx)

	cmds = append(
		cmds,
		cmd,
	)
	return m, tea.Batch(cmds...)
}

// View renders the current state of the `EmptyModel` as a string, displaying "Empty Page" vertically aligned to the left.
func (m EmptyModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, "Empty Page")
}

// UpdateProgramContext updates the ProgramContext for the EmptyModel, syncing any relevant data for rendering or behavior.
func (m EmptyModel) UpdateProgramContext(_ *context.ProgramContext) {}
