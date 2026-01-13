package views

import (
	"fmt"
	"strings"

	"tui-example/types"

	"github.com/charmbracelet/lipgloss"
)

var (
	detailTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true).
			Padding(1, 2)

	detailLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Bold(true).
			Width(15)

	detailValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			PaddingLeft(2)

	detailSectionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("62")).
				Bold(true).
				MarginTop(1).
				Padding(0, 1)

	detailHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(2).
			Padding(1)
)

// RenderResourceDetail renders the resource detail view
func RenderResourceDetail(resource types.Resource, width, height int) string {
	var b strings.Builder

	// Title
	title := detailTitleStyle.Render(
		fmt.Sprintf("📋 %s Details", resource.Type),
	)
	b.WriteString(title)
	b.WriteString("\n\n")

	// Basic Information
	b.WriteString(detailSectionStyle.Render("Basic Information"))
	b.WriteString("\n")

	b.WriteString(detailLabelStyle.Render("Name:"))
	b.WriteString(detailValueStyle.Render(resource.Name))
	b.WriteString("\n")

	b.WriteString(detailLabelStyle.Render("Type:"))
	b.WriteString(detailValueStyle.Render(string(resource.Type)))
	b.WriteString("\n")

	b.WriteString(detailLabelStyle.Render("Namespace:"))
	b.WriteString(detailValueStyle.Render(resource.Namespace))
	b.WriteString("\n")

	b.WriteString(detailLabelStyle.Render("Status:"))
	statusStr := renderStatus(resource.Status)
	b.WriteString(detailValueStyle.Render(statusStr))
	b.WriteString("\n")

	b.WriteString(detailLabelStyle.Render("Age:"))
	b.WriteString(detailValueStyle.Render(resource.Age))
	b.WriteString("\n")

	// Additional Details
	if len(resource.Details) > 0 {
		b.WriteString("\n")
		b.WriteString(detailSectionStyle.Render("Additional Details"))
		b.WriteString("\n")

		for key, value := range resource.Details {
			b.WriteString(detailLabelStyle.Render(key + ":"))
			b.WriteString(detailValueStyle.Render(value))
			b.WriteString("\n")
		}
	}

	// Help text
	helpText := detailHelpStyle.Render(
		"Esc: Back  |  q: Quit",
	)
	b.WriteString("\n")
	b.WriteString(helpText)

	return b.String()
}
