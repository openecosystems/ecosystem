package theme

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// FormTheme returns a new theme for huh forms
func FormTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Base = t.Focused.Base.BorderForeground(lipgloss.Color("238"))
	t.Focused.Title = t.Focused.Title.Foreground(DefaultTheme.PrimaryColor).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(DefaultTheme.PrimaryColor).Bold(true).MarginBottom(1)
	t.Focused.Directory = t.Focused.Directory.Foreground(DefaultTheme.PrimaryColor)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(DefaultTheme.ErrorColor)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(DefaultTheme.ErrorColor)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(DefaultTheme.SecondaryColor)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(DefaultTheme.SecondaryColor)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(DefaultTheme.SecondaryColor)
	t.Focused.Option = t.Focused.Option.Foreground(DefaultTheme.FocusedForeground)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(DefaultTheme.SecondaryColor)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(DefaultTheme.TertiaryColor)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(DefaultTheme.FocusedForeground)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(DefaultTheme.CreamColor).Background(DefaultTheme.SecondaryColor)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(DefaultTheme.FocusedForeground).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(DefaultTheme.TertiaryColor)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(DefaultTheme.SecondaryColor)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
