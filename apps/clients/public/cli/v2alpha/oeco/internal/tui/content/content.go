package content

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// BaseModel represents a foundational structure containing viewport and program context for UI components.
type BaseModel struct {
	Viewport *viewport.Model
	Ctx      *context.ProgramContext
}

// NewBaseOptions is a configuration structure used to initialize a BaseModel with customizable properties like Viewport.
type NewBaseOptions struct {
	Viewport viewport.Model
}

// NewBaseModel initializes and returns a new BaseModel using the given ProgramContext and options.
func NewBaseModel(ctx *context.ProgramContext, options NewBaseOptions) BaseModel {
	return BaseModel{
		Viewport: &options.Viewport,
		Ctx:      ctx,
	}
}

// UpdateBase updates the BaseModel's program context and dimensions, returning the updated BaseModel and a batch of commands.
func (m BaseModel) UpdateBase(_ tea.Msg) (BaseModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmds = append(
		cmds,
		cmd,
	)

	m.UpdateProgramContext(m.Ctx)
	m.SyncDimensions(m.Ctx)
	// TODO: Investigate why we need this
	m.Viewport.Width = m.Ctx.MainContentWidth
	m.Viewport.Height = m.Ctx.PageContentHeight

	return m, tea.Batch(cmds...)
}

// ViewBase returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m BaseModel) ViewBase() string {
	height := m.Ctx.PageContentHeight
	style := m.Ctx.Styles.MainContent.ContainerStyle.
		Height(height).
		MaxHeight(height).
		Width(m.Ctx.MainContentWidth).
		MaxWidth(m.Ctx.MainContentWidth)

	return style.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		m.Viewport.View(),
	))

	// return m.Viewport.View()
	// s := m.ViewDebug()
	// return s.String()
	// return m.Viewport.View()
}

// ViewDebug generates a debug view of the BaseModel's context and dimensions and returns it as a strings.Builder pointer.
func (m BaseModel) ViewDebug() *strings.Builder {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("SideBar Config Width: " + strconv.Itoa(m.Ctx.Config.Defaults.Sidebar.Width) + "\n")
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
	s.WriteString("\n")

	return &s
}

// ScrollToTop resets the viewport's position to the top.
func (m BaseModel) ScrollToTop() {
	m.Viewport.GotoTop()
}

// ScrollToBottom scrolls the viewport content to the bottom.
func (m BaseModel) ScrollToBottom() {
	m.Viewport.GotoBottom()
}

// UpdateProgramContext updates the program context of the BaseModel.
// If the provided context is nil, the function exits without making changes.
// Assigns the passed context to the BaseModel's Ctx field.
func (m BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}
}

// OnWindowSizeChanged updates the BaseModel's context with synchronized dimensions based on the provided ProgramContext.
func (m BaseModel) OnWindowSizeChanged(ctx *context.ProgramContext) {
	m.Ctx = m.SyncDimensions(ctx)
}

// SyncDimensions updates the BaseModel's Viewport dimensions to match the dimensions provided in the ProgramContext.
// If the provided context is nil, it uses the current BaseModel context.
// Returns the updated ProgramContext.
func (m BaseModel) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx
	m.Viewport.Width = m.Ctx.MainContentWidth
	m.Viewport.Height = m.Ctx.PageContentHeight

	return m.Ctx
}
