package main

import (
	"testing"
	"time"
	"tui-example/types"
)

func TestTickCommand(t *testing.T) {
	cmd := tickCommand()
	if cmd == nil {
		t.Fatal("Expected tickCommand to return a command")
	}

	// Execute the command to get the message
	msg := cmd()
	if msg == nil {
		t.Fatal("Expected command to return a message")
	}

	tickMsg, ok := msg.(tickMsg)
	if !ok {
		t.Fatalf("Expected tickMsg, got %T", msg)
	}

	if tickMsg.Time.IsZero() {
		t.Error("Expected tick message to have a non-zero time")
	}
}

func TestLoadResourcesCommand(t *testing.T) {
	cmd := loadResourcesCommand()
	if cmd == nil {
		t.Fatal("Expected loadResourcesCommand to return a command")
	}

	// Execute the command (this will block for 500ms)
	start := time.Now()
	msg := cmd()
	duration := time.Since(start)

	if duration < 100*time.Millisecond {
		t.Error("Expected command to simulate some delay")
	}

	if msg == nil {
		t.Fatal("Expected command to return a message")
	}

	resourceMsg, ok := msg.(resourceLoadedMsg)
	if !ok {
		t.Fatalf("Expected resourceLoadedMsg, got %T", msg)
	}

	if len(resourceMsg.resources) == 0 {
		t.Error("Expected loaded resources to be populated")
	}

	// Verify resources have required fields
	for _, r := range resourceMsg.resources {
		if r.Name == "" {
			t.Error("Expected resource to have a name")
		}
		if r.Type == "" {
			t.Error("Expected resource to have a type")
		}
	}
}

func TestUpdateMetricsCommand(t *testing.T) {
	cmd := updateMetricsCommand()
	if cmd == nil {
		t.Fatal("Expected updateMetricsCommand to return a command")
	}

	// Execute the command
	msg := cmd()
	if msg == nil {
		t.Fatal("Expected command to return a message")
	}

	metricsMsg, ok := msg.(metricsUpdatedMsg)
	if !ok {
		t.Fatalf("Expected metricsUpdatedMsg, got %T", msg)
	}

	metrics := metricsMsg.metrics

	// Verify metrics are within reasonable ranges
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

func TestResourceLoadedMessage(t *testing.T) {
	m := initialModel()
	m.isLoading = true

	resources := []types.Resource{
		{
			Name:      "test-pod",
			Type:      types.ResourceTypePod,
			Namespace: "default",
			Status:    types.StatusRunning,
		},
	}

	msg := resourceLoadedMsg{resources: resources}
	updated, cmd := m.Update(msg)

	if cmd != nil {
		t.Error("Expected no command after resource loaded")
	}

	updatedModel, ok := updated.(model)
	if !ok {
		t.Fatal("Expected updated model to be of type model")
	}

	if updatedModel.isLoading {
		t.Error("Expected isLoading to be false after resources loaded")
	}

	if len(updatedModel.resources) != len(resources) {
		t.Errorf("Expected %d resources, got %d", len(resources), len(updatedModel.resources))
	}
}

func TestMetricsUpdatedMessage(t *testing.T) {
	m := initialModel()
	m.currentView = types.ViewDashboard

	newMetrics := types.Metrics{
		CPUUsage:        50.0,
		MemoryUsage:     60.0,
		PodCount:        10,
		ServiceCount:    5,
		DeploymentCount: 3,
	}

	msg := metricsUpdatedMsg{metrics: newMetrics}
	updated, cmd := m.Update(msg)

	if cmd == nil {
		t.Error("Expected a command (tick) after metrics updated")
	}

	updatedModel, ok := updated.(model)
	if !ok {
		t.Fatal("Expected updated model to be of type model")
	}

	if updatedModel.metrics.CPUUsage != newMetrics.CPUUsage {
		t.Errorf("Expected CPUUsage to be %f, got %f", newMetrics.CPUUsage, updatedModel.metrics.CPUUsage)
	}

	if updatedModel.metrics.MemoryUsage != newMetrics.MemoryUsage {
		t.Errorf("Expected MemoryUsage to be %f, got %f", newMetrics.MemoryUsage, updatedModel.metrics.MemoryUsage)
	}
}
