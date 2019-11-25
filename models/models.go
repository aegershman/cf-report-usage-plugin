package models

// App -
type App struct {
	Actual int
	Desire int
	RAM    int
}

// Apps -
type Apps []App

// Service -
type Service struct {
	Label       string
	ServicePlan string
}

// Services -
type Services []Service

// Space -
type Space struct {
	Name     string
	Apps     Apps
	Services Services
}

// Spaces -
type Spaces []Space

// Org -
type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      Spaces
}

// Orgs -
type Orgs []Org

// SpaceStats is a way to represent the 'business logic'
// of Spaces; we can use it as a way to decorate
// a Space with extra info like billableAIs, etc.
type SpaceStats struct {
	Name                                 string
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	ConsumedMemory                       int
	SpringCloudServicesCount             int
	ServicesSuiteForPivotalPlatformCount int
	BillableAppInstancesCount            int
	BillableServicesCount                int
}

// OrgStats -
type OrgStats struct {
	Name                                 string
	MemoryQuota                          int
	MemoryUsage                          int
	Spaces                               Spaces
	SpaceStats                           []SpaceStats // TODO unsure if this is best model...?
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	SpringCloudServicesCount             int
	ServicesSuiteForPivotalPlatformCount int
	BillableAppInstancesCount            int
	BillableServicesCount                int
}

// AggregateOrgStats describes an aggregated view
// of multiple OrgStats after a Report Execution run
type AggregateOrgStats struct {
	AppInstancesCount         int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	BillableAppInstancesCount int
	SpringCloudServicesCount  int
	BillableServicesCount     int
}

// Report -
// TODO consider breaking into "pre-init" and "post-init" structs,
// e.g. "reportPlan" and "report"? Possibly makes it clearer that you're
// supposed to "execute" the reportPlan to get it to generate the data?
type Report struct {
	Orgs              Orgs
	OrgStats          []OrgStats
	AggregateOrgStats AggregateOrgStats
}
