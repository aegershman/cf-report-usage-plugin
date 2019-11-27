package models

// Report -
type Report struct {
	AppInstancesCount                    int
	AppsCount                            int
	BillableAppInstancesCount            int
	BillableServicesCount                int
	MemoryQuota                          int
	MemoryUsage                          int
	Name                                 string
	RunningAppInstancesCount             int
	RunningAppsCount                     int
	ServicesCount                        int
	ServicesSuiteForPivotalPlatformCount int
	SpringCloudServicesCount             int
	StoppedAppInstancesCount             int
	StoppedAppsCount                     int
}

// Reporter -
type Reporter interface {
	AppInstancesCount() int
	AppsCount() int
	BillableAppInstancesCount() int
	BillableServicesCount() int
	MemoryQuota() int
	MemoryUsage() int
	Name() string
	RunningAppInstancesCount() int
	RunningAppsCount() int
	ServicesCount() int
	ServicesSuiteForPivotalPlatformCount() int
	SpringCloudServicesCount() int
	StoppedAppInstancesCount() int
	StoppedAppsCount() int
}
