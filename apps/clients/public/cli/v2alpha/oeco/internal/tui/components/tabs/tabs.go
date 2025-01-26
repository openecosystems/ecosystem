//nolint:revive
package tabs

import (
	"strings"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the state and logic for managing pages, context, and the current page within the application.
type Model struct {
	currentPageId int
	pages         []contract.Page
	ctx           *context.ProgramContext
}

// NewModel initializes and returns a new Model based on the provided context and pages, identifying the current page.
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

// Update applies a message to the Model, updating its state and possibly returning a command to process asynchronously.
func (m Model) Update(_ tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

// View renders the current state of tabs, including their styles and layout, as a string to be displayed in the UI.
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

// GetPages returns the list of pages associated with the model.
func (m Model) GetPages() []contract.Page {
	return m.pages
}

// GetCurrentPageId returns the ID of the currently active page in the model.
func (m Model) GetCurrentPageId() int {
	return m.currentPageId
}

// SetCurrentPageId updates the current page ID of the Model and returns the updated Model.
func (m Model) SetCurrentPageId(id int) Model {
	m.currentPageId = id

	return m
}

// UpdateProgramContext updates the ProgramContext of the Model with the provided context.
func (m Model) UpdateProgramContext(ctx *context.ProgramContext) {
	// m.ctx = ctx
}
