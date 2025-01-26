package sidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// BaseModel represents a foundational model for managing UI state, viewport properties, and program context.
type BaseModel struct {
	Opened   bool
	Viewport viewport.Model
	Ctx      *context.ProgramContext
}

// NewBaseOptions represents configuration options for initializing a BaseModel.
// It includes a viewport model and an opened state.
type NewBaseOptions struct {
	Viewport viewport.Model
	Opened   bool
}

// NewBaseModel initializes and returns a new BaseModel instance using the provided ProgramContext and options.
func NewBaseModel(ctx *context.ProgramContext, options NewBaseOptions) BaseModel {
	return BaseModel{
		Opened:   options.Opened,
		Viewport: options.Viewport,
		Ctx:      ctx,
	}
}

// UpdateBase updates the BaseModel's context and dimensions, then returns the updated model along with a batched command.
func (m BaseModel) UpdateBase(_ tea.Msg) (BaseModel, tea.Cmd) {
	var cmds []tea.Cmd

	m.UpdateProgramContext(m.Ctx)
	m.SyncDimensions(m.Ctx)
	// TODO: Investigate why we need this
	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	return m, tea.Batch(cmds...)
}

// View renders the sidebar layout with the viewport content and a scroll percentage indicator.
func (m BaseModel) View() string {
	height := m.Ctx.PageContentHeight
	style := m.Ctx.Styles.Sidebar.Root.
		Height(height).
		MaxHeight(height).
		Width(m.Ctx.Config.Defaults.Sidebar.Width).
		MaxWidth(m.Ctx.Config.Defaults.Sidebar.Width)

	return style.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		m.Viewport.View(),
		m.Ctx.Styles.Sidebar.PagerStyle.
			Render(fmt.Sprintf("%d%%", int(m.Viewport.ScrollPercent()*100))),
	))

	// return m.Viewport.View()
	// s := m.ViewDebug()
	// return s.String()
}

// ViewDebug generates and returns a debug view of the current BaseModel's context, including layout dimensions and state.
func (m BaseModel) ViewDebug() *strings.Builder {
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
	s.WriteString("         Open?: " + strconv.FormatBool(m.IsOpen()) + "\n")
	s.WriteString("\n")

	return &s
}

// ScrollToTop moves the viewport of the BaseModel to the top.
func (m BaseModel) ScrollToTop() {
	m.Viewport.GotoTop()
}

// ScrollToBottom scrolls the viewport of the BaseModel to its bottom-most position.
func (m BaseModel) ScrollToBottom() {
	m.Viewport.GotoBottom()
}

// UpdateProgramContext updates the ProgramContext of the BaseModel if the provided context is not nil.
func (m BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}
	// m.Ctx = ctx
}

// OnWindowSizeChanged updates the ProgramContext by syncing its dimensions with the provided context.
func (m BaseModel) OnWindowSizeChanged(ctx *context.ProgramContext) {
	m.Ctx = m.SyncDimensions(ctx)
}

// SyncDimensions updates the BaseModel's viewport dimensions with values from the provided context and returns the updated context.
func (m BaseModel) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx
	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	return m.Ctx
}

// IsOpen returns true if the BaseModel instance is open; otherwise, false.
func (m BaseModel) IsOpen() bool {
	return m.Opened
}

// Open sets the `Opened` property of the `BaseModel` to true, indicating that the model is now open.
func (m BaseModel) Open() {
	// m.Opened = true
}

// Close sets the Opened field of the BaseModel to false, indicating that the model is closed.
func (m BaseModel) Close() {
	// m.Opened = false
}
