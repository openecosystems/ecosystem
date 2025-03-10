package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
)

// BaseModel defines a generic model structure that manages UI context, key configuration, and content layout components.
type BaseModel[Cfg any] struct {
	Ctx         *context.ProgramContext
	Keys        *keys.KeyMap
	KeyBindings *config.KeyBindings

	ComponentConfig  *Cfg
	CurrentComponent contract.Component
}

// NewBaseOptions defines options for initializing a base model, including default settings, page configuration,
// main content, sidebar, key mappings, and key bindings.
type NewBaseOptions[Cfg any] struct {
	Default          bool
	ComponentConfig  *Cfg
	CurrentComponent contract.Component
	Keys             *keys.KeyMap
	KeyBindings      *config.KeyBindings
}

// NewBaseModel initializes and returns a new BaseModel with the given ProgramContext and configuration options.
func NewBaseModel[Cfg any](ctx *context.ProgramContext, options NewBaseOptions[Cfg]) BaseModel[Cfg] {
	m := BaseModel[Cfg]{
		ComponentConfig: options.ComponentConfig,
		Ctx:             ctx,
		Keys:            options.Keys,
		KeyBindings:     options.KeyBindings,
	}

	m.Ctx.Logger.Debug("Component: Base Model Initial Configuration")

	return m
}

// UpdateBase processes incoming messages and updates the BaseModel state, returning the updated model and commands.
func (m BaseModel[Cfg]) UpdateBase(msg tea.Msg) (BaseModel[Cfg], tea.Cmd) {
	var cmds []tea.Cmd

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.Help):
			m.Ctx.Logger.Debug("Component: Handling keys.Keys.Help", "msg", msg)

		case key.Matches(message, keys.Keys.Quit):
			m.Ctx.Logger.Debug("Component: Handling keys.Keys.Quit", "msg", msg)
		}
	}

	return m, tea.Batch(cmds...)
}

// ViewBase generates a styled horizontal layout for rendering the provided content within the model's context.
func (m BaseModel[Cfg]) ViewBase(content string) string {
	return m.Ctx.Styles.Page.ContainerStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			content,
		),
	)
}
