package models

// OrgDecorator -
type OrgDecorator struct {
	org                                  Org
	Name                                 string
	MemoryQuota                          int
	MemoryUsage                          int
	Spaces                               []Space
	SpaceDecorator                       []SpaceDecorator
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	ServicesSuiteForPivotalPlatformCount int
}

// PopulateOrgDecorators -
func PopulateOrgDecorators(orgs []Org, c chan OrgDecorator) {
	for _, org := range orgs {
		OrgDecorator := NewOrgDecorator(org)
		c <- OrgDecorator
	}
	close(c)
}

// NewOrgDecorator -
func NewOrgDecorator(org Org) OrgDecorator {
	return OrgDecorator{
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
func (o *OrgDecorator) SpringCloudServicesCount() int {
	count := 0
	for _, s := range o.SpaceDecorator {
		count += s.SpringCloudServicesCount()
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (o *OrgDecorator) BillableAppInstancesCount() int {
	count := 0
	for _, s := range o.SpaceDecorator {
		count += s.BillableAppInstancesCount()
	}
	return count
}

// BillableServicesCount returns the count of "billable" SIs across all spaces of the org
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (o *OrgDecorator) BillableServicesCount() int {
	count := 0
	for _, s := range o.SpaceDecorator {
		count += s.BillableServicesCount()
	}
	return count
}
