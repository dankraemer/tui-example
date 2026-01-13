package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"tui-example/types"
	"tui-example/views"
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}

// View renders the current view
func (m model) View() string {
	switch m.currentView {
	case types.ViewMenu:
		return views.RenderMenu(m.selectedIndex, m.width, m.height)
	case types.ViewResourceList:
		filtered := filterResources(m.resources, m.resourceType, m.filterText)
		// Adjust selected index to match filtered list
		adjustedIndex := m.selectedIndex
		if adjustedIndex >= len(filtered) {
			adjustedIndex = len(filtered) - 1
		}
		if adjustedIndex < 0 {
			adjustedIndex = 0
		}
		return views.RenderResourceList(
			m.resources,
			adjustedIndex,
			m.resourceType,
			m.filterText,
			m.width,
			m.height,
		)
	case types.ViewResourceDetail:
		filtered := filterResources(m.resources, m.resourceType, m.filterText)
		if len(filtered) > 0 && m.selectedIndex < len(filtered) {
			return views.RenderResourceDetail(
				filtered[m.selectedIndex],
				m.width,
				m.height,
			)
		}
		return "Resource not found"
	case types.ViewDashboard:
		return views.RenderDashboard(m.metrics, m.lastUpdate, m.width, m.height)
	case types.ViewSettings:
		return views.RenderSettings(m.width, m.height)
	default:
		return "Unknown view"
	}
}
