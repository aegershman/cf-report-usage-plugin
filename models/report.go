package models

// Reporter -
type Reporter interface {
	AppInstancesCount() int
	AppsCount() int
	BillableAppInstancesCount() int
	BillableServicesCount() int
	MemoryQuota() int
	MemoryUsage() int
	RunningAppInstancesCount() int
	RunningAppsCount() int
	ServicesCount() int
	SpringCloudServicesCount() int
	StoppedAppInstancesCount() int
	StoppedAppsCount() int
}

// Report -
type Report struct {
	Orgs          []Org
	SummaryReport SummaryReporter
}

// NewReport -
func NewReport(orgs []Org) Report {
	return Report{
		Orgs:          orgs,
		SummaryReport: NewSummaryReport(orgs),
	}
}
