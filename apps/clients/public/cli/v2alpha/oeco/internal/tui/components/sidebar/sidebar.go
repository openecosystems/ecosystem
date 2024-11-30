package sidebar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

type BaseModel struct {
	Opened   bool
	Viewport viewport.Model
	Ctx      *context.ProgramContext
}

type NewBaseOptions struct {
	Viewport viewport.Model
	Opened   bool
}

func NewBaseModel(ctx *context.ProgramContext, options NewBaseOptions) BaseModel {
	return BaseModel{
		Opened:   options.Opened,
		Viewport: options.Viewport,
		Ctx:      ctx,
	}
}

func (m BaseModel) UpdateBase(_ tea.Msg) (BaseModel, tea.Cmd) {
	var cmds []tea.Cmd

	m.UpdateProgramContext(m.Ctx)
	m.SyncDimensions(m.Ctx)
	// TODO: Investigate why we need this
	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	cmds = append(
		cmds,
	)

	return m, tea.Batch(cmds...)
}

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

func (m BaseModel) ScrollToTop() {
	m.Viewport.GotoTop()
}

func (m BaseModel) ScrollToBottom() {
	m.Viewport.GotoBottom()
}

func (m BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}
	m.Ctx = ctx
}

func (m BaseModel) OnWindowSizeChanged(ctx *context.ProgramContext) {
	m.Ctx = m.SyncDimensions(ctx)
}

func (m BaseModel) SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext {
	if ctx == nil {
		return m.Ctx
	}

	m.Ctx = ctx
	m.Viewport.Width = m.Ctx.SidebarContentWidth
	m.Viewport.Height = m.Ctx.SidebarContentHeight

	return m.Ctx
}

func (m BaseModel) IsOpen() bool {
	return m.Opened
}

func (m BaseModel) Open() {
	m.Opened = true
}

func (m BaseModel) Close() {
	m.Opened = false
}
