package models

// Report should be present in all reports at all levels,
// e.g. summary level, org level, space level, etc.
// this allows us to calculate common data at discrete layers and be
// confident it'll be available when presenting the data
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
