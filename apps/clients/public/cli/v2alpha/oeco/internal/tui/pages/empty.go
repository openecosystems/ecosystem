package pages

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/content"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	sidebar "apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
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
func NewEmptyModel(ctx *context.ProgramContext) contract.Page {
	m := EmptyModel{}

	m.BaseModel = NewBaseModel[EmptyModelConfig](
		ctx,
		NewBaseOptions[EmptyModelConfig]{
			Default: false,
			PageConfig: &EmptyModelConfig{
				Title: "Empty",
			},
			CurrentMainContent: content.NewEmptyModel(ctx),
			CurrentSidebar:     sidebar.NewEmptyModel(ctx),
			// Keys:               nil,
			// KeyBindings:        nil,
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
		cmd            tea.Cmd
		cmds           []tea.Cmd
		mainContentCmd tea.Cmd
		sidebarCmd     tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil
		_ = message
		//switch {
		//}
	}

	m.UpdateProgramContext(m.Ctx)
	// m.CurrentSidebar, sidebarCmd = m.CurrentSidebar.Update(msg)
	// m.CurrentMainContent, mainContentCmd = m.CurrentMainContent.Update(msg)

	cmds = append(
		cmds,
		cmd,
		mainContentCmd,
		sidebarCmd,
	)
	return m, tea.Batch(cmds...)
}

// View renders the current state of the `EmptyModel` as a string, displaying "Empty Page" vertically aligned to the left.
func (m EmptyModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, "Empty Page")
}

// GetPageSettings returns the page settings for an EmptyModel, including title, default status, key bindings, and type.
func (m EmptyModel) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:         "Empty Page",
		IsDefault:     false,
		KeyBindings:   []key.Binding{},
		ContentHeight: 0,
		Type:          config.EmptyPage,
	}
}

// UpdateProgramContext updates the ProgramContext for the EmptyModel, syncing any relevant data for rendering or behavior.
func (m EmptyModel) UpdateProgramContext(_ *context.ProgramContext) {}
