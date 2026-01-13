package main

import (
	"testing"
	"tui-example/types"
)

func TestInitialModel(t *testing.T) {
	m := initialModel()

	if m.currentView != types.ViewMenu {
		t.Errorf("Expected initial view to be ViewMenu, got %v", m.currentView)
	}

	if m.selectedIndex != 0 {
		t.Errorf("Expected selectedIndex to be 0, got %d", m.selectedIndex)
	}

	if m.resourceType != types.ResourceTypePod {
		t.Errorf("Expected resourceType to be ResourceTypePod, got %v", m.resourceType)
	}

	if len(m.resources) == 0 {
		t.Error("Expected resources to be populated")
	}

	if m.filterText != "" {
		t.Errorf("Expected filterText to be empty, got %s", m.filterText)
	}

	if m.isLoading {
		t.Error("Expected isLoading to be false")
	}
}

func TestGenerateSampleResources(t *testing.T) {
	resources := generateSampleResources()

	if len(resources) == 0 {
		t.Error("Expected sample resources to be generated")
	}

	// Check that we have different resource types
	hasPod := false
	hasService := false
	hasDeployment := false

	for _, r := range resources {
		switch r.Type {
		case types.ResourceTypePod:
			hasPod = true
		case types.ResourceTypeService:
			hasService = true
		case types.ResourceTypeDeployment:
			hasDeployment = true
		}

		if r.Name == "" {
			t.Error("Expected resource to have a name")
		}

		if r.Namespace == "" {
			t.Error("Expected resource to have a namespace")
		}
	}

	if !hasPod {
		t.Error("Expected at least one Pod resource")
	}

	if !hasService {
		t.Error("Expected at least one Service resource")
	}

	if !hasDeployment {
		t.Error("Expected at least one Deployment resource")
	}
}

func TestGenerateSampleMetrics(t *testing.T) {
	metrics := generateSampleMetrics()

	if metrics.CPUUsage < 0 || metrics.CPUUsage > 100 {
		t.Errorf("Expected CPUUsage to be between 0 and 100, got %f", metrics.CPUUsage)
	}

	if metrics.MemoryUsage < 0 || metrics.MemoryUsage > 100 {
		t.Errorf("Expected MemoryUsage to be between 0 and 100, got %f", metrics.MemoryUsage)
	}

	if metrics.PodCount < 0 {
		t.Errorf("Expected PodCount to be non-negative, got %d", metrics.PodCount)
	}

	if metrics.ServiceCount < 0 {
		t.Errorf("Expected ServiceCount to be non-negative, got %d", metrics.ServiceCount)
	}

	if metrics.DeploymentCount < 0 {
		t.Errorf("Expected DeploymentCount to be non-negative, got %d", metrics.DeploymentCount)
	}
}
