package sidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	"github.com/charmbracelet/bubbles/key"
)

// BaseModel represents a foundational model for managing UI state, viewport properties, and program context.
type BaseModel struct {
	Ctx        *context.ProgramContext
	Viewport   *viewport.Model
	Opened     bool
	Data       string
	EmptyState string
}

// NewBaseOptions represents configuration options for initializing a BaseModel.
// It includes a viewport model and an opened state.
type NewBaseOptions struct {
	Viewport   *viewport.Model
	Opened     bool
	EmptyState string
}

// NewBaseModel initializes and returns a new BaseModel instance using the provided ProgramContext and options.
func NewBaseModel(ctx *context.ProgramContext, options *NewBaseOptions) *BaseModel {
	return &BaseModel{
		Opened:   options.Opened,
		Viewport: options.Viewport,
		Ctx:      ctx,
	}
}

// InitBase initializes the BaseModel by batching the execution of the base `init` method and returning a command.
func (m *BaseModel) InitBase() tea.Cmd {
	return tea.Batch()
}

// UpdateBase updates the BaseModel's context and dimensions, then returns the updated model along with a batched command.
func (m *BaseModel) UpdateBase(msg tea.Msg) (*BaseModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch message := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(message, keys.Keys.PageDown):
			m.Viewport.HalfViewDown()

		case key.Matches(message, keys.Keys.PageUp):
			m.Viewport.HalfViewUp()
		}
	}

	m.UpdateProgramContext(m.Ctx)
	m.SyncDimensions(m.Ctx)

	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	return m, tea.Batch(cmds...)
}

// ViewBase renders the sidebar layout with the viewport content and a scroll percentage indicator.
func (m *BaseModel) ViewBase() string {
	if !m.Opened {
		return ""
	}

	height := m.Ctx.PageContentHeight
	style := m.Ctx.Styles.Sidebar.Root.
		Height(height).
		MaxHeight(height).
		Width(m.Ctx.Config.Defaults.Sidebar.Width).
		MaxWidth(m.Ctx.Config.Defaults.Sidebar.Width)

	if m.Data == "" {
		return style.Align(lipgloss.Center).Render(
			lipgloss.PlaceVertical(height, lipgloss.Center, m.EmptyState),
		)
	}

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
	s.WriteString("         Open?: " + strconv.FormatBool(m.IsOpen()) + "\n")
	s.WriteString("\n")

	return &s
}

// ScrollToTop moves the viewport of the BaseModel to the top.
func (m *BaseModel) ScrollToTop() {
	m.Viewport.GotoTop()
}

// ScrollToBottom scrolls the viewport of the BaseModel to its bottom-most position.
func (m *BaseModel) ScrollToBottom() {
	m.Viewport.GotoBottom()
}

// UpdateProgramContext updates the ProgramContext of the BaseModel if the provided context is not nil.
func (m *BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}
	m.Ctx = ctx
	m.Viewport.Height = m.Ctx.MainContentHeight - m.Ctx.Styles.Sidebar.PagerHeight
	m.Viewport.Width = m.GetSidebarContentWidth()
}

// OnWindowSizeChanged updates the ProgramContext by syncing its dimensions with the provided context.
func (m *BaseModel) OnWindowSizeChanged(ctx *context.ProgramContext) {
	m.Ctx = m.SyncDimensions(ctx)
}

// SyncDimensions updates the BaseModel's viewport dimensions with values from the provided context and returns the updated context.
func (m *BaseModel) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx
	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	return m.Ctx
}

// SetContent sets the content of the BaseModel and updates the viewport with the provided string data.
func (m *BaseModel) SetContent(data string) {
	m.Data = data
	m.Viewport.SetContent(data)
}

// GetSidebarContentWidth calculates the width of the sidebar content by considering padding and border dimensions.
func (m *BaseModel) GetSidebarContentWidth() int {
	if m.Ctx.Config == nil {
		return 0
	}
	return m.Ctx.Config.Defaults.Sidebar.Width - 2*m.Ctx.Styles.Sidebar.ContentPadding - m.Ctx.Styles.Sidebar.BorderWidth
}

// SyncSidebar synchronizes the sidebar state and returns a command for subsequent UI updates.
func (m *BaseModel) SyncSidebar() tea.Cmd {
	var cmd tea.Cmd
	return cmd
}

// IsOpen returns true if the BaseModel instance is open; otherwise, false.
func (m *BaseModel) IsOpen() bool {
	return m.Opened
}

// Open sets the `Opened` property of the `BaseModel` to true, indicating that the model is now open.
func (m *BaseModel) Open() {
	m.Opened = true
}

// Close sets the Opened field of the BaseModel to false, indicating that the model is closed.
func (m *BaseModel) Close() {
	m.Opened = false
}
