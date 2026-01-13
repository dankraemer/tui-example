package types

import "testing"

func TestViewConstants(t *testing.T) {
	views := []View{
		ViewMenu,
		ViewResourceList,
		ViewResourceDetail,
		ViewDashboard,
		ViewSettings,
	}

	// Verify all views are distinct
	for i, v1 := range views {
		for j, v2 := range views {
			if i != j && v1 == v2 {
				t.Errorf("Views at index %d and %d are equal", i, j)
			}
		}
	}
}

func TestResourceTypeConstants(t *testing.T) {
	if ResourceTypePod == "" {
		t.Error("ResourceTypePod should not be empty")
	}

	if ResourceTypeService == "" {
		t.Error("ResourceTypeService should not be empty")
	}

	if ResourceTypeDeployment == "" {
		t.Error("ResourceTypeDeployment should not be empty")
	}

	// Verify they are distinct
	types := []ResourceType{
		ResourceTypePod,
		ResourceTypeService,
		ResourceTypeDeployment,
	}

	for i, t1 := range types {
		for j, t2 := range types {
			if i != j && t1 == t2 {
				t.Errorf("Resource types at index %d and %d are equal", i, j)
			}
		}
	}
}

func TestStatusConstants(t *testing.T) {
	statuses := []Status{
		StatusRunning,
		StatusPending,
		StatusFailed,
		StatusSucceeded,
		StatusUnknown,
	}

	for _, status := range statuses {
		if status == "" {
			t.Errorf("Status %v should not be empty", status)
		}
	}
}

func TestResource(t *testing.T) {
	r := Resource{
		Name:      "test-resource",
		Type:      ResourceTypePod,
		Namespace: "default",
		Status:    StatusRunning,
		Age:       "1d",
		Details:   make(map[string]string),
	}

	if r.Name == "" {
		t.Error("Resource should have a name")
	}

	if r.Type == "" {
		t.Error("Resource should have a type")
	}

	if r.Namespace == "" {
		t.Error("Resource should have a namespace")
	}

	if r.Status == "" {
		t.Error("Resource should have a status")
	}

	if r.Details == nil {
		t.Error("Resource details should be initialized")
	}
}

func TestMetrics(t *testing.T) {
	m := Metrics{
		CPUUsage:        50.0,
		MemoryUsage:     60.0,
		PodCount:        10,
		ServiceCount:    5,
		DeploymentCount: 3,
	}

	if m.CPUUsage < 0 || m.CPUUsage > 100 {
		t.Errorf("CPUUsage should be between 0 and 100, got %f", m.CPUUsage)
	}

	if m.MemoryUsage < 0 || m.MemoryUsage > 100 {
		t.Errorf("MemoryUsage should be between 0 and 100, got %f", m.MemoryUsage)
	}

	if m.PodCount < 0 {
		t.Errorf("PodCount should be non-negative, got %d", m.PodCount)
	}

	if m.ServiceCount < 0 {
		t.Errorf("ServiceCount should be non-negative, got %d", m.ServiceCount)
	}

	if m.DeploymentCount < 0 {
		t.Errorf("DeploymentCount should be non-negative, got %d", m.DeploymentCount)
	}
}
