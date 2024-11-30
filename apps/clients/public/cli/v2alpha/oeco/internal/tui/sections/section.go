package sections

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/footer"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/tabs"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/pages"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

type BaseModel struct {
	Ctx           *context.ProgramContext
	Pages         []contract.Page
	CurrentPage   contract.Page
	CurrentPageId int
	Spinner       spinner.Model
	Tabs          tabs.Model
	Footer        footer.Model
	Keys          keys.KeyMap
	SingularForm  string
	PluralForm    string
}

type NewBaseOptions struct {
	Singular      string
	Plural        string
	Pages         []contract.Page
	CurrentPageId int
	Settings      *specv2pb.SpecSettings
}

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

	m.CurrentPage, m.CurrentPageId = m.SetCurrentPage(options.CurrentPageId)
	m.Ctx.Page = m.CurrentPage.GetPageSettings().Type
	m.Ctx.Settings = options.Settings
	m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageId)

	return m
}

// initialize The initial tea.Msg in the ELM architecture
type initialize struct {
	Config config.Config
}

func (m BaseModel) init() tea.Msg {

	cfg, err := config.ParseConfig()
	if err != nil {
		utils.ShowError(err)
		return initialize{Config: cfg}
	}

	return initialize{Config: cfg}
}

// InitBase Returns a [tea.Cmd] that can be used in Updates to change state
func (m BaseModel) InitBase() tea.Cmd {
	return tea.Batch(m.init)
}

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
			m.CurrentPage, m.CurrentPageId = m.SetCurrentPage(m.GetPrevPageId())
			m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageId)
			m.Ctx.Page = m.CurrentPage.GetPageSettings().Type

		case key.Matches(message, keys.Keys.NextPage):
			m.CurrentPage, m.CurrentPageId = m.SetCurrentPage(m.GetNextPageId())
			m.Tabs = m.Tabs.SetCurrentPageId(m.CurrentPageId)
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
			cmd = tea.Quit

		}
	case tea.WindowSizeMsg:
		m.OnWindowSizeChanged(message)
	case initialize:
		m.Ctx.Config = &message.Config
		m.Ctx.Theme = theme.ParseTheme(m.Ctx.Config)
		m.Ctx.Styles = theme.InitStyles(m.Ctx.Theme)
		m.Ctx.Section = m.Ctx.Config.Defaults.Section
		m.Ctx.Page = m.Ctx.Config.Defaults.Page
		m.CurrentPage, m.CurrentPageId = m.SetCurrentPage(m.GetDefaultPageId())
		cmds = append(cmds)
	}

	m.UpdateProgramContext(m.Ctx)

	cmds = append(
		cmds,
	)

	return m, tea.Batch(cmds...)
}

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

func (m BaseModel) ViewDebug() *strings.Builder {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("Section: " + string(m.Ctx.Section) + "\n")
	s.WriteString("   Screen Width: " + strconv.Itoa(m.Ctx.ScreenWidth) + "\n")
	s.WriteString("   Screen Height: " + strconv.Itoa(m.Ctx.ScreenHeight) + "\n")
	s.WriteString("   Default Page Id: " + strconv.Itoa(m.GetDefaultPageId()) + "\n")
	s.WriteString("   Number of Section Pages: " + strconv.Itoa(len(m.Pages)) + "\n")
	for _, p := range m.Pages {
		s.WriteString("      Section Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")
	s.WriteString("   Current Page: " + string(m.Ctx.Page) + "\n")
	s.WriteString("      Id: " + strconv.Itoa(m.CurrentPageId) + "\n")
	s.WriteString("      Title: " + string(m.CurrentPage.GetPageSettings().Title) + "\n")
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
	nextPage := m.GetPageAt(m.GetNextPageId())
	s.WriteString("    Next Page: " + string(nextPage.GetPageSettings().Title) + "\n")
	s.WriteString("       Id: " + strconv.Itoa(m.GetNextPageId()) + "\n")
	prevPage := m.GetPageAt(m.GetPrevPageId())
	s.WriteString("    Previous Page: " + string(prevPage.GetPageSettings().Title) + "\n")
	s.WriteString("       Id: " + strconv.Itoa(m.GetPrevPageId()) + "\n")
	tabPage := m.GetPageAt(m.Tabs.GetCurrentPageId())
	s.WriteString("    Current Tab: " + string(tabPage.GetPageSettings().Title) + "\n")
	s.WriteString("       Id: " + strconv.Itoa(m.Tabs.GetCurrentPageId()) + "\n")
	s.WriteString("    Available Tab Page Count: " + strconv.Itoa(len(m.Tabs.GetPages())) + "\n")
	for _, p := range m.Tabs.GetPages() {
		s.WriteString("       Tab Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")

	return &s
}

func (m BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	m.Ctx = ctx
	m.Tabs.UpdateProgramContext(ctx)
	if m.CurrentPage != nil {
		m.CurrentPage.UpdateProgramContext(ctx)
	}
	m.Footer.UpdateProgramContext(ctx)
}

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

func (m BaseModel) SetCurrentPage(id int) (contract.Page, int) {
	p := m.GetPageAt(id)
	if p == nil {
		p = pages.NewEmptyModel(m.Ctx)
	}
	m.Ctx.Page = p.GetPageSettings().Type
	m.CurrentPageId = id
	m.CurrentPage = p
	m.Tabs = m.Tabs.SetCurrentPageId(id)

	return p, id
}

func (m BaseModel) GetCurrentPage() contract.Page {
	p := m.Pages
	if len(p) == 0 || m.CurrentPageId >= len(p) {
		return nil
	}
	return p[m.CurrentPageId]
}

func (m BaseModel) GetPageAt(id int) contract.Page {
	p := m.Pages
	if len(p) <= id {
		return nil
	}
	return p[id]
}

func (m BaseModel) GetPrevPageId() int {
	return (m.CurrentPageId - 1 + len(m.Pages)) % len(m.Pages)
}

func (m BaseModel) GetNextPageId() int {
	return (m.CurrentPageId + 1) % len(m.Pages)
}

func (m BaseModel) GetDefaultPageId() int {

	for i, page := range m.Pages {
		if page.GetPageSettings().IsDefault {
			return i
		}
	}

	return 0

}
