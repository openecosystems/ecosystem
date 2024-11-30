package tabs

import (
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
)

type Model struct {
	currentPageId int
	pages         []contract.Page
	ctx           *context.ProgramContext
}

func NewModel(ctx *context.ProgramContext, pages []contract.Page) Model {

	pageId := 0
	for i, page := range pages {
		if page.GetPageSettings().Type == ctx.Page {
			pageId = i
			break
		}
	}

	return Model{
		currentPageId: pageId,
		pages:         pages,
		ctx:           ctx,
	}
}

func (m Model) Update(_ tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {

	var tabs []string
	tabs = append(tabs, m.ctx.Styles.Tabs.Logo.Render(""))
	for i, page := range m.pages {
		if m.currentPageId == i {
			tabs = append(tabs, m.ctx.Styles.Tabs.ActiveTab.Render(page.GetPageSettings().Title))
		} else {
			tabs = append(tabs, m.ctx.Styles.Tabs.Tab.Render(page.GetPageSettings().Title))
		}
	}

	renderedTabs := lipgloss.NewStyle().
		Width(m.ctx.ScreenWidth).
		MaxWidth(m.ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, strings.Join(tabs, m.ctx.Styles.Tabs.TabSeparator.Render("|"))))

	return m.ctx.Styles.Tabs.TabsRow.
		Width(m.ctx.ScreenWidth).
		MaxWidth(m.ctx.ScreenWidth).
		Render(renderedTabs)
}

func (m Model) GetPages() []contract.Page {
	return m.pages
}

func (m Model) GetCurrentPageId() int {
	return m.currentPageId
}

func (m Model) SetCurrentPageId(id int) Model {
	m.currentPageId = id

	return m
}

func (m Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
