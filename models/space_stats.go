package models

import "strings"

// SpaceStats is a way to represent the 'business logic'
// of Spaces; we can use it as a way to decorate
// a Space with extra info like billableAIs, etc.
type SpaceStats struct {
	Space                     Space
	Name                      string
	AppsCount                 int
	RunningAppsCount          int
	StoppedAppsCount          int
	AppInstancesCount         int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	ServicesCount             int
	ConsumedMemory            int
	BillableAppInstancesCount int
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
// func NewSpaceStats(space Space, skipSIcount bool) SpaceStats {
// 	totalUniqueApps := space.AppsCount()
// 	runningUniqueApps := space.RunningAppsCount()
// 	stoppedUniqueApps := totalUniqueApps - runningUniqueApps

// 	billableAppInstancesCount := space.BillableAppInstancesCount()
// 	appInstancesCount := space.AppInstancesCount()
// 	runningAppInstancesCount := space.RunningAppInstancesCount()
// 	stoppedAppInstancesCount := appInstancesCount - runningAppInstancesCount

// 	consumedMemory := space.ConsumedMemory()

// 	billableServicesCount := space.BillableServicesCount()
// 	servicesCount := space.ServicesCount()
// 	springCloudServicesCount := space.SpringCloudServicesCount()
// 	servicesSuiteForPivotalPlatformCount := space.ServicesSuiteForPivotalPlatformCount()

// 	if skipSIcount {
// 		servicesCount = 0
// 	}

// 	return SpaceStats{
// 		Name:                      space.Name,
// 		AppsCount:                 totalUniqueApps,
// 		RunningAppsCount:          runningUniqueApps,
// 		StoppedAppsCount:          stoppedUniqueApps,
// 		AppInstancesCount:         appInstancesCount,
// 		RunningAppInstancesCount:  runningAppInstancesCount,
// 		StoppedAppInstancesCount:  stoppedAppInstancesCount,
// 		ServicesCount:             servicesCount,
// 		ConsumedMemory:            consumedMemory,
// 		SpringCloudServicesCount:  springCloudServicesCount,
// 		BillableAppInstancesCount: billableAppInstancesCount,
// 		BillableServicesCount:     billableServicesCount,
// 	}
// }

func NewSpaceStats(space Space, skipSIcount bool) SpaceStats {
	if skipSIcount {
		//
	}

	return SpaceStats{
		Space: space,
	}
}

// ServicesCountByServiceLabel returns the number of service instances
// within a space which contain the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (ss *SpaceStats) ServicesCountByServiceLabel(serviceType string) int {
	count := 0
	for _, service := range ss.Space.Services {
		if strings.Contains(service.Label, serviceType) {
			count++
		}
	}
	return count
}

// ServicesSuiteForPivotalPlatformCount returns the number of service instances
// part of the "services suite for pivotal platform", e.g. Pivotal's MySQL/Redis/RMQ
//
// see: https://network.pivotal.io/products/pcf-services
// (I know right? It's an intense function name)
func (ss *SpaceStats) ServicesSuiteForPivotalPlatformCount() int {
	count := 0

	count += ss.ServicesCountByServiceLabel("p-dataflow-servers") // TODO

	count += ss.ServicesCountByServiceLabel("p-mysql")
	count += ss.ServicesCountByServiceLabel("p.mysql")
	count += ss.ServicesCountByServiceLabel("pivotal-mysql")

	count += ss.ServicesCountByServiceLabel("p-redis")
	count += ss.ServicesCountByServiceLabel("p.redis")

	count += ss.ServicesCountByServiceLabel("p-rabbitmq")
	count += ss.ServicesCountByServiceLabel("p.rabbitmq")

	return count
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (ss *SpaceStats) SpringCloudServicesCount() int {
	count := 0

	// scs 2.x
	count += ss.ServicesCountByServiceLabel("p-config-server")
	count += ss.ServicesCountByServiceLabel("p-service-registry")
	count += ss.ServicesCountByServiceLabel("p-circuit-breaker")

	// scs 3.x
	count += ss.ServicesCountByServiceLabel("p.config-server")
	count += ss.ServicesCountByServiceLabel("p.service-registry")

	return count
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (ss *SpaceStats) BillableServicesCount() int {
	count := ss.Space.ServicesCount()
	count -= ss.SpringCloudServicesCount()
	return count
}
