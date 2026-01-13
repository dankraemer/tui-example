package types

// View represents the current view in the application
type View int

const (
	ViewMenu View = iota
	ViewResourceList
	ViewResourceDetail
	ViewDashboard
	ViewSettings
)

// ResourceType represents the type of Kubernetes resource
type ResourceType string

const (
	ResourceTypePod        ResourceType = "Pod"
	ResourceTypeService    ResourceType = "Service"
	ResourceTypeDeployment ResourceType = "Deployment"
)

// Status represents the status of a resource
type Status string

const (
	StatusRunning   Status = "Running"
	StatusPending   Status = "Pending"
	StatusFailed    Status = "Failed"
	StatusSucceeded Status = "Succeeded"
	StatusUnknown   Status = "Unknown"
)

// Resource represents a Kubernetes resource
type Resource struct {
	Name      string
	Type      ResourceType
	Namespace string
	Status    Status
	Age       string
	Details   map[string]string
}

// Metrics represents monitoring metrics
type Metrics struct {
	CPUUsage        float64
	MemoryUsage     float64
	PodCount        int
	ServiceCount    int
	DeploymentCount int
}
