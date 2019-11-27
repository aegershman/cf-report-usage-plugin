package models

// OrgReporter -
type OrgReporter interface {
	SpaceReports() []SpaceReport
	Reporter
}

// OrgReport -
type OrgReport struct {
	org          Org
	SpaceReports []SpaceReport
	Spaces       []Space
}

// NewOrgReport -
func NewOrgReport(org Org) OrgReport {
	var spaceReports []SpaceReport
	for _, space := range org.Spaces {
		spaceReports = append(spaceReports, NewSpaceReport(space))
	}

	return OrgReport{
		org:          org,
		SpaceReports: spaceReports,
		Spaces:       org.Spaces,
	}
}

// SpringCloudServicesCount returns total count of SCS services across all spaces of the org
func (o *OrgReport) SpringCloudServicesCount() int {
	count := 0
	for _, s := range o.SpaceReports {
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
	for _, s := range o.SpaceReports {
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
	for _, s := range o.SpaceReports {
		count += s.BillableServicesCount()
	}
	return count
}
