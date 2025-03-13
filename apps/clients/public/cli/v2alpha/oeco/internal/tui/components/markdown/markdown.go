package markdown

import (
	"github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

// markdownStyle defines the ANSI style configuration for rendering Markdown content based on the terminal’s color scheme.
var markdownStyle *ansi.StyleConfig

// InitializeMarkdownStyle initializes the markdownStyle variable with a dark or light theme based on the hasDarkBackground flag.
// It does nothing if the markdownStyle has already been initialized.
func InitializeMarkdownStyle(hasDarkBackground bool) {
	if markdownStyle != nil {
		return
	}
	if hasDarkBackground {
		markdownStyle = &theme.DarkStyleConfig
	} else {
		markdownStyle = &theme.LightStyleConfig
	}
}

// GetMarkdownRenderer creates and returns a Markdown renderer with a specified word-wrap width.
func GetMarkdownRenderer(width int) glamour.TermRenderer {
	markdownRenderer, _ := glamour.NewTermRenderer(
		glamour.WithStyles(*markdownStyle),
		glamour.WithWordWrap(width),
	)

	return *markdownRenderer
}
