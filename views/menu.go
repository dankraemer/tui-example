package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	menuTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true).
			Padding(1, 2)

	menuItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Margin(1, 0)

	menuItemSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("219")).
				Bold(true).
				PaddingLeft(2).
				Margin(1, 0)

	menuHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(2).
			Padding(1)
)

// RenderMenu renders the main menu view
func RenderMenu(selectedIndex int, width, height int) string {
	menuItems := []string{
		"📦 View Pods",
		"🔌 View Services",
		"🚀 View Deployments",
		"📊 Dashboard",
		"⚙️  Settings",
	}

	var b strings.Builder

	// Title
	title := menuTitleStyle.Render("Cloud Native Resource Manager")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Menu items
	for i, item := range menuItems {
		if i == selectedIndex {
			b.WriteString(menuItemSelectedStyle.Render(fmt.Sprintf("> %s", item)))
		} else {
			b.WriteString(menuItemStyle.Render(fmt.Sprintf("  %s", item)))
		}
		b.WriteString("\n")
	}

	// Help text
	helpText := menuHelpStyle.Render(
		"↑/↓: Navigate  |  Enter: Select  |  q: Quit",
	)
	b.WriteString("\n")
	b.WriteString(helpText)

	return b.String()
}
