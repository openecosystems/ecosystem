package pages

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/content"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/sidebar"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BaseModel defines a generic model structure that manages UI context, key configuration, and content layout components.
type BaseModel[Cfg any] struct {
	Ctx         *context.ProgramContext
	Keys        *keys.KeyMap
	KeyBindings *config.KeyBindings

	Default            bool
	PageConfig         *Cfg
	CurrentMainContent contract.MainContent
	CurrentSidebar     contract.Sidebar
}

// NewBaseOptions defines options for initializing a base model, including default settings, page configuration,
// main content, sidebar, key mappings, and key bindings.
type NewBaseOptions[Cfg any] struct {
	Default            bool
	PageConfig         *Cfg
	CurrentMainContent contract.MainContent
	CurrentSidebar     contract.Sidebar
	Keys               *keys.KeyMap
	KeyBindings        *config.KeyBindings
}

// NewBaseModel initializes and returns a new BaseModel with the given ProgramContext and configuration options.
func NewBaseModel[Cfg any](ctx *context.ProgramContext, options NewBaseOptions[Cfg]) BaseModel[Cfg] {
	m := BaseModel[Cfg]{
		Default:            options.Default,
		PageConfig:         options.PageConfig,
		CurrentMainContent: content.NewEmptyModel(ctx),
		CurrentSidebar:     sidebar.NewEmptyModel(ctx),
		Ctx:                ctx,
		Keys:               options.Keys,
		KeyBindings:        options.KeyBindings,
	}

	if options.CurrentMainContent != nil {
		m.CurrentMainContent = options.CurrentMainContent
	}

	if options.CurrentSidebar != nil {
		m.CurrentSidebar = options.CurrentSidebar
	}

	m.Ctx = m.SyncDimensions(m.Ctx)

	return m
}

// UpdateBase processes incoming messages and updates the BaseModel state, returning the updated model and commands.
func (m BaseModel[Cfg]) UpdateBase(msg tea.Msg) (BaseModel[Cfg], tea.Cmd) {
	var cmds []tea.Cmd

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.TogglePreview):
			if m.CurrentSidebar.IsOpen() {
				m.CurrentSidebar.Close()
			} else {
				m.CurrentSidebar.Open()
			}
			m.Ctx = m.SyncDimensions(m.Ctx)
		case key.Matches(message, keys.Keys.PrevPage):
			m.CurrentSidebar.Close()
			m.Ctx = m.SyncDimensions(m.Ctx)

		case key.Matches(message, keys.Keys.NextPage):

		case key.Matches(message, keys.Keys.Help):

		case key.Matches(message, keys.Keys.Quit):
		}
	}

	m.UpdateProgramContext(m.Ctx)

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

// ViewDebug generates a debug view as a *strings.Builder, displaying detailed layout and contextual information.
func (m BaseModel[Cfg]) ViewDebug() *strings.Builder {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("Section: " + string(m.Ctx.Section) + "\n")
	s.WriteString("   Screen Width: " + strconv.Itoa(m.Ctx.ScreenWidth) + "\n")
	s.WriteString("   Screen Height: " + strconv.Itoa(m.Ctx.ScreenHeight) + "\n")
	s.WriteString("\n")
	s.WriteString("   Page Content " + string(m.Ctx.Page) + "\n")
	s.WriteString("      Width: " + strconv.Itoa(m.Ctx.PageContentWidth) + "\n")
	s.WriteString("      Height: " + strconv.Itoa(m.Ctx.PageContentHeight) + "\n")
	s.WriteString("      Main Content\n")
	s.WriteString("         Width: " + strconv.Itoa(m.Ctx.MainContentWidth) + "\n")
	s.WriteString("         Height: " + strconv.Itoa(m.Ctx.MainContentHeight) + "\n")
	s.WriteString("         Body Width: " + strconv.Itoa(m.Ctx.MainContentBodyWidth) + "\n")
	s.WriteString("         Body Height: " + strconv.Itoa(m.Ctx.MainContentBodyHeight) + "\n")
	s.WriteString("      Sidebar Content\n")
	s.WriteString("         Width: " + strconv.Itoa(m.Ctx.SidebarContentWidth) + "\n")
	s.WriteString("         Height: " + strconv.Itoa(m.Ctx.SidebarContentHeight) + "\n")
	s.WriteString("         Body Width: " + strconv.Itoa(m.Ctx.SidebarContentBodyWidth) + "\n")
	s.WriteString("         Body Height: " + strconv.Itoa(m.Ctx.SidebarContentBodyHeight) + "\n")
	s.WriteString("         Open?: " + strconv.FormatBool(m.CurrentSidebar.IsOpen()) + "\n")
	s.WriteString("\n")
	return &s
}

// UpdateProgramContext updates the program context for the BaseModel and its associated components: main content and sidebar.
func (m BaseModel[Cfg]) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}

	// m.Ctx = ctx
	m.CurrentMainContent.UpdateProgramContext(ctx)
	m.CurrentSidebar.UpdateProgramContext(ctx)
}

// OnWindowSizeChanged updates the program context and synchronizes dimensions when the window size changes.
func (m BaseModel[Cfg]) OnWindowSizeChanged(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}

	m.Ctx = ctx
	m.SyncDimensions(m.Ctx)
	m.CurrentMainContent.OnWindowSizeChanged(m.Ctx)
	m.CurrentSidebar.OnWindowSizeChanged(m.Ctx)
}

// SyncDimensions synchronizes the dimensions of the main content and sidebar based on the provided ProgramContext.
func (m BaseModel[Cfg]) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = m.SyncMainContentDimensions(ctx)
	m.Ctx = m.SyncSidebarDimensions(m.Ctx)

	return m.Ctx
}

// SyncMainContentDimensions adjusts dimensions of the main content area based on the context and sidebar visibility.
func (m BaseModel[Cfg]) SyncMainContentDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx

	sideBarOffset := 50
	if m.CurrentSidebar.IsOpen() {
		sideBarOffset = m.Ctx.Config.Defaults.Sidebar.Width
	}
	m.Ctx.MainContentWidth = m.Ctx.PageContentWidth - sideBarOffset
	m.Ctx.MainContentHeight = m.Ctx.PageContentHeight - m.Ctx.Styles.Sidebar.PagerHeight
	m.Ctx.MainContentBodyWidth = m.Ctx.MainContentWidth - m.Ctx.Styles.MainContent.ContainerStyle.GetHorizontalPadding()
	m.Ctx.MainContentBodyHeight = m.Ctx.PageContentHeight - theme.SearchHeight
	m.Ctx = m.CurrentMainContent.SyncDimensions(m.Ctx)

	return m.Ctx
}

// SyncSidebarDimensions recalculates and updates the dimensions of the sidebar and its content in the ProgramContext.
func (m BaseModel[Cfg]) SyncSidebarDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx

	sideBarOffset := 50
	if m.CurrentSidebar.IsOpen() {
		sideBarOffset = m.Ctx.Config.Defaults.Sidebar.Width
	}
	m.Ctx.SidebarContentWidth = sideBarOffset
	m.Ctx.SidebarContentHeight = m.Ctx.PageContentHeight - theme.SearchHeight
	m.Ctx.SidebarContentBodyWidth = m.Ctx.SidebarContentWidth - 2*m.Ctx.Styles.Sidebar.ContentPadding - m.Ctx.Styles.Sidebar.BorderWidth
	m.Ctx.SidebarContentBodyHeight = m.Ctx.PageContentHeight - m.Ctx.Styles.Sidebar.PagerHeight
	m.Ctx = m.CurrentSidebar.SyncDimensions(m.Ctx)

	return m.Ctx
}
