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
	spacesRef       []Space
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
		spacesRef:       org.Spaces,
	}
}

func (o *OrgReport) SpaceReports() []SpaceReport {
	return o.spaceReportsRef
}

// AppInstancesCount -
func (o *OrgReport) AppInstancesCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.AppInstancesCount()
	}
	return count
}

// AppsCount -
func (o *OrgReport) AppsCount() int {
	// count := 0
	// for _, space := range o.spaceReportsRef {
	// 	count += len(space.Apps)
	// }
	// return count
	return 0
}

func (o *OrgReport) MemoryQuota() int { return 0 }
func (o *OrgReport) MemoryUsage() int { return 0 }
func (o *OrgReport) Name() string     { return "" }

// RunningAppInstancesCount -
func (o *OrgReport) RunningAppInstancesCount() int {
	// count := 0
	// for _, space := range org.Spaces {
	// 	count += space.RunningAppInstancesCount()
	// }
	// return count
	return 0
}

// RunningAppsCount -
func (o *OrgReport) RunningAppsCount() int {
	// count := 0
	// for _, space := range org.Spaces {
	// 	count += space.RunningAppsCount()
	// }
	// return count
	return 0
}

// ServicesCount -
func (o *OrgReport) ServicesCount() int {
	// count := 0
	// for _, space := range org.Spaces {
	// 	count += len(space.Services)
	// }
	// return count
	return 0
}

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
