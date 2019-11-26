package models

// OrgDecorator -
type OrgDecorator struct {
	Org                                  Org
	Name                                 string
	MemoryQuota                          int
	MemoryUsage                          int
	Spaces                               []Space
	SpaceStats                           []SpaceStats
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	ServicesSuiteForPivotalPlatformCount int
}

// NewOrgsStats -
func NewOrgsStats(orgs []Org, c chan OrgDecorator) {
	for _, org := range orgs {
		OrgDecorator := NewOrgDecorator(org)
		c <- OrgDecorator
	}
	close(c)
}

// NewOrgDecorator -
func NewOrgDecorator(org Org) OrgDecorator {
	return OrgDecorator{
		Org:                      org,
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
func (os *OrgDecorator) SpringCloudServicesCount() int {
	count := 0
	for _, ss := range os.SpaceStats {
		count += ss.SpringCloudServicesCount()
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (os *OrgDecorator) BillableAppInstancesCount() int {
	count := 0
	for _, ss := range os.SpaceStats {
		count += ss.BillableAppInstancesCount()
	}
	return count
}

// BillableServicesCount returns the count of "billable" SIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (os *OrgDecorator) BillableServicesCount() int {
	count := 0
	for _, ss := range os.SpaceStats {
		count += ss.BillableServicesCount()
	}
	return count
}
