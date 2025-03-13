package theme

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// FormTheme defines and returns a customized theme for form elements, including focused and blurred styles.
func FormTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Base = t.Focused.Base.BorderForeground(DefaultTheme.SelectedBackground)
	t.Focused.Title = t.Focused.Title.Foreground(DefaultTheme.PrimaryColor500).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(DefaultTheme.PrimaryColor500).Bold(true).MarginBottom(1)
	t.Focused.Description = t.Focused.Description.Foreground(DefaultTheme.FaintText)
	t.Focused.Directory = t.Focused.Directory.Foreground(DefaultTheme.PrimaryColor500)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(DefaultTheme.ErrorColor)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(DefaultTheme.ErrorColor)
	t.Focused.Directory = t.Focused.Directory.Foreground(DefaultTheme.PrimaryColor500)
	t.Focused.File = t.Focused.File.Foreground(DefaultTheme.FocusedForeground)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(DefaultTheme.ErrorColor)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(DefaultTheme.SecondaryColor500)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(DefaultTheme.SecondaryColor500)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(DefaultTheme.SecondaryColor500)
	t.Focused.Option = t.Focused.Option.Foreground(DefaultTheme.FocusedForeground)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(DefaultTheme.SecondaryColor500)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(DefaultTheme.TertiaryColor500)

	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.Foreground(DefaultTheme.TertiaryColor500).SetString("✓ ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(DefaultTheme.FocusedForeground)
	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.Foreground(DefaultTheme.PrimaryColor400).SetString("• ")
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(DefaultTheme.CreamColor).Background(DefaultTheme.SecondaryColor500).Bold(true)
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(DefaultTheme.FocusedForeground).Background(DefaultTheme.SelectedBackground)
	t.Focused.Next = t.Focused.FocusedButton

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(DefaultTheme.TertiaryColor500)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(DefaultTheme.PrimaryColor400)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(DefaultTheme.SecondaryColor500)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})

	return t
}
