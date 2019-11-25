package models

// Stats -
func (spaces Spaces) Stats(c chan SpaceStats, skipSIcount bool) {
	for _, space := range spaces {

		totalUniqueApps := space.AppsCount()
		runningUniqueApps := space.RunningAppsCount()
		stoppedUniqueApps := totalUniqueApps - runningUniqueApps

		// What _used_ to be reported as just "services"
		servicesSuiteForPivotalPlatformCount := space.ServicesSuiteForPivotalPlatformCount()

		appInstancesCount := space.AppInstancesCount()
		runningAppInstancesCount := space.RunningAppInstancesCount()
		stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

		billableAppInstancesCount := space.AppInstancesCount()
		billableAppInstancesCount += space.SpringCloudServicesCount()

		consumedMemory := space.ConsumedMemory()
		servicesCount := space.ServicesCount()
		billableServicesCount := servicesCount - space.SpringCloudServicesCount()
		if skipSIcount {
			servicesCount = 0
		}

		c <- SpaceStats{
			Name:                                 space.Name,
			AppsCount:                            totalUniqueApps,
			RunningAppsCount:                     runningUniqueApps,
			StoppedAppsCount:                     stoppedUniqueApps,
			AppInstancesCount:                    appInstancesCount,
			RunningAppInstancesCount:             runningAppInstancesCount,
			StoppedAppInstancesCount:             stoppedAppInstancesCount,
			ServicesCount:                        servicesCount,
			ConsumedMemory:                       consumedMemory,
			ServicesSuiteForPivotalPlatformCount: servicesSuiteForPivotalPlatformCount,
			BillableAppInstancesCount:            billableAppInstancesCount,
			BillableServicesCount:                billableServicesCount,
		}
	}
	close(c)
}
