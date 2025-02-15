package sections

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	footer "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/footer"
	tabs "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/tabs"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	pages "apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	theme "apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
	cliv2alphalib "libs/public/go/cli/v2alpha"
)

// BaseModel encapsulates the primary state and components for managing UI layout, navigation, and application context.
type BaseModel struct {
	Ctx           *context.ProgramContext
	Pages         []contract.Page
	CurrentPage   contract.Page
	CurrentPageID int
	Spinner       spinner.Model
	Tabs          tabs.Model
	Footer        *footer.Model
	Keys          keys.KeyMap
	SingularForm  string
	PluralForm    string
}

// NewBaseOptions holds configuration for initializing a base model with singular/plural names, pages, and settings.
type NewBaseOptions struct {
	Singular      string
	Plural        string
	Pages         []contract.Page
	CurrentPageID int
	Settings      *cliv2alphalib.Configuration
}

// NewBaseModel creates and initializes a new BaseModel instance with the provided context and options.
func NewBaseModel(ctx *context.ProgramContext, options NewBaseOptions) BaseModel {
	var p []contract.Page
	if options.Pages != nil {
		p = options.Pages
	}

	m := BaseModel{
		Ctx:          ctx,
		Spinner:      spinner.Model{Spinner: spinner.Dot},
		SingularForm: options.Singular,
		PluralForm:   options.Plural,
		Pages:        p,
		Tabs:         tabs.NewModel(ctx, p),
		Footer:       footer.NewModel(ctx),
	}

	m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(options.CurrentPageID)
	m.Ctx.Page = m.CurrentPage.GetPageSettings().Type
	m.Ctx.Settings = options.Settings
	m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageID)

	return m
}

// initialize represents a configuration initialization containing the application's main configuration.
type initialize struct {
	Config config.Config
}

// init initializes the application by parsing the configuration file and handling potential parsing errors.
// Returns an `initialize` message containing the parsed configuration.
func (m BaseModel) init() tea.Msg {
	cfg, err := config.ParseConfig()
	if err != nil {
		utils.ShowError(err)
		return initialize{Config: cfg}
	}

	return initialize{Config: cfg}
}

// InitBase initializes the BaseModel by batching the execution of the base `init` method and returning a command.
func (m BaseModel) InitBase() tea.Cmd {
	return tea.Batch(m.init)
}

// UpdateBase processes incoming messages to update the state of the BaseModel and returns the updated model and command.
func (m BaseModel) UpdateBase(msg tea.Msg) (BaseModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.PrevPage):
			m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetPrevPageID())
			m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageID)
			m.Ctx.Page = m.CurrentPage.GetPageSettings().Type

		case key.Matches(message, keys.Keys.NextPage):
			m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetNextPageID())
			m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageID)
			m.Ctx.Page = m.CurrentPage.GetPageSettings().Type

		case key.Matches(message, keys.Keys.Help):
			if !m.Footer.ShowAll {
				m.Ctx.PageContentHeight = m.Ctx.PageContentHeight + theme.FooterHeight - theme.ExpandedHelpHeight
			} else {
				m.Ctx.PageContentHeight = m.Ctx.PageContentHeight + theme.ExpandedHelpHeight - theme.FooterHeight
			}

		case key.Matches(message, keys.Keys.Quit):
			if m.Ctx.Config.ConfirmQuit {
				m.Footer, cmd = m.Footer.Update(msg)
				return m, cmd
			}
			// cmd = tea.Quit
		}
	case tea.WindowSizeMsg:
		m.OnWindowSizeChanged(message)
	case initialize:
		m.Ctx.Config = &message.Config
		m.Ctx.Theme = theme.ParseTheme(m.Ctx.Config)
		m.Ctx.Styles = theme.InitStyles(m.Ctx.Theme)
		m.Ctx.Section = m.Ctx.Config.Defaults.Section
		m.Ctx.Page = m.Ctx.Config.Defaults.Page
		m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetDefaultPageID())
	}

	m.UpdateProgramContext(m.Ctx)

	return m, tea.Batch(cmds...)
}

// ViewBase generates and returns a string representation of the current UI, including tabs, content, error, and footer.
func (m BaseModel) ViewBase(content string) string {
	s := strings.Builder{}
	s.WriteString(m.Tabs.View())
	s.WriteString("\n")
	if content == "" {
		content = "No page defined"
	}
	c := lipgloss.JoinHorizontal(
		lipgloss.Top,
		content,
	)
	s.WriteString(c)
	s.WriteString("\n")
	if m.Ctx.Error != nil {
		s.WriteString(
			m.Ctx.Styles.Common.ErrorStyle.
				Width(m.Ctx.ScreenWidth).
				Render(fmt.Sprintf("%s %s",
					m.Ctx.Styles.Common.FailureGlyph,
					lipgloss.NewStyle().
						Foreground(m.Ctx.Theme.ErrorText).
						Render(m.Ctx.Error.Error()),
				)),
		)
	} else {
		s.WriteString(m.Footer.View())
	}

	return s.String()
}

// ViewDebug generates and returns a structured debug representation of the BaseModel's current state as a strings.Builder.
func (m BaseModel) ViewDebug() *strings.Builder {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("Section: " + string(m.Ctx.Section) + "\n")
	s.WriteString("   Screen Width: " + strconv.Itoa(m.Ctx.ScreenWidth) + "\n")
	s.WriteString("   Screen Height: " + strconv.Itoa(m.Ctx.ScreenHeight) + "\n")
	s.WriteString("   Default Page ID: " + strconv.Itoa(m.GetDefaultPageID()) + "\n")
	s.WriteString("   Number of Section Pages: " + strconv.Itoa(len(m.Pages)) + "\n")
	for _, p := range m.Pages {
		s.WriteString("      Section Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")
	s.WriteString("   Current Page: " + string(m.Ctx.Page) + "\n")
	s.WriteString("      ID: " + strconv.Itoa(m.CurrentPageID) + "\n")
	s.WriteString("      Title: " + m.CurrentPage.GetPageSettings().Title + "\n")
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
	nextPage := m.GetPageAt(m.GetNextPageID())
	s.WriteString("    Next Page: " + nextPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.GetNextPageID()) + "\n")
	prevPage := m.GetPageAt(m.GetPrevPageID())
	s.WriteString("    Previous Page: " + prevPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.GetPrevPageID()) + "\n")
	tabPage := m.GetPageAt(m.Tabs.GetCurrentPageId())
	s.WriteString("    Current Tab: " + tabPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.Tabs.GetCurrentPageId()) + "\n")
	s.WriteString("    Available Tab Page Count: " + strconv.Itoa(len(m.Tabs.GetPages())) + "\n")
	for _, p := range m.Tabs.GetPages() {
		s.WriteString("       Tab Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")

	return &s
}

// UpdateProgramContext updates the BaseModel's context and propagates it to tabs, the current page, and the footer.
func (m BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	// m.Ctx = ctx
	m.Tabs.UpdateProgramContext(ctx)
	if m.CurrentPage != nil {
		m.CurrentPage.UpdateProgramContext(ctx)
	}
	m.Footer.UpdateProgramContext(ctx)
}

// OnWindowSizeChanged updates context dimensions and page content size based on the new window size from the message.
func (m BaseModel) OnWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.Ctx.ScreenWidth = msg.Width
	m.Ctx.ScreenHeight = msg.Height
	m.Ctx.PageContentWidth = m.Ctx.ScreenWidth
	if m.Footer.ShowAll {
		m.Ctx.PageContentHeight = msg.Height - theme.TabsHeight - theme.ExpandedHelpHeight
	} else {
		m.Ctx.PageContentHeight = msg.Height - theme.TabsHeight - theme.FooterHeight
	}
	m.CurrentPage.OnWindowSizeChanged(m.Ctx)
}

// SetCurrentPage sets the current page of the BaseModel to the page at the given ID, updates the context and tabs, and returns the page and its ID.
func (m BaseModel) SetCurrentPage(id int) (contract.Page, int) {
	p := m.GetPageAt(id)
	if p == nil {
		p = pages.NewEmptyModel(m.Ctx)
	}
	m.Ctx.Page = p.GetPageSettings().Type
	m.CurrentPageID = id
	m.CurrentPage = p
	m.Tabs = m.Tabs.SetCurrentPageId(id)

	return p, id
}

// GetCurrentPage returns the currently active page from the Pages slice based on the CurrentPageID. Returns nil if no valid page exists.
func (m BaseModel) GetCurrentPage() contract.Page {
	p := m.Pages
	if len(p) == 0 || m.CurrentPageID >= len(p) {
		return nil
	}
	return p[m.CurrentPageID]
}

// GetPageAt retrieves the page at the specified index from the Pages slice. If the index is out of range, it returns nil.
func (m BaseModel) GetPageAt(id int) contract.Page {
	p := m.Pages
	if len(p) <= id {
		return nil
	}
	return p[id]
}

// GetPrevPageID calculates and returns the ID of the previous page in the Pages slice, wrapping around if necessary.
func (m BaseModel) GetPrevPageID() int {
	return (m.CurrentPageID - 1 + len(m.Pages)) % len(m.Pages)
}

// GetNextPageID calculates and returns the ID of the next page, cycling back to the start if the end is reached.
func (m BaseModel) GetNextPageID() int {
	return (m.CurrentPageID + 1) % len(m.Pages)
}

// GetDefaultPageID identifies and returns the index of the default page in the Pages slice. Defaults to 0 if none is found.
func (m BaseModel) GetDefaultPageID() int {
	for i, page := range m.Pages {
		if page.GetPageSettings().IsDefault {
			return i
		}
	}

	return 0
}
