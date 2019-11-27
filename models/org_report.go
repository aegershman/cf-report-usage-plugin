package models

// OrgReporter -
type OrgReporter interface {
	SpaceReports() []SpaceReport
	Reporter
}

// OrgReport -
type OrgReport struct {
	org             Org
	spaceReportsRef []SpaceReport
	Spaces          []Space
}

// NewOrgReport -
func NewOrgReport(org Org) *OrgReport {
	var spaceReports []SpaceReport
	for _, space := range org.Spaces {
		spaceReports = append(spaceReports, *NewSpaceReport(space))
	}

	return &OrgReport{
		org:             org,
		spaceReportsRef: spaceReports,
		Spaces:          org.Spaces,
	}
}

func (o *OrgReport) SpaceReports() []SpaceReport {
	return o.spaceReportsRef
}

func (o *OrgReport) AppInstancesCount() int { return 0 }
func (o *OrgReport) AppsCount() int         { return 0 }

// func (o *OrgReport) BillableAppInstancesCount() int            { return 0 }
// func (o *OrgReport) BillableServicesCount() int                { return 0 }
func (o *OrgReport) MemoryQuota() int                          { return 0 }
func (o *OrgReport) MemoryUsage() int                          { return 0 }
func (o *OrgReport) Name() string                              { return "" }
func (o *OrgReport) RunningAppInstancesCount() int             { return 0 }
func (o *OrgReport) RunningAppsCount() int                     { return 0 }
func (o *OrgReport) ServicesCount() int                        { return 0 }
func (o *OrgReport) ServicesSuiteForPivotalPlatformCount() int { return 0 }

// func (o *OrgReport) SpringCloudServicesCount() int             { return 0 }
func (o *OrgReport) StoppedAppInstancesCount() int { return 0 }
func (o *OrgReport) StoppedAppsCount() int         { return 0 }

// SpringCloudServicesCount returns total count of SCS services across all spaces of the org
func (o *OrgReport) SpringCloudServicesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
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
	for _, s := range o.spaceReportsRef {
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
	for _, s := range o.spaceReportsRef {
		count += s.BillableServicesCount()
	}
	return count
}
