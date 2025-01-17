package footer

import (
	"strings"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx             *context.ProgramContext
	leftSection     *string
	rightSection    *string
	help            help.Model
	ShowAll         bool
	ShowConfirmQuit bool
}

func NewModel(ctx *context.ProgramContext) Model {
	h := help.New()
	h.ShowAll = true
	h.Styles = ctx.Styles.Help.BubbleStyles
	l := ""
	r := ""
	return Model{
		ctx:          ctx,
		help:         h,
		leftSection:  &l,
		rightSection: &r,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			if m.ShowConfirmQuit {
				return m, tea.Quit
			} else {
				m.ShowConfirmQuit = true
			}
		case m.ShowConfirmQuit && !key.Matches(msg, keys.Keys.Quit):
			m.ShowConfirmQuit = false
		case key.Matches(msg, keys.Keys.Help):
			m.ShowAll = !m.ShowAll
		}
	}

	return m, nil
}

func (m Model) View() string {
	var footer string

	if m.ShowConfirmQuit {
		footer = lipgloss.NewStyle().Render("Really quit? (Press q/esc again to quit)")
	} else {
		helpIndicator := lipgloss.NewStyle().
			Background(m.ctx.Theme.FaintText).
			Foreground(m.ctx.Theme.SelectedBackground).
			Padding(0, 1).
			Render("? help")
		viewSwitcher := m.renderViewSwitcher(*m.ctx)
		leftSection := ""
		if m.leftSection != nil {
			leftSection = *m.leftSection
		}
		rightSection := ""
		if m.rightSection != nil {
			rightSection = *m.rightSection
		}
		spacing := lipgloss.NewStyle().
			Background(m.ctx.Theme.SelectedBackground).
			Render(
				strings.Repeat(
					" ",
					utils.Max(0,
						m.ctx.ScreenWidth-lipgloss.Width(
							viewSwitcher,
						)-lipgloss.Width(leftSection)-
							lipgloss.Width(rightSection)-
							lipgloss.Width(
								helpIndicator,
							),
					)))

		footer = m.ctx.Styles.Common.FooterStyle.
			Render(lipgloss.JoinHorizontal(lipgloss.Top, viewSwitcher, leftSection, spacing, rightSection, helpIndicator))
	}

	if m.ShowAll {
		keymap := keys.CreateKeyMapForView(m.ctx.Page)
		fullHelp := m.help.View(keymap)
		return lipgloss.JoinVertical(lipgloss.Top, footer, fullHelp)
	}

	return footer
}

func (m Model) SetWidth(width int) {
	m.help.Width = width
}

func (m Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
	m.help.Styles = ctx.Styles.Help.BubbleStyles
}

func (m Model) renderViewSwitcher(ctx context.ProgramContext) string {
	var view string
	if ctx.Section == config.EnclaveSection {
		view += " Enclave"
	} else if ctx.Section == config.ContextSection {
		view += " Context"
	} else if ctx.Section == config.OrganizationSection {
		view += "⏱ Organization"
	} else if ctx.Section == config.PackageSection {
		view += " Package"
	} else if ctx.Section == config.ConnectorSection {
		view += " Connector"
	} else if ctx.Section == config.ApiSection {
		view += " Api"
	} else if ctx.Section == config.EcosystemSection {
		view += " Ecosystem"
	}

	var user string
	if ctx.User != "" {
		user = ctx.Styles.Tabs.ViewSwitcher.Background(ctx.Theme.FaintText).Render("@" + ctx.User)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, ctx.Styles.Tabs.ViewSwitcher.
		Render(view), user)
}

func (m Model) SetLeftSection(leftSection string) {
	*m.leftSection = leftSection
}

func (m Model) SetRightSection(rightSection string) {
	*m.rightSection = rightSection
}
