package models

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

// Report -
type Report struct {
	orgsRef       []Org
	SummaryReport SummaryReporter
}

// NewReport -
func NewReport(orgs []Org) Report {
	return Report{
		orgsRef:       orgs,
		SummaryReport: NewSummaryReport(orgs),
	}
}
