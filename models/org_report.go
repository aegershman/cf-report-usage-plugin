package models

// OrgReporter -
type OrgReporter interface {
	SpaceReports() []SpaceReport
	Reporter
}

// OrgReport -
type OrgReport struct {
	orgRef          Org
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
		orgRef:          org,
		spaceReportsRef: spaceReports,
		Spaces:          org.Spaces,
	}
}

func (o *OrgReport) SpaceReports() []SpaceReport {
	return o.spaceReportsRef
}
func (o *OrgReport) AppInstancesCount() int                    { return 0 }
func (o *OrgReport) AppsCount() int                            { return 0 }
func (o *OrgReport) MemoryQuota() int                          { return 0 }
func (o *OrgReport) MemoryUsage() int                          { return 0 }
func (o *OrgReport) Name() string                              { return "" }
func (o *OrgReport) RunningAppInstancesCount() int             { return 0 }
func (o *OrgReport) RunningAppsCount() int                     { return 0 }
func (o *OrgReport) ServicesCount() int                        { return 0 }
func (o *OrgReport) ServicesSuiteForPivotalPlatformCount() int { return 0 }
func (o *OrgReport) StoppedAppInstancesCount() int             { return 0 }
func (o *OrgReport) StoppedAppsCount() int                     { return 0 }

// SpringCloudServicesCount -
func (o *OrgReport) SpringCloudServicesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.SpringCloudServicesCount()
	}
	return count
}

// BillableAppInstancesCount -
func (o *OrgReport) BillableAppInstancesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.BillableAppInstancesCount()
	}
	return count
}

// BillableServicesCount -
func (o *OrgReport) BillableServicesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.BillableServicesCount()
	}
	return count
}
