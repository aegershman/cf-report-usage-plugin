package models

// Report should be present in all reports at all levels,
// e.g. summary level, org level, space level, etc.
// this allows us to calculate common data at discrete layers and be
// confident it'll be available when presenting the data
type Report struct {
	AppInstancesCount                    int    `json:"app_instances_count"`
	AppsCount                            int    `json:"apps_count"`
	BillableAppInstancesCount            int    `json:"billable_app_instances_count"`
	BillableServicesCount                int    `json:"billable_services_count"`
	MemoryQuota                          int    `json:"memory_quota"`
	MemoryUsage                          int    `json:"memory_usage"`
	Name                                 string `json:"name"`
	RunningAppInstancesCount             int    `json:"running_app_instances_count"`
	RunningAppsCount                     int    `json:"running_apps_count"`
	ServicesCount                        int    `json:"services_count"`
	ServicesSuiteForPivotalPlatformCount int    `json:"services_suite_for_pivotal_platform_count"`
	SpringCloudServicesCount             int    `json:"spring_cloud_services_count"`
	StoppedAppInstancesCount             int    `json:"stopped_app_instances_count"`
	StoppedAppsCount                     int    `json:"stopped_apps_count"`
}
