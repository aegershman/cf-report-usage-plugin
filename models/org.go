package models

// AppInstancesCount returns the count of declared canonical app instances
// regardless of start/stop state across all spaces within the org
func (org *Org) AppInstancesCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += space.AppInstancesCount()
	}
	return count
}

// RunningAppsCount returns the count of unique canonical app
// guids with at least 1 running app instance across all spaces within the org
func (org *Org) RunningAppsCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += space.RunningAppsCount()
	}
	return count
}

// RunningAppInstancesCount returns the count of declared canonical app instances
// which are actively running across all spaces within the org
func (org *Org) RunningAppInstancesCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += space.RunningAppInstancesCount()
	}
	return count
}

// AppsCount returns the count of unique canonical app guids
// regardless of start/stop state across all spaces within the org
//
// for example, within a space, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "3 unique apps"
func (org *Org) AppsCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += len(space.Apps)
	}
	return count
}

// ServicesCount returns total count of registered services in all spaces of the org
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (org *Org) ServicesCount() int {
	count := 0
	for _, space := range org.Spaces {
		count += len(space.Services)
	}
	return count
}

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
