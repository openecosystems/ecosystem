package search

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// Model represents the main state for managing a text input field within the UI, paired with a program context.
type Model struct {
	ctx          *context.ProgramContext
	initialValue string
	textInput    textinput.Model
}

// Options defines the configuration for initializing a search input field, including prefix, value, and placeholder text.
type Options struct {
	Prefix       string
	InitialValue string
	Placeholder  string
}

// NewModel creates and returns a new Model instance, initializing the text input with specified options and program context.
func NewModel(ctx *context.ProgramContext, opts Options) Model {
	prompt := fmt.Sprintf(" %s ", opts.Prefix)
	ti := textinput.New()
	ti.Placeholder = opts.Placeholder
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(ctx.Theme.FaintText)
	ti.Width = ctx.MainContentWidth - lipgloss.Width(prompt) - 6
	ti.PromptStyle = ti.PromptStyle.Foreground(ctx.Theme.SecondaryText)
	ti.Prompt = prompt
	ti.TextStyle = ti.TextStyle.Faint(true)
	ti.Cursor.Style = ti.Cursor.Style.Faint(true)
	ti.Cursor.TextStyle = ti.Cursor.TextStyle.Faint(true)
	ti.Blur()
	ti.SetValue(opts.InitialValue)
	ti.CursorStart()

	return Model{
		ctx:          ctx,
		textInput:    ti,
		initialValue: opts.InitialValue,
	}
}

// Init initializes the model and returns a command to enable text input blinking.
func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update processes the given message, updates the textInput field, and returns the updated model and any resulting command.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View returns a string representation of the model's text input styled with a border and adjusted widths based on context.
func (m *Model) View(ctx *context.ProgramContext) string {
	return lipgloss.NewStyle().
		Width(ctx.MainContentWidth - 4).
		MaxHeight(3).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.ctx.Theme.PrimaryBorder).
		Render(m.textInput.View())
}

// Focus adjusts the text input's style, places the cursor at the end, and sets the input to focused state.
func (m *Model) Focus() {
	m.textInput.TextStyle = m.textInput.TextStyle.Faint(false)
	m.textInput.CursorEnd()
	m.textInput.Focus()
}

// Blur removes focus from the text input, makes its text appear faint, and resets the cursor to the starting position.
func (m *Model) Blur() {
	m.textInput.TextStyle = m.textInput.TextStyle.Faint(true)
	m.textInput.CursorStart()
	m.textInput.Blur()
}

// SetValue updates the value of the underlying text input model with the provided string.
func (m *Model) SetValue(val string) {
	m.textInput.SetValue(val)
}

// UpdateProgramContext updates the context of the program by adjusting input width and resetting the text input's value and state.
func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.textInput.Width = m.getInputWidth(ctx)
	m.textInput.SetValue(m.textInput.Value())
	m.textInput.Blur()
}

// getInputWidth calculates and returns the appropriate width for the text input based on the available main content width.
func (m *Model) getInputWidth(ctx *context.ProgramContext) int {
	textWidth := 0
	if m.textInput.Value() == "" {
		textWidth = lipgloss.Width(m.textInput.Placeholder)
	}
	return ctx.MainContentWidth - lipgloss.Width(m.textInput.Prompt) - textWidth - 6
}

// Value returns the current value of the text input associated with the Model.
func (m *Model) Value() string {
	return m.textInput.Value()
}
