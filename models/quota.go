package models

// Quota -
type Quota struct {
	Name                string `json:"name"`
	TotalServices       string `json:"total_services"`
	TotalRoutes         string `json:"total_routes"`
	TotalPrivateDomains string `json:"total_private_domains"`
	MemoryLimit         int    `json:"memory_limit"`
	InstanceMemoryLimit int    `json:"instance_memory_limit"`
	AppInstanceLimit    int    `json:"app_instance_limit"`
	AppTaskLimit        int    `json:"app_task_limit"`
	TotalServiceKeys    int    `json:"total_service_keys"`
}
