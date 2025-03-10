package prompt

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// Model represents a UI component managing text input and application context.
type Model struct {
	ctx    *context.ProgramContext
	prompt *textinput.Model
}

// NewModel initializes and returns a new instance of Model with the provided ProgramContext and a focused text input.
func NewModel(ctx *context.ProgramContext) *Model {
	ti := textinput.New()
	ti.Focus()
	ti.Blur()
	ti.CursorStart()

	return &Model{
		ctx:    ctx,
		prompt: &ti,
	}
}

// Update handles the update cycle for the Model, updating its prompt state and returning the updated Model and command.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	prompt, cmd := m.prompt.Update(msg)
	m.prompt = &prompt
	return m, cmd
}

// View renders the prompt view as a string representing the current state of the text input model.
func (m *Model) View() string {
	return m.prompt.View()
}

// Init initializes the model and returns a command for blinking text input.
func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

// Blur disables focus for the prompt component within the Model, making it inactive for user input.
func (m *Model) Blur() {
	m.prompt.Blur()
}

// Focus sets the text input focus to the associated prompt and returns the command to apply it.
func (m *Model) Focus() tea.Cmd {
	return m.prompt.Focus()
}

// SetValue sets the value of the underlying prompt to the provided string.
func (m *Model) SetValue(value string) {
	m.prompt.SetValue(value)
}

// Value retrieves the current value from the text input model associated with the Model instance.
func (m *Model) Value() string {
	return m.prompt.Value()
}

// SetPrompt sets the prompt text for the model's text input.
func (m *Model) SetPrompt(prompt string) {
	m.prompt.Prompt = prompt
}

// Reset clears the input model state, restoring it to its initial form.
func (m *Model) Reset() {
	m.prompt.Reset()
}

// UpdateProgramContext updates the program context of the model to the provided context instance.
func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
