package models

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

// Stats -
func (orgs Orgs) Stats(c chan OrgStats) {
	for _, org := range orgs {
		orgStats := NewOrgStats(org)
		c <- orgStats
	}
	close(c)
}

// NewOrgStats converts an Org into something decorated with more information
// that can be used for presenting business logic and such
func NewOrgStats(org Org) OrgStats {
	totalUniqueApps := org.AppsCount()
	runningUniqueApps := org.RunningAppsCount()
	stoppedUniqueApps := totalUniqueApps - runningUniqueApps

	billableAppInstancesCount := org.BillableAppInstancesCount()
	appInstancesCount := org.AppInstancesCount()
	runningAppInstancesCount := org.RunningAppInstancesCount()
	stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

	billableServicesCount := org.BillableServicesCount()
	servicesCount := org.ServicesCount()
	springCloudServicesCount := org.SpringCloudServicesCount()

	return OrgStats{
		Name:                      org.Name,
		MemoryQuota:               org.MemoryQuota,
		MemoryUsage:               org.MemoryUsage,
		Spaces:                    org.Spaces,
		AppsCount:                 totalUniqueApps,
		RunningAppsCount:          runningUniqueApps,
		StoppedAppsCount:          stoppedUniqueApps,
		AppInstancesCount:         appInstancesCount,
		RunningAppInstancesCount:  runningAppInstancesCount,
		StoppedAppInstancesCount:  stoppedAppInstancesCount,
		SpringCloudServicesCount:  springCloudServicesCount,
		BillableAppInstancesCount: billableAppInstancesCount,
		BillableServicesCount:     billableServicesCount,
		ServicesCount:             servicesCount,
	}
}
