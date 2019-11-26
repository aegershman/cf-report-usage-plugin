package models

// OrgReport -
type OrgReport struct {
	org                                  Org
	Name                                 string
	MemoryQuota                          int
	MemoryUsage                          int
	Spaces                               []Space
	SpaceReport                          []SpaceReport
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	ServicesSuiteForPivotalPlatformCount int
}

// PopulateOrgReports -
func PopulateOrgReports(orgs []Org, c chan OrgReport) {
	for _, org := range orgs {
		OrgReport := NewOrgReport(org)
		c <- OrgReport
	}
	close(c)
}

// NewOrgReport -
func NewOrgReport(org Org) OrgReport {
	return OrgReport{
		org:                      org,
		Name:                     org.Name,
		MemoryQuota:              org.MemoryQuota,
		MemoryUsage:              org.MemoryUsage,
		Spaces:                   org.Spaces,
		AppsCount:                org.AppsCount(),
		RunningAppsCount:         org.RunningAppsCount(),
		StoppedAppsCount:         org.AppsCount() - org.RunningAppsCount(),
		AppInstancesCount:        org.AppInstancesCount(),
		RunningAppInstancesCount: org.RunningAppInstancesCount(),
		StoppedAppInstancesCount: org.AppInstancesCount() - org.RunningAppInstancesCount(),
		ServicesCount:            org.ServicesCount(),
	}
}

// SpringCloudServicesCount returns total count of SCS services across all spaces of the org
func (o *OrgReport) SpringCloudServicesCount() int {
	count := 0
	for _, s := range o.SpaceReport {
		count += s.SpringCloudServicesCount()
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (o *OrgReport) BillableAppInstancesCount() int {
	count := 0
	for _, s := range o.SpaceReport {
		count += s.BillableAppInstancesCount()
	}
	return count
}

// BillableServicesCount returns the count of "billable" SIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (o *OrgReport) BillableServicesCount() int {
	count := 0
	for _, s := range o.SpaceReport {
		count += s.BillableServicesCount()
	}
	return count
}
