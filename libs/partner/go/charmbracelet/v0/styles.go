package charmbraceletloggerv0

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// GetDefaultStyles returns a custom log.Styles instance with predefined styles for log levels, keys, and values.
func GetDefaultStyles() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR!").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	// Add a custom style for key `err`
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)

	return styles
}
