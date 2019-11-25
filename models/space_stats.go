package models

// SpaceStats is a way to represent the 'business logic'
// of Spaces; we can use it as a way to decorate
// a Space with extra info like billableAIs, etc.
type SpaceStats struct {
	Name                                 string
	AppsCount                            int
	RunningAppsCount                     int
	StoppedAppsCount                     int
	AppInstancesCount                    int
	RunningAppInstancesCount             int
	StoppedAppInstancesCount             int
	ServicesCount                        int
	ConsumedMemory                       int
	SpringCloudServicesCount             int
	ServicesSuiteForPivotalPlatformCount int
	BillableAppInstancesCount            int
	BillableServicesCount                int
}

// Stats -
func (spaces Spaces) Stats(c chan SpaceStats, skipSIcount bool) {
	for _, space := range spaces {
		spaceStats := NewSpaceStats(space, skipSIcount)
		c <- spaceStats
	}
	close(c)
}

// NewSpaceStats converts a Space into something decorated with more information
// that can be used for presenting business logic and such
func NewSpaceStats(space Space, skipSIcount bool) SpaceStats {
	totalUniqueApps := space.AppsCount()
	runningUniqueApps := space.RunningAppsCount()
	stoppedUniqueApps := totalUniqueApps - runningUniqueApps

	billableAppInstancesCount := space.BillableAppInstancesCount()
	appInstancesCount := space.AppInstancesCount()
	runningAppInstancesCount := space.RunningAppInstancesCount()
	stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

	consumedMemory := space.ConsumedMemory()

	billableServicesCount := space.BillableServicesCount()
	servicesCount := space.ServicesCount()
	springCloudServicesCount := space.SpringCloudServicesCount()
	servicesSuiteForPivotalPlatformCount := space.ServicesSuiteForPivotalPlatformCount()

	if skipSIcount {
		servicesCount = 0
	}

	return SpaceStats{
		Name:                                 space.Name,
		AppsCount:                            totalUniqueApps,
		RunningAppsCount:                     runningUniqueApps,
		StoppedAppsCount:                     stoppedUniqueApps,
		AppInstancesCount:                    appInstancesCount,
		RunningAppInstancesCount:             runningAppInstancesCount,
		StoppedAppInstancesCount:             stoppedAppInstancesCount,
		ServicesCount:                        servicesCount,
		ConsumedMemory:                       consumedMemory,
		SpringCloudServicesCount:             springCloudServicesCount,
		ServicesSuiteForPivotalPlatformCount: servicesSuiteForPivotalPlatformCount,
		BillableAppInstancesCount:            billableAppInstancesCount,
		BillableServicesCount:                billableServicesCount,
	}

}
