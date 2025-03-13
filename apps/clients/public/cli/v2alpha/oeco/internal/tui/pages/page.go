package pages

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	content "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	sidebar "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/sidebar"
	tasks "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
)

// BaseModel defines a generic model structure that manages UI context, key configuration, and content layout components.
type BaseModel struct {
	Ctx          *context.ProgramContext
	Keys         *keys.KeyMap
	KeyBindings  *config.KeyBindings
	PageSettings *contract.PageSettings

	Default            bool
	CurrentMainContent contract.MainContent
	CurrentSidebar     contract.Sidebar
}

// NewBaseOptions defines options for initializing a base model, including default settings, page configuration,
// main content, sidebar, key mappings, and key bindings.
type NewBaseOptions struct {
	Default            bool
	CurrentMainContent contract.MainContent
	CurrentSidebar     contract.Sidebar
	Keys               *keys.KeyMap
	KeyBindings        *config.KeyBindings
	PageSettings       *contract.PageSettings
}

// NewBaseModel initializes and returns a new BaseModel with the given ProgramContext and configuration options.
func NewBaseModel(ctx *context.ProgramContext, options *NewBaseOptions) *BaseModel {
	if options == nil || options.PageSettings == nil {
		panic("NewBaseModel: PageSettings cannot be nil")
	}

	m := &BaseModel{
		Default:            options.Default,
		CurrentMainContent: content.NewEmptyModel(ctx),
		CurrentSidebar:     sidebar.NewEmptyModel(ctx),
		Ctx:                ctx,
		Keys:               options.Keys,
		KeyBindings:        options.KeyBindings,
		PageSettings:       options.PageSettings,
	}

	if options.CurrentMainContent != nil {
		m.CurrentMainContent = options.CurrentMainContent
	}

	if options.CurrentSidebar != nil {
		m.CurrentSidebar = options.CurrentSidebar
	}

	m.Ctx = m.SyncDimensions(m.Ctx)

	m.Ctx.Logger.Debug("Page: Base Model Initial Configuration")

	return m
}

// InitBase initializes the BaseModel by batching the execution of the base `init` method and returning a command.
func (m *BaseModel) InitBase() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds,
		m.CurrentMainContent.Init(),
		m.CurrentSidebar.Init(),
	)

	return tea.Batch(cmds...)
}

// UpdateBase processes incoming messages and updates the BaseModel state, returning the updated model and commands.
func (m *BaseModel) UpdateBase(msg tea.Msg) (*BaseModel, tea.Cmd) {
	var (
		cmds           []tea.Cmd
		mainContentCmd tea.Cmd
		sidebarCmd     tea.Cmd
	)

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
			m.Ctx.Logger.Debug("Page: Handling keys.Keys.PrevPage", "msg", msg)
			m.CurrentSidebar.Close()
			m.Ctx = m.SyncDimensions(m.Ctx)

		case key.Matches(message, keys.Keys.NextPage):
			m.Ctx.Logger.Debug("Page: Handling keys.Keys.NextPage", "msg", msg)

		case key.Matches(message, keys.Keys.Help):
			m.Ctx.Logger.Debug("Page: Handling keys.Keys.Help", "msg", msg)

		case key.Matches(message, keys.Keys.Quit):
			m.Ctx.Logger.Debug("Page: Handling keys.Keys.Quit", "msg", msg)
		}
	case tasks.TaskFinishedMsg:

		if message.Task.Error != nil {
			m.Ctx.Logger.Error("Page Task finished with error", "id", message.Task.ID, "err", message.Task.Error)
		} else {
			m.Ctx.Logger.Debug("Page Task finished", "id", message.Task.ID)
		}

	case tasks.ClearTaskMsg:
		// m.Footer.SetRightSection("")
		//	delete(m.Tasks, message.TaskID)
		// case spinner.TickMsg:
		//	if len(m.Tasks) > 0 {
		//		taskSpinner, internalTickCmd := m.Spinner.Update(msg)
		//		m.Spinner = taskSpinner
		//		rTask := m.RenderRunningTask()
		//		m.Footer.SetRightSection(rTask)
		//		cmds = append(cmds, internalTickCmd)
		//	}
	}

	m.UpdateProgramContext(m.Ctx)
	_, mainContentCmd = m.CurrentMainContent.Update(msg)
	_, sidebarCmd = m.CurrentSidebar.Update(msg)

	cmds = append(
		cmds,
		mainContentCmd,
		sidebarCmd,
	)

	return m, tea.Batch(cmds...)
}

// ViewBase generates a styled horizontal layout for rendering the provided content within the model's context.
func (m *BaseModel) ViewBase(content string) string {
	return m.Ctx.Styles.Page.ContainerStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			content,
		),
	)
}

// ViewDebug generates a debug view as a *strings.Builder, displaying detailed layout and contextual information.
func (m *BaseModel) ViewDebug() *strings.Builder {
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
func (m *BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}

	m.CurrentMainContent.UpdateProgramContext(ctx)
	m.CurrentSidebar.UpdateProgramContext(ctx)
}

// OnWindowSizeChanged updates the program context and synchronizes dimensions when the window size changes.
func (m *BaseModel) OnWindowSizeChanged(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}

	m.Ctx = ctx
	m.SyncDimensions(m.Ctx)
	m.CurrentMainContent.OnWindowSizeChanged(m.Ctx)
	m.CurrentSidebar.OnWindowSizeChanged(m.Ctx)
}

// SyncDimensions synchronizes the dimensions of the main content and sidebar based on the provided ProgramContext.
func (m *BaseModel) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = m.SyncMainContentDimensions(ctx)
	m.Ctx = m.SyncSidebarDimensions(m.Ctx)

	return m.Ctx
}

// SyncMainContentDimensions adjusts dimensions of the main content area based on the context and sidebar visibility.
func (m *BaseModel) SyncMainContentDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx

	sideBarOffset := theme.SideBarOffset
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
func (m *BaseModel) SyncSidebarDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx

	sideBarOffset := theme.SideBarOffset
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

// GetPageSettings returns the settings for the "Create an Ecosystems" page, including title, default status, and page type.
func (m *BaseModel) GetPageSettings() *contract.PageSettings {
	return m.PageSettings
}
