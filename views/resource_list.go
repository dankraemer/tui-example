package views

import (
	"fmt"
	"strings"

	"tui-example/types"

	"github.com/charmbracelet/lipgloss"
)

var (
	resourceListTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("62")).
				Bold(true).
				Padding(1, 2)

	resourceListHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Bold(true).
				Padding(0, 1)

	resourceItemStyle = lipgloss.NewStyle().
			Padding(0, 1)

	resourceItemSelectedStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("219")).
					Bold(true).
					Background(lipgloss.Color("236")).
					Padding(0, 1)

	statusRunningStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("46")).
				Bold(true)

	statusPendingStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("226")).
				Bold(true)

	statusFailedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("196")).
				Bold(true)

	statusSucceededStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("46")).
				Bold(true)

	resourceListHelpStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("241")).
				MarginTop(1).
				Padding(1)

	filterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1)
)

// RenderResourceList renders the resource list view
func RenderResourceList(
	resources []types.Resource,
	selectedIndex int,
	resourceType types.ResourceType,
	filterText string,
	width, height int,
) string {
	var b strings.Builder

	// Filter resources by type and filter text
	filtered := filterResources(resources, resourceType, filterText)

	// Title
	title := resourceListTitleStyle.Render(
		fmt.Sprintf("📦 %s Resources", resourceType),
	)
	b.WriteString(title)
	b.WriteString("\n")

	// Filter display
	if filterText != "" {
		filterDisplay := filterStyle.Render(fmt.Sprintf("Filter: %s", filterText))
		b.WriteString(filterDisplay)
		b.WriteString("\n")
	}

	// Header
	header := resourceListHeaderStyle.Render(
		fmt.Sprintf("%-30s %-15s %-15s %-10s",
			"NAME", "NAMESPACE", "STATUS", "AGE"),
	)
	b.WriteString(header)
	b.WriteString("\n")
	b.WriteString(strings.Repeat("─", width))
	b.WriteString("\n")

	// Resource items
	startIdx := 0
	if selectedIndex >= height-10 {
		startIdx = selectedIndex - (height - 10)
	}
	if startIdx < 0 {
		startIdx = 0
	}

	endIdx := startIdx + (height - 10)
	if endIdx > len(filtered) {
		endIdx = len(filtered)
	}

	if len(filtered) == 0 {
		b.WriteString(resourceItemStyle.Render("No resources found"))
		b.WriteString("\n")
	} else {
		for i := startIdx; i < endIdx; i++ {
			resource := filtered[i]
			isSelected := (i == selectedIndex)

			statusStr := renderStatus(resource.Status)

			line := fmt.Sprintf("%-30s %-15s %-15s %-10s",
				truncate(resource.Name, 28),
				truncate(resource.Namespace, 13),
				statusStr,
				resource.Age,
			)

			if isSelected {
				b.WriteString(resourceItemSelectedStyle.Render("> " + line))
			} else {
				b.WriteString(resourceItemStyle.Render("  " + line))
			}
			b.WriteString("\n")
		}
	}

	// Footer with count
	footer := fmt.Sprintf("\nShowing %d of %d resources", len(filtered), len(resources))
	b.WriteString(resourceListHelpStyle.Render(footer))
	b.WriteString("\n")

	// Help text
	helpText := resourceListHelpStyle.Render(
		"↑/↓: Navigate  |  Enter: View Details  |  /: Filter  |  r: Refresh  |  Esc: Back  |  q: Quit",
	)
	b.WriteString(helpText)

	return b.String()
}

func filterResources(
	resources []types.Resource,
	resourceType types.ResourceType,
	filterText string,
) []types.Resource {
	var filtered []types.Resource

	for _, r := range resources {
		// Filter by type
		if r.Type != resourceType {
			continue
		}

		// Filter by text if provided
		if filterText != "" {
			searchText := strings.ToLower(filterText)
			nameMatch := strings.Contains(strings.ToLower(r.Name), searchText)
			namespaceMatch := strings.Contains(strings.ToLower(r.Namespace), searchText)
			if !nameMatch && !namespaceMatch {
				continue
			}
		}

		filtered = append(filtered, r)
	}

	return filtered
}

func renderStatus(status types.Status) string {
	switch status {
	case types.StatusRunning:
		return statusRunningStyle.Render(string(status))
	case types.StatusPending:
		return statusPendingStyle.Render(string(status))
	case types.StatusFailed:
		return statusFailedStyle.Render(string(status))
	case types.StatusSucceeded:
		return statusSucceededStyle.Render(string(status))
	default:
		return string(status)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
