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
	for _, space := range o.spacesRef {
		count += space.AppInstancesCount()
	}
	return count
}

// AppsCount -
func (org *Org) AppsCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += len(space.Apps)
	}
	return count
}

func (o *OrgReport) MemoryQuota() int { return 0 }
func (o *OrgReport) MemoryUsage() int { return 0 }
func (o *OrgReport) Name() string     { return "" }

// RunningAppInstancesCount -
func (org *Org) RunningAppInstancesCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += space.RunningAppInstancesCount()
	}
	return count
}

// RunningAppsCount -
func (org *Org) RunningAppsCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += space.RunningAppsCount()
	}
	return count
}

// ServicesCount -
func (org *Org) ServicesCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += len(space.Services)
	}
	return count
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
