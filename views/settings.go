package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	settingsTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true).
			Padding(1, 2)

	settingsItemStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0)

	settingsLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Bold(true).
			Width(20)

	settingsValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			PaddingLeft(2)

	settingsHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(2).
			Padding(1)
)

// RenderSettings renders the settings view
func RenderSettings(width, height int) string {
	var b strings.Builder

	// Title
	title := settingsTitleStyle.Render("⚙️  Settings")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Settings items
	settings := []struct {
		label string
		value string
	}{
		{"Auto-refresh:", "Enabled"},
		{"Refresh interval:", "1s"},
		{"Theme:", "Default"},
		{"Log level:", "Info"},
	}

	for _, setting := range settings {
		b.WriteString(settingsItemStyle.Render(
			fmt.Sprintf("%s%s",
				settingsLabelStyle.Render(setting.label),
				settingsValueStyle.Render(setting.value),
			),
		))
		b.WriteString("\n")
	}

	// Help text
	helpText := settingsHelpStyle.Render(
		"Esc: Back  |  q: Quit",
	)
	b.WriteString("\n")
	b.WriteString(helpText)

	return b.String()
}
