package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"tui-example/types"
)

// Update handles messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			if m.currentView != types.ViewMenu {
				m.currentView = types.ViewMenu
				m.selectedIndex = 0
				m.filterText = ""
			}
			return m, nil

		default:
			// Handle view-specific key bindings
			return handleViewKeys(m, msg)
		}

	case tickMsg:
		if m.currentView == types.ViewDashboard {
			m.lastUpdate = msg.Time
			// Update metrics periodically
			return m, updateMetricsCommand()
		}
		return m, tickCommand()

	case resourceLoadedMsg:
		m.resources = msg.resources
		m.isLoading = false
		return m, nil

	case metricsUpdatedMsg:
		m.metrics = msg.metrics
		m.lastUpdate = time.Now()
		return m, tickCommand()
	}

	return m, tea.Batch(cmds...)
}

// handleViewKeys handles key bindings specific to each view
func handleViewKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case types.ViewMenu:
		return handleMenuKeys(m, msg)
	case types.ViewResourceList:
		return handleResourceListKeys(m, msg)
	case types.ViewResourceDetail:
		return handleResourceDetailKeys(m, msg)
	case types.ViewDashboard:
		return handleDashboardKeys(m, msg)
	case types.ViewSettings:
		return handleSettingsKeys(m, msg)
	default:
		return m, nil
	}
}

// handleMenuKeys handles key bindings for the menu view
func handleMenuKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.selectedIndex > 0 {
			m.selectedIndex--
		}
		return m, nil

	case "down", "j":
		if m.selectedIndex < 4 {
			m.selectedIndex++
		}
		return m, nil

	case "enter", " ":
		switch m.selectedIndex {
		case 0: // View Pods
			m.currentView = types.ViewResourceList
			m.resourceType = types.ResourceTypePod
			m.selectedIndex = 0
			m.filterText = ""
			return m, nil

		case 1: // View Services
			m.currentView = types.ViewResourceList
			m.resourceType = types.ResourceTypeService
			m.selectedIndex = 0
			m.filterText = ""
			return m, nil

		case 2: // View Deployments
			m.currentView = types.ViewResourceList
			m.resourceType = types.ResourceTypeDeployment
			m.selectedIndex = 0
			m.filterText = ""
			return m, nil

		case 3: // Dashboard
			m.currentView = types.ViewDashboard
			return m, tickCommand()

		case 4: // Settings
			m.currentView = types.ViewSettings
			return m, nil
		}
	}

	return m, nil
}

// handleResourceListKeys handles key bindings for the resource list view
func handleResourceListKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	filtered := filterResources(m.resources, m.resourceType, m.filterText)

	switch msg.String() {
	case "up", "k":
		if m.selectedIndex > 0 {
			m.selectedIndex--
		}
		return m, nil

	case "down", "j":
		if m.selectedIndex < len(filtered)-1 {
			m.selectedIndex++
		}
		return m, nil

	case "enter":
		if len(filtered) > 0 && m.selectedIndex < len(filtered) {
			m.currentView = types.ViewResourceDetail
		}
		return m, nil

	case "/":
		// Start filtering (simplified - in a real app, you'd use a text input component)
		m.filterText = ""
		return m, nil

	case "r":
		m.isLoading = true
		return m, loadResourcesCommand()
	}

	return m, nil
}

// handleResourceDetailKeys handles key bindings for the resource detail view
func handleResourceDetailKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Most keys just go back
	return m, nil
}

// handleDashboardKeys handles key bindings for the dashboard view
func handleDashboardKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "r":
		return m, updateMetricsCommand()
	}
	return m, nil
}

// handleSettingsKeys handles key bindings for the settings view
func handleSettingsKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

// filterResources filters resources by type and filter text
func filterResources(resources []types.Resource, resourceType types.ResourceType, filterText string) []types.Resource {
	var filtered []types.Resource

	for _, r := range resources {
		if r.Type != resourceType {
			continue
		}

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
