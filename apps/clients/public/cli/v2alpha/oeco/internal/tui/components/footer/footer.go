package footer

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

// Model represents the state and behavior of the application UI, managing sections, help views, and user interactions.
type Model struct {
	ctx             *context.ProgramContext
	leftSection     *string
	rightSection    *string
	help            help.Model
	ShowAll         bool
	ShowConfirmQuit bool
}

// NewModel initializes and returns a new Model instance with default help settings and empty left and right sections.
func NewModel(ctx *context.ProgramContext) *Model {
	h := help.New()
	h.ShowAll = true
	h.Styles = ctx.Styles.Help.BubbleStyles
	l := ""
	r := ""
	return &Model{
		ctx:          ctx,
		help:         h,
		leftSection:  &l,
		rightSection: &r,
	}
}

// Update handles incoming messages, updating the model state and determining the command to execute next.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
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

// View generates the string representation of the current model state, including footer and optional help view.
func (m *Model) View() string {
	var footer string

	if m.ShowConfirmQuit {
		footer = lipgloss.NewStyle().Render("Really quit? (Press q/esc again to quit)")
	} else {
		helpIndicator := lipgloss.NewStyle().
			Background(m.ctx.Theme.InvertedText).
			Foreground(m.ctx.Theme.SelectedBackground).
			Padding(0, 1).
			Render("? help")
		viewSwitcher := m.renderViewSwitcher(m.ctx)
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

// SetWidth sets the width of the help model to the specified value.
func (m *Model) SetWidth(width int) {
	m.help.Width = width
}

// UpdateProgramContext updates the model's context and applies styles from the updated context to the help view.
//
//nolint:staticcheck
func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
	m.help.Styles = ctx.Styles.Help.BubbleStyles
}

// renderViewSwitcher generates a horizontal view switcher string based on the current section and user context.
func (m *Model) renderViewSwitcher(ctx *context.ProgramContext) string {
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
	} else if ctx.Section == config.APISection {
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

// SetLeftSection sets the content of the left section in the model footer view. Updates the `leftSection` field.
func (m *Model) SetLeftSection(leftSection string) {
	m.leftSection = &leftSection
}

// SetRightSection sets the value of the right section in the Model.
func (m *Model) SetRightSection(rightSection string) {
	m.rightSection = &rightSection
}
