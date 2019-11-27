package models

// OrgReport -
type OrgReport struct {
	orgRef          Org
	Report          Report
	spaceReportsRef []SpaceReport
}

// NewOrgReport -
func NewOrgReport(org Org) *OrgReport {
	var spaceReports []SpaceReport
	for _, space := range org.Spaces {
		spaceReports = append(spaceReports, *NewSpaceReport(space))
	}

	self := &OrgReport{
		orgRef:          org,
		spaceReportsRef: spaceReports,
	}

	self.Report = Report{
		AppInstancesCount:                    self.appInstancesCount(),
		AppsCount:                            self.appsCount(),
		BillableAppInstancesCount:            self.billableAppInstancesCount(),
		BillableServicesCount:                self.billableServicesCount(),
		MemoryQuota:                          self.memoryQuota(),
		MemoryUsage:                          self.memoryUsage(),
		Name:                                 self.name(),
		RunningAppInstancesCount:             self.runningAppInstancesCount(),
		RunningAppsCount:                     self.runningAppsCount(),
		ServicesCount:                        self.servicesCount(),
		ServicesSuiteForPivotalPlatformCount: self.servicesSuiteForPivotalPlatformCount(),
		SpringCloudServicesCount:             self.springCloudServicesCount(),
		StoppedAppInstancesCount:             self.stoppedAppInstancesCount(),
		StoppedAppsCount:                     self.stoppedAppsCount(),
	}

	return self
}

// SpaceReports -
func (o *OrgReport) SpaceReports() []SpaceReport {
	return o.spaceReportsRef
}

func (o *OrgReport) appInstancesCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.AppInstancesCount()
	}
	return count
}

// AppsCount -
func (o *OrgReport) appsCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.appsCount()
	}
	return count
}

// MemoryQuota -
func (o *OrgReport) memoryQuota() int {
	return o.orgRef.MemoryQuota
}

// MemoryUsage -
func (o *OrgReport) memoryUsage() int {
	return o.orgRef.MemoryUsage
}

// Name -
func (o *OrgReport) name() string {
	return o.orgRef.Name
}

// RunningAppInstancesCount -
func (o *OrgReport) runningAppInstancesCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.runningAppInstancesCount()
	}
	return count
}

// RunningAppsCount -
func (o *OrgReport) runningAppsCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.runningAppsCount()
	}
	return count
}

// ServicesCount -
func (o *OrgReport) servicesCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.servicesCount()
	}
	return count
}

// ServicesSuiteForPivotalPlatformCount -
func (o *OrgReport) servicesSuiteForPivotalPlatformCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.servicesSuiteForPivotalPlatformCount()
	}
	return count

}

// StoppedAppInstancesCount -
func (o *OrgReport) stoppedAppInstancesCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.stoppedAppInstancesCount()
	}
	return count
}

// StoppedAppsCount -
func (o *OrgReport) stoppedAppsCount() int {
	count := 0
	for _, space := range o.spaceReportsRef {
		count += space.stoppedAppsCount()
	}
	return count
}

// SpringCloudServicesCount -
func (o *OrgReport) springCloudServicesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.springCloudServicesCount()
	}
	return count
}

// BillableAppInstancesCount -
func (o *OrgReport) billableAppInstancesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.billableAppInstancesCount()
	}
	return count
}

// BillableServicesCount -
func (o *OrgReport) billableServicesCount() int {
	count := 0
	for _, s := range o.spaceReportsRef {
		count += s.billableServicesCount()
	}
	return count
}
