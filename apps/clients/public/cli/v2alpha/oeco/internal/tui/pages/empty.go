package pages

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

type EmptyModelConfig struct {
	Title string
}

type EmptyModel struct {
	BaseModel[EmptyModelConfig]
}

func NewEmptyModel(ctx *context.ProgramContext) contract.Page {

	m := &EmptyModel{}

	m.BaseModel = NewBaseModel[EmptyModelConfig](
		ctx,
		NewBaseOptions[EmptyModelConfig]{
			Default: false,
			PageConfig: &EmptyModelConfig{
				Title: "Empty",
			},
			CurrentMainContent: content.NewEmptyModel(ctx),
			CurrentSidebar:     sidebar.NewEmptyModel(ctx),
			//Keys:               nil,
			//KeyBindings:        nil,
		},
	)

	return m

}

func (m EmptyModel) Update(msg tea.Msg) (EmptyModel, tea.Cmd) {

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
		switch {

		}
	}

	m.UpdateProgramContext(m.Ctx)
	//m.CurrentSidebar, sidebarCmd = m.CurrentSidebar.Update(msg)
	//m.CurrentMainContent, mainContentCmd = m.CurrentMainContent.Update(msg)

	cmds = append(
		cmds,
		cmd,
		mainContentCmd,
		sidebarCmd,
	)
	return m, tea.Batch(cmds...)
}

func (m EmptyModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, "Empty Page")
}

func (m EmptyModel) GetPageSettings() contract.PageSettings {
	return contract.PageSettings{
		Title:         "Empty Page",
		IsDefault:     false,
		KeyBindings:   []key.Binding{},
		ContentHeight: 0,
		Type:          config.EmptyPage,
	}
}

func (m EmptyModel) UpdateProgramContext(_ *context.ProgramContext) {}
