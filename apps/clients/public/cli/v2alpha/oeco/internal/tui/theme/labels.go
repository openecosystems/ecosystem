package theme

import (
	"github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/data"

	"github.com/charmbracelet/lipgloss"
)

// RenderLabels generates a styled, multi-line string of labels formatted to fit within a given sidebar width.
// It arranges the labels horizontally until the width limit is reached, then wraps them to a new line.
// Labels are styled using the provided pillStyle and their background is set based on the color property of each label.
func RenderLabels(sidebarWidth int, labels []data.Label, pillStyle lipgloss.Style) string {
	width := sidebarWidth

	renderedRows := []string{}

	rowContentsWidth := 0
	currentRowLabels := []string{}

	for _, l := range labels {
		currentLabel := pillStyle.
			Background(lipgloss.Color("#" + l.Color)).
			Render(l.Name)

		currentLabelWidth := lipgloss.Width(currentLabel)

		if rowContentsWidth+currentLabelWidth <= width {
			currentRowLabels = append(
				currentRowLabels,
				currentLabel,
			)
			rowContentsWidth += currentLabelWidth
		} else {
			currentRowLabels = append(currentRowLabels, "\n")
			renderedRows = append(renderedRows, lipgloss.JoinHorizontal(lipgloss.Top, currentRowLabels...))

			currentRowLabels = []string{currentLabel}
			rowContentsWidth = currentLabelWidth
		}

		// +1 for the space between labels
		currentRowLabels = append(currentRowLabels, " ")
		rowContentsWidth++
	}

	renderedRows = append(renderedRows, lipgloss.JoinHorizontal(lipgloss.Top, currentRowLabels...))

	return lipgloss.JoinVertical(lipgloss.Left, renderedRows...)
}
