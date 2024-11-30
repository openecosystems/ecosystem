package theme

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Common CommonStyles
	Tabs   struct {
		Logo           lipgloss.Style
		Tab            lipgloss.Style
		ActiveTab      lipgloss.Style
		TabSeparator   lipgloss.Style
		TabsRow        lipgloss.Style
		ViewSwitcher   lipgloss.Style
		ActiveView     lipgloss.Style
		ViewsSeparator lipgloss.Style
		InactiveView   lipgloss.Style
	}
	Section struct {
		ContainerPadding int
		ContainerStyle   lipgloss.Style
		SpinnerStyle     lipgloss.Style
		EmptyStateStyle  lipgloss.Style
		KeyStyle         lipgloss.Style
	}
	Page struct {
		ContainerPadding int
		ContainerStyle   lipgloss.Style
		SpinnerStyle     lipgloss.Style
		EmptyStateStyle  lipgloss.Style
		KeyStyle         lipgloss.Style
	}
	MainContent struct {
		ContainerPadding int
		ContainerStyle   lipgloss.Style
		SpinnerStyle     lipgloss.Style
		EmptyStateStyle  lipgloss.Style
		KeyStyle         lipgloss.Style
		PagerStyle       lipgloss.Style
	}
	Sidebar struct {
		BorderWidth    int
		PagerHeight    int
		ContentPadding int
		Root           lipgloss.Style
		PagerStyle     lipgloss.Style
	}
	Table struct {
		CellStyle                lipgloss.Style
		SelectedCellStyle        lipgloss.Style
		TitleCellStyle           lipgloss.Style
		SingleRuneTitleCellStyle lipgloss.Style
		HeaderStyle              lipgloss.Style
		RowStyle                 lipgloss.Style
	}
	Help struct {
		Text         lipgloss.Style
		KeyText      lipgloss.Style
		BubbleStyles help.Styles
	}
	CommentBox struct {
		Text lipgloss.Style
	}
	Pager struct {
		Height int
		Root   lipgloss.Style
	}
}

func InitStyles(theme Theme) Styles {
	var s Styles

	s.Common = BuildStyles(theme)

	s.Tabs.Logo = lipgloss.NewStyle().
		Faint(false).
		Bold(true).
		//Background(lipgloss.Color("#323DD6")).
		Background(theme.PrimaryColor).
		Padding(0, 1)
	s.Tabs.Tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)
	s.Tabs.ActiveTab = s.Tabs.Tab.
		Faint(false).
		Bold(true).
		Background(theme.SelectedBackground).
		Foreground(theme.PrimaryText)
	s.Tabs.TabSeparator = lipgloss.NewStyle().
		Foreground(theme.SecondaryBorder)
	s.Tabs.TabsRow = lipgloss.NewStyle().
		Height(TabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(theme.PrimaryBorder)
	s.Tabs.ViewSwitcher = lipgloss.NewStyle().
		Background(theme.PrimaryColor).
		Foreground(theme.InvertedText).
		Padding(0, 1).
		Bold(true)
	s.Tabs.ActiveView = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Bold(true).
		Background(theme.SelectedBackground)
	s.Tabs.ViewsSeparator = lipgloss.NewStyle().
		BorderForeground(theme.PrimaryBorder).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true)
	s.Tabs.InactiveView = lipgloss.NewStyle().
		Background(theme.FaintBorder).
		Foreground(theme.SecondaryText)

	s.Section.ContainerPadding = 1
	s.Section.ContainerStyle = lipgloss.NewStyle().
		Padding(0, s.Section.ContainerPadding)
	s.Section.SpinnerStyle = lipgloss.NewStyle().Padding(0, 1)
	s.Section.EmptyStateStyle = lipgloss.NewStyle().
		Faint(true).
		PaddingLeft(1).
		MarginBottom(1)
	s.Section.KeyStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Background(theme.SelectedBackground).
		Padding(0, 1)

	s.Page.ContainerPadding = 1
	s.Page.ContainerStyle = lipgloss.NewStyle().
		Padding(0, s.Section.ContainerPadding)
	s.Page.SpinnerStyle = lipgloss.NewStyle().Padding(0, 1)
	s.Page.EmptyStateStyle = lipgloss.NewStyle().
		Faint(true).
		PaddingLeft(1).
		MarginBottom(1)
	s.Page.KeyStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Background(theme.SelectedBackground).
		Padding(0, 1)

	s.MainContent.ContainerPadding = 1
	s.MainContent.ContainerStyle = lipgloss.NewStyle().
		Padding(0, s.Section.ContainerPadding)
	s.MainContent.SpinnerStyle = lipgloss.NewStyle().Padding(0, 1)
	s.MainContent.EmptyStateStyle = lipgloss.NewStyle().
		Faint(true).
		PaddingLeft(1).
		MarginBottom(1)
	s.MainContent.KeyStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Background(theme.SelectedBackground).
		Padding(0, 1)
	s.MainContent.PagerStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Background(theme.SelectedBackground).
		Foreground(theme.FaintText)

	s.Sidebar.BorderWidth = 1
	s.Sidebar.ContentPadding = 2
	s.Sidebar.Root = lipgloss.NewStyle().
		Padding(0, s.Sidebar.ContentPadding).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{
			Top:         "",
			Bottom:      "",
			Left:        "│",
			Right:       "",
			TopLeft:     "",
			TopRight:    "",
			BottomRight: "",
			BottomLeft:  "",
		}).
		BorderForeground(theme.PrimaryBorder)
	s.Sidebar.PagerStyle = lipgloss.NewStyle().
		Height(s.Sidebar.PagerHeight).
		Bold(true).
		Foreground(theme.FaintText)

	s.Table.CellStyle = lipgloss.NewStyle().PaddingLeft(1).
		PaddingRight(1).
		MaxHeight(1)
	s.Table.SelectedCellStyle = s.Table.CellStyle.
		Background(theme.SelectedBackground)
	s.Table.TitleCellStyle = s.Table.CellStyle.
		Bold(true).
		Foreground(theme.PrimaryText)
	s.Table.SingleRuneTitleCellStyle = s.Table.TitleCellStyle.
		Width(SingleRuneWidth)
	s.Table.HeaderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(theme.FaintBorder).
		BorderBottom(true)
	s.Table.RowStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(theme.FaintBorder)

	s.Help.Text = lipgloss.NewStyle().Foreground(theme.SecondaryText)
	s.Help.KeyText = lipgloss.NewStyle().Foreground(theme.PrimaryText)
	s.Help.BubbleStyles = help.Styles{
		ShortDesc:      s.Help.Text.Foreground(theme.FaintText),
		FullDesc:       s.Help.Text.Foreground(theme.FaintText),
		ShortSeparator: s.Help.Text.Foreground(theme.SecondaryBorder),
		FullSeparator:  s.Help.Text,
		FullKey:        s.Help.KeyText,
		ShortKey:       s.Help.KeyText,
		Ellipsis:       s.Help.Text,
	}

	s.CommentBox.Text = s.Help.Text

	s.Pager.Height = 2
	s.Pager.Root = lipgloss.NewStyle().
		Height(s.Pager.Height).
		MaxHeight(s.Pager.Height).
		PaddingTop(1).
		Bold(true).
		Foreground(theme.FaintText)

	return s
}
