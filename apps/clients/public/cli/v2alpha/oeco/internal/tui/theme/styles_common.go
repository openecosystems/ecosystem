package theme

import (
	"github.com/charmbracelet/lipgloss"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
)

var (
	SearchHeight       = 3
	FooterHeight       = 1
	ExpandedHelpHeight = 14
	InputBoxHeight     = 8
	SingleRuneWidth    = 4
	MainContentPadding = 1
	TabsBorderHeight   = 1
	TabsContentHeight  = 2
	TabsHeight         = TabsBorderHeight + TabsContentHeight
	ViewSwitcherMargin = 1
	TableHeaderHeight  = 2
)

type CommonStyles struct {
	MainTextStyle lipgloss.Style
	FooterStyle   lipgloss.Style
	ErrorStyle    lipgloss.Style
	WaitingGlyph  string
	FailureGlyph  string
	SuccessGlyph  string
}

func BuildStyles(theme Theme) CommonStyles {
	var s CommonStyles

	s.MainTextStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Bold(true)
	s.FooterStyle = lipgloss.NewStyle().
		Background(theme.SelectedBackground).
		Height(FooterHeight)

	s.ErrorStyle = s.FooterStyle.
		Foreground(theme.ErrorText).
		MaxHeight(FooterHeight)

	s.WaitingGlyph = lipgloss.NewStyle().
		Foreground(theme.FaintText).
		Render(constants.WaitingIcon)
	s.FailureGlyph = lipgloss.NewStyle().
		Foreground(theme.ErrorText).
		Render(constants.FailureIcon)
	s.SuccessGlyph = lipgloss.NewStyle().
		Foreground(theme.SuccessText).
		Render(constants.SuccessIcon)

	return s
}
