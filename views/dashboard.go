package views

import (
	"fmt"
	"strings"
	"time"

	"tui-example/types"

	"github.com/charmbracelet/lipgloss"
)

var (
	dashboardTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("62")).
				Bold(true).
				Padding(1, 2)

	dashboardCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(1, 2).
				Margin(1, 0).
				Width(30)

	dashboardCardTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("62")).
				Bold(true).
				MarginBottom(1)

	dashboardMetricStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				MarginTop(1)

	dashboardMetricValueStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("219")).
					Bold(true)

	dashboardHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(2).
			Padding(1)

	progressBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("46"))
)

// RenderDashboard renders the monitoring dashboard view
func RenderDashboard(metrics types.Metrics, lastUpdate time.Time, width, height int) string {
	var b strings.Builder

	// Title
	title := dashboardTitleStyle.Render("📊 Cluster Dashboard")
	b.WriteString(title)
	b.WriteString("\n")

	// Last update time
	updateTime := fmt.Sprintf("Last updated: %s", lastUpdate.Format("15:04:05"))
	b.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 2).
		Render(updateTime))
	b.WriteString("\n\n")

	// Metrics cards in a row
	var cards []string

	// CPU Usage Card
	cpuCard := dashboardCardStyle.Render(
		fmt.Sprintf("%s\n%s\n%s",
			dashboardCardTitleStyle.Render("CPU Usage"),
			renderProgressBar(metrics.CPUUsage, 100),
			dashboardMetricStyle.Render(
				fmt.Sprintf("%s %.1f%%",
					dashboardMetricValueStyle.Render("Current:"),
					metrics.CPUUsage,
				),
			),
		),
	)
	cards = append(cards, cpuCard)

	// Memory Usage Card
	memoryCard := dashboardCardStyle.Render(
		fmt.Sprintf("%s\n%s\n%s",
			dashboardCardTitleStyle.Render("Memory Usage"),
			renderProgressBar(metrics.MemoryUsage, 100),
			dashboardMetricStyle.Render(
				fmt.Sprintf("%s %.1f%%",
					dashboardMetricValueStyle.Render("Current:"),
					metrics.MemoryUsage,
				),
			),
		),
	)
	cards = append(cards, memoryCard)

	// Resource Counts Card
	countsCard := dashboardCardStyle.Render(
		fmt.Sprintf("%s\n%s\n%s\n%s",
			dashboardCardTitleStyle.Render("Resources"),
			dashboardMetricStyle.Render(
				fmt.Sprintf("Pods: %s %d",
					dashboardMetricValueStyle.Render(""),
					metrics.PodCount,
				),
			),
			dashboardMetricStyle.Render(
				fmt.Sprintf("Services: %s %d",
					dashboardMetricValueStyle.Render(""),
					metrics.ServiceCount,
				),
			),
			dashboardMetricStyle.Render(
				fmt.Sprintf("Deployments: %s %d",
					dashboardMetricValueStyle.Render(""),
					metrics.DeploymentCount,
				),
			),
		),
	)
	cards = append(cards, countsCard)

	// Render cards in a row
	cardsRow := lipgloss.JoinHorizontal(lipgloss.Top, cards...)
	b.WriteString(cardsRow)
	b.WriteString("\n")

	// Help text
	helpText := dashboardHelpStyle.Render(
		"r: Refresh  |  Esc: Back  |  q: Quit",
	)
	b.WriteString("\n")
	b.WriteString(helpText)

	return b.String()
}

func renderProgressBar(value, max float64) string {
	barWidth := 20
	filled := int((value / max) * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	if filled < 0 {
		filled = 0
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
	return progressBarStyle.Render(bar)
}
