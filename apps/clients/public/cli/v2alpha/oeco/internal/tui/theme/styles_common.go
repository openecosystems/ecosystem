package theme

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"

	"github.com/charmbracelet/lipgloss"
)

// SearchHeight defines the height of the search bar.
// FooterHeight defines the height of the footer section.
// ExpandedHelpHeight defines the height of the expanded help section.
// InputBoxHeight defines the height of the input box.
// SingleRuneWidth defines the width allocated for a single rune.
// MainContentPadding defines the padding used for the main content area.
// TabsBorderHeight defines the height of the tab borders.
// TabsContentHeight defines the height of the tab content area.
// TabsHeight combines border and content heights to define total tab height.
// ViewSwitcherMargin defines the margin around the view switcher.
// TableHeaderHeight defines the height of the table header section.
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

// CommonStyles defines a set of reusable styles and symbols used for consistent UI rendering throughout the application.
type CommonStyles struct {
	MainTextStyle lipgloss.Style
	FooterStyle   lipgloss.Style
	ErrorStyle    lipgloss.Style
	WaitingGlyph  string
	FailureGlyph  string
	SuccessGlyph  string
}

// BuildStyles generates and returns a CommonStyles struct by applying styles based on the provided Theme configuration.
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
