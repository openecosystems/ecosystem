package markdown

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

var markdownStyle *ansi.StyleConfig

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

func GetMarkdownRenderer(width int) glamour.TermRenderer {
	markdownRenderer, _ := glamour.NewTermRenderer(
		glamour.WithStyles(*markdownStyle),
		glamour.WithWordWrap(width),
	)

	return *markdownRenderer
}
