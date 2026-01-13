package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"tui-example/types"
)

// model represents the application state
type model struct {
	currentView    types.View
	resources      []types.Resource
	selectedIndex  int
	resourceType   types.ResourceType
	metrics        types.Metrics
	filterText     string
	isLoading      bool
	lastUpdate     time.Time
	width          int
	height         int
}

// initialModel returns the initial model state
func initialModel() model {
	return model{
		currentView:   types.ViewMenu,
		resources:     generateSampleResources(),
		selectedIndex: 0,
		resourceType:  types.ResourceTypePod,
		metrics:       generateSampleMetrics(),
		filterText:    "",
		isLoading:     false,
		lastUpdate:    time.Now(),
	}
}

// generateSampleResources creates sample Kubernetes resources for demonstration
func generateSampleResources() []types.Resource {
	return []types.Resource{
		{
			Name:      "web-server-1",
			Type:      types.ResourceTypePod,
			Namespace: "default",
			Status:    types.StatusRunning,
			Age:       "2d",
			Details: map[string]string{
				"Image":    "nginx:1.21",
				"Node":     "node-1",
				"IP":       "10.244.1.5",
				"Restarts": "0",
			},
		},
		{
			Name:      "web-server-2",
			Type:      types.ResourceTypePod,
			Namespace: "default",
			Status:    types.StatusRunning,
			Age:       "1d",
			Details: map[string]string{
				"Image":    "nginx:1.21",
				"Node":     "node-2",
				"IP":       "10.244.2.3",
				"Restarts": "1",
			},
		},
		{
			Name:      "api-service",
			Type:      types.ResourceTypeService,
			Namespace: "default",
			Status:    types.StatusRunning,
			Age:       "3d",
			Details: map[string]string{
				"Type":      "ClusterIP",
				"ClusterIP":  "10.96.0.1",
				"Port":      "80/TCP",
				"Selector":  "app=api",
			},
		},
		{
			Name:      "frontend-deployment",
			Type:      types.ResourceTypeDeployment,
			Namespace: "production",
			Status:    types.StatusRunning,
			Age:       "5d",
			Details: map[string]string{
				"Replicas":    "3/3",
				"Strategy":    "RollingUpdate",
				"Image":       "frontend:v2.1",
			},
		},
		{
			Name:      "db-pod",
			Type:      types.ResourceTypePod,
			Namespace: "database",
			Status:    types.StatusPending,
			Age:       "1h",
			Details: map[string]string{
				"Image":    "postgres:14",
				"Node":     "node-3",
				"IP":       "Pending",
				"Restarts": "0",
			},
		},
		{
			Name:      "cache-service",
			Type:      types.ResourceTypeService,
			Namespace: "default",
			Status:    types.StatusRunning,
			Age:       "1w",
			Details: map[string]string{
				"Type":      "ClusterIP",
				"ClusterIP":  "10.96.0.2",
				"Port":      "6379/TCP",
				"Selector":  "app=cache",
			},
		},
		{
			Name:      "failed-pod",
			Type:      types.ResourceTypePod,
			Namespace: "default",
			Status:    types.StatusFailed,
			Age:       "30m",
			Details: map[string]string{
				"Image":    "broken:latest",
				"Node":     "node-1",
				"IP":       "N/A",
				"Restarts": "5",
			},
		},
	}
}

// generateSampleMetrics creates sample monitoring metrics
func generateSampleMetrics() types.Metrics {
	return types.Metrics{
		CPUUsage:        45.2,
		MemoryUsage:     62.8,
		PodCount:        12,
		ServiceCount:    8,
		DeploymentCount: 5,
	}
}

// Init is called when the program starts
func (m model) Init() tea.Cmd {
	return nil
}
