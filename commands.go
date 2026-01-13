package main

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"tui-example/types"
)

// tickMsg is sent periodically to update the dashboard
type tickMsg struct {
	Time time.Time
}

// resourceLoadedMsg is sent when resources are loaded
type resourceLoadedMsg struct {
	resources []types.Resource
}

// metricsUpdatedMsg is sent when metrics are updated
type metricsUpdatedMsg struct {
	metrics types.Metrics
}

// tickCommand returns a command that sends a tick message every second
func tickCommand() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg{Time: t}
	})
}

// loadResourcesCommand simulates an async resource loading operation
func loadResourcesCommand() tea.Cmd {
	return func() tea.Msg {
		// Simulate network delay
		time.Sleep(500 * time.Millisecond)
		return resourceLoadedMsg{
			resources: generateSampleResources(),
		}
	}
}

// updateMetricsCommand simulates updating metrics with some variation
func updateMetricsCommand() tea.Cmd {
	return func() tea.Msg {
		// Simulate slight delay
		time.Sleep(100 * time.Millisecond)
		
		// Generate slightly varied metrics
		rand.Seed(time.Now().UnixNano())
		return metricsUpdatedMsg{
			metrics: types.Metrics{
				CPUUsage:        40.0 + rand.Float64()*20.0,
				MemoryUsage:     55.0 + rand.Float64()*20.0,
				PodCount:        10 + rand.Intn(5),
				ServiceCount:    6 + rand.Intn(4),
				DeploymentCount: 4 + rand.Intn(3),
			},
		}
	}
}
