# TUI Example - Cloud Native Resource Manager

A comprehensive Terminal User Interface (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that demonstrates features commonly needed in cloud native tooling applications.

## Overview

This application showcases a multi-view TUI that simulates a Kubernetes resource manager, demonstrating key Bubble Tea concepts and patterns useful for building low-level tooling interfaces.

## Features Demonstrated

- **Multi-view Navigation**: Seamless navigation between different views (menu, resource lists, details, dashboard, settings)
- **Real-time Updates**: Live dashboard with auto-refreshing metrics using ticker commands
- **List Management**: Scrollable resource lists with keyboard navigation and selection
- **Status Indicators**: Color-coded status displays for resource states
- **Keyboard Controls**: Comprehensive key bindings including vim-style navigation (j/k)
- **Async Operations**: Simulated async resource loading and metric updates
- **Filtering**: Resource filtering by type and search text
- **Styled UI**: Beautiful terminal UI using Lip Gloss for styling

## Prerequisites

- Go 1.21 or later
- A terminal that supports ANSI colors

## Installation

1. Clone or navigate to this directory
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Building and Running

Build the application:
```bash
go build -o tui-example
```

Run the application:
```bash
./tui-example
```

Or run directly:
```bash
go run .
```

## Usage

### Navigation

- **Arrow Keys / j/k**: Navigate up and down in lists and menus
- **Enter / Space**: Select an item or confirm action
- **Esc**: Go back to previous view
- **q / Ctrl+C**: Quit the application

### Views

#### Main Menu
The entry point with options to:
- View Pods
- View Services
- View Deployments
- View Dashboard
- Settings

#### Resource Lists
Displays a scrollable list of resources with:
- Name, Namespace, Status, and Age columns
- Color-coded status indicators (Running, Pending, Failed, etc.)
- Selection highlighting
- Filter support (press `/` to filter)

**Controls:**
- `↑/↓` or `j/k`: Navigate
- `Enter`: View resource details
- `r`: Refresh resources
- `Esc`: Return to menu

#### Resource Detail
Shows detailed information about a selected resource including:
- Basic information (name, type, namespace, status, age)
- Additional details (image, node, IP, etc.)

**Controls:**
- `Esc`: Return to resource list

#### Dashboard
Real-time monitoring dashboard displaying:
- CPU usage with progress bar
- Memory usage with progress bar
- Resource counts (pods, services, deployments)
- Auto-updating metrics

**Controls:**
- `r`: Manually refresh metrics
- `Esc`: Return to menu

#### Settings
Configuration view (demonstration purposes)

**Controls:**
- `Esc`: Return to menu

## Architecture

The application follows Bubble Tea's Model-View-Update pattern:

- **Model** (`model.go`): Application state management
- **View** (`views/`): UI rendering logic for each view
- **Update** (`update.go`): Message handling and state transitions
- **Commands** (`commands.go`): Async operations and side effects
- **Types** (`types/`): Shared data structures

### Key Concepts

1. **Messages**: Events that trigger state updates (key presses, ticks, async results)
2. **Commands**: Functions that return messages, enabling async operations
3. **Views**: Functions that render the current state to a string
4. **Update Loop**: Handles messages and returns updated models and commands

## Use Cases for Cloud Native Tooling

This demo showcases patterns useful for:

- **Kubernetes Resource Viewers**: Browse and inspect cluster resources
- **Container Log Viewers**: Stream and filter container logs
- **Cluster Monitoring Dashboards**: Real-time metrics and health monitoring
- **Deployment Status Trackers**: Track deployment progress and status
- **Resource Management Tools**: Manage and configure resources
- **Debugging Tools**: Inspect and troubleshoot cluster issues

## Project Structure

```
tui-example/
├── main.go                 # Entry point and View() method
├── model.go               # Model struct and initial state
├── update.go              # Update function with message handling
├── commands.go            # Command definitions for async operations
├── types/
│   └── types.go          # Shared data types
├── views/
│   ├── menu.go           # Main menu view
│   ├── resource_list.go  # Resource list view
│   ├── resource_detail.go # Resource detail view
│   ├── dashboard.go      # Monitoring dashboard
│   └── settings.go       # Settings view
├── go.mod                # Go module file
└── README.md            # This file
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout

## Learning Resources

- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Bubble Tea Tutorial](https://github.com/charmbracelet/bubbletea/tree/main/tutorials)
- [Lip Gloss Documentation](https://github.com/charmbracelet/lipgloss)

## License

Apache License 2.0
