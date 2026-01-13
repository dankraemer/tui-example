package main

import (
	"testing"
	"tui-example/types"

	tea "github.com/charmbracelet/bubbletea"
)

func TestFilterResources(t *testing.T) {
	resources := []types.Resource{
		{
			Name:      "web-server",
			Type:      types.ResourceTypePod,
			Namespace: "default",
			Status:    types.StatusRunning,
		},
		{
			Name:      "api-service",
			Type:      types.ResourceTypeService,
			Namespace: "default",
			Status:    types.StatusRunning,
		},
		{
			Name:      "db-pod",
			Type:      types.ResourceTypePod,
			Namespace: "database",
			Status:    types.StatusPending,
		},
		{
			Name:      "cache-service",
			Type:      types.ResourceTypeService,
			Namespace: "default",
			Status:    types.StatusRunning,
		},
	}

	tests := []struct {
		name         string
		resources    []types.Resource
		resourceType types.ResourceType
		filterText   string
		expectedLen  int
	}{
		{
			name:         "Filter by Pod type",
			resources:    resources,
			resourceType: types.ResourceTypePod,
			filterText:   "",
			expectedLen:  2,
		},
		{
			name:         "Filter by Service type",
			resources:    resources,
			resourceType: types.ResourceTypeService,
			filterText:   "",
			expectedLen:  2,
		},
		{
			name:         "Filter by type and name",
			resources:    resources,
			resourceType: types.ResourceTypePod,
			filterText:   "web",
			expectedLen:  1,
		},
		{
			name:         "Filter by type and namespace",
			resources:    resources,
			resourceType: types.ResourceTypePod,
			filterText:   "database",
			expectedLen:  1,
		},
		{
			name:         "Filter with no matches",
			resources:    resources,
			resourceType: types.ResourceTypePod,
			filterText:   "nonexistent",
			expectedLen:  0,
		},
		{
			name:         "Case insensitive filter",
			resources:    resources,
			resourceType: types.ResourceTypePod,
			filterText:   "WEB",
			expectedLen:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterResources(tt.resources, tt.resourceType, tt.filterText)
			if len(result) != tt.expectedLen {
				t.Errorf("Expected %d resources, got %d", tt.expectedLen, len(result))
			}

			// Verify all results match the resource type
			for _, r := range result {
				if r.Type != tt.resourceType {
					t.Errorf("Expected resource type %v, got %v", tt.resourceType, r.Type)
				}
			}
		})
	}
}

func TestUpdateWindowSize(t *testing.T) {
	m := initialModel()
	msg := tea.WindowSizeMsg{
		Width:  100,
		Height: 50,
	}

	updated, cmd := m.Update(msg)

	if cmd != nil {
		t.Error("Expected no command for window size update")
	}

	updatedModel, ok := updated.(model)
	if !ok {
		t.Fatal("Expected updated model to be of type model")
	}

	if updatedModel.width != 100 {
		t.Errorf("Expected width to be 100, got %d", updatedModel.width)
	}

	if updatedModel.height != 50 {
		t.Errorf("Expected height to be 50, got %d", updatedModel.height)
	}
}

func TestUpdateQuit(t *testing.T) {
	m := initialModel()
	
	// Test with 'q' key
	msg := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'q'},
	}

	_, cmd := m.Update(msg)

	if cmd == nil {
		t.Error("Expected quit command for 'q' key")
	}

	// Test with Ctrl+C
	msg2 := tea.KeyMsg{
		Type:  tea.KeyCtrlC,
		Runes: []rune{},
	}

	_, cmd2 := m.Update(msg2)

	if cmd2 == nil {
		t.Error("Expected quit command for Ctrl+C")
	}
}

func TestUpdateEscapeFromNonMenu(t *testing.T) {
	m := initialModel()
	m.currentView = types.ViewResourceList

	msg := tea.KeyMsg{
		Type:  tea.KeyEscape,
		Runes: []rune{},
	}

	updated, cmd := m.Update(msg)

	if cmd != nil {
		t.Error("Expected no command for escape key")
	}

	updatedModel, ok := updated.(model)
	if !ok {
		t.Fatal("Expected updated model to be of type model")
	}

	if updatedModel.currentView != types.ViewMenu {
		t.Errorf("Expected view to be ViewMenu after escape, got %v", updatedModel.currentView)
	}

	if updatedModel.selectedIndex != 0 {
		t.Errorf("Expected selectedIndex to be reset to 0, got %d", updatedModel.selectedIndex)
	}

	if updatedModel.filterText != "" {
		t.Errorf("Expected filterText to be cleared, got %s", updatedModel.filterText)
	}
}

func TestUpdateEscapeFromMenu(t *testing.T) {
	m := initialModel()
	m.currentView = types.ViewMenu

	msg := tea.KeyMsg{
		Type:  tea.KeyEscape,
		Runes: []rune{},
	}

	updated, cmd := m.Update(msg)

	if cmd != nil {
		t.Error("Expected no command for escape key from menu")
	}

	updatedModel, ok := updated.(model)
	if !ok {
		t.Fatal("Expected updated model to be of type model")
	}

	// Should stay in menu
	if updatedModel.currentView != types.ViewMenu {
		t.Errorf("Expected view to remain ViewMenu, got %v", updatedModel.currentView)
	}
}
