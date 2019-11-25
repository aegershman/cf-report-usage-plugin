package models

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

	servicesSuiteForPivotalPlatformCount := space.ServicesSuiteForPivotalPlatformCount()

	appInstancesCount := space.AppInstancesCount()
	runningAppInstancesCount := space.RunningAppInstancesCount()
	stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

	springCloudServicesCount := space.SpringCloudServicesCount()

	billableAppInstancesCount := space.BillableAppInstancesCount()

	consumedMemory := space.ConsumedMemory()
	servicesCount := space.ServicesCount()
	billableServicesCount := servicesCount - space.SpringCloudServicesCount()
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
