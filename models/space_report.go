package models

import "strings"

// SpaceReport is a way to represent the 'business logic'
// of Spaces; we can use it as a way to decorate
// a Space with extra info like billableAIs, etc.
type SpaceReport struct {
	space                    Space
	Name                     string
	AppsCount                int
	RunningAppsCount         int
	StoppedAppsCount         int
	AppInstancesCount        int
	RunningAppInstancesCount int
	StoppedAppInstancesCount int
	ServicesCount            int
	ConsumedMemory           int
}

// NewSpaceReport -
func NewSpaceReport(space Space) SpaceReport {
	return SpaceReport{
		space:                    space,
		Name:                     space.Name,
		AppsCount:                space.AppsCount(),
		RunningAppsCount:         space.RunningAppsCount(),
		StoppedAppsCount:         space.AppsCount() - space.RunningAppsCount(),
		AppInstancesCount:        space.AppInstancesCount(),
		RunningAppInstancesCount: space.RunningAppInstancesCount(),
		StoppedAppInstancesCount: space.AppInstancesCount() - space.RunningAppInstancesCount(),
		ServicesCount:            space.ServicesCount(),
		ConsumedMemory:           space.ConsumedMemory(),
	}
}

// ServicesCountByServiceLabel returns the number of service instances
// within a space which contain the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (s *SpaceReport) ServicesCountByServiceLabel(serviceType string) int {
	count := 0
	for _, service := range s.space.Services {
		if strings.Contains(service.ServicePlanLabel, serviceType) {
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
func (s *SpaceReport) ServicesSuiteForPivotalPlatformCount() int {
	count := 0

	count += s.ServicesCountByServiceLabel("p-dataflow-servers") // TODO

	count += s.ServicesCountByServiceLabel("p-mysql")
	count += s.ServicesCountByServiceLabel("p.mysql")
	count += s.ServicesCountByServiceLabel("pivotal-mysql")

	count += s.ServicesCountByServiceLabel("p-redis")
	count += s.ServicesCountByServiceLabel("p.redis")

	count += s.ServicesCountByServiceLabel("p-rabbitmq")
	count += s.ServicesCountByServiceLabel("p.rabbitmq")

	return count
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (s *SpaceReport) SpringCloudServicesCount() int {
	count := 0

	// scs 2.x
	count += s.ServicesCountByServiceLabel("p-config-server")
	count += s.ServicesCountByServiceLabel("p-service-registry")
	count += s.ServicesCountByServiceLabel("p-circuit-breaker")

	// scs 3.x
	count += s.ServicesCountByServiceLabel("p.config-server")
	count += s.ServicesCountByServiceLabel("p.service-registry")

	return count
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (s *SpaceReport) BillableServicesCount() int {
	count := s.space.ServicesCount()
	count -= s.SpringCloudServicesCount()
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (s *SpaceReport) BillableAppInstancesCount() int {
	count := 0
	count += s.space.AppInstancesCount()
	count += s.SpringCloudServicesCount()
	return count
}
