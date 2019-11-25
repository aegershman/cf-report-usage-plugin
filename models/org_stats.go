package models

// Stats -
func (orgs Orgs) Stats(c chan OrgStats) {
	for _, org := range orgs {

		totalUniqueApps := org.AppsCount()
		runningUniqueApps := org.RunningAppsCount()
		stoppedUniqueApps := totalUniqueApps - runningUniqueApps

		appInstancesCount := org.AppInstancesCount()
		runningAppInstancesCount := org.RunningAppInstancesCount()
		stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

		servicesCount := org.ServicesCount()

		c <- OrgStats{
			Name:                     org.Name,
			MemoryQuota:              org.MemoryQuota,
			MemoryUsage:              org.MemoryUsage,
			Spaces:                   org.Spaces,
			AppsCount:                totalUniqueApps,
			RunningAppsCount:         runningUniqueApps,
			StoppedAppsCount:         stoppedUniqueApps,
			AppInstancesCount:        appInstancesCount,
			RunningAppInstancesCount: runningAppInstancesCount,
			StoppedAppInstancesCount: stoppedAppInstancesCount,
			ServicesCount:            servicesCount,
		}
	}
	close(c)
}
