package models

import "strings"

// App -
type App struct {
	Actual int
	Desire int
	RAM    int
}

// Apps -
type Apps []App

// Service -
type Service struct {
	Label       string
	ServicePlan string
}

// Services -
type Services []Service

// Space -
type Space struct {
	Name     string
	Apps     Apps
	Services Services
}

// Spaces -
type Spaces []Space

// ConsumedMemory returns the amount of memory consumed by all
// running canonical application instances within a space
func (space *Space) ConsumedMemory() int {
	count := 0
	for _, app := range space.Apps {
		count += int(app.Actual * app.RAM)
	}
	return count
}

// AppsCount returns the count of unique canonical app guids
// regardless of start/stop state
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "3 unique apps"
func (space *Space) AppsCount() int {
	count := len(space.Apps)
	return count
}

// RunningAppsCount returns the count of unique canonical app
// guids with at least 1 running app instance
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "2 running apps"
func (space *Space) RunningAppsCount() int {
	count := 0
	for _, app := range space.Apps {
		if app.Actual > 0 {
			count++
		}
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (space *Space) BillableAppInstancesCount() int {
	count := 0
	count += space.AppInstancesCount()
	count += space.SpringCloudServicesCount()
	return count
}

// AppInstancesCount returns the count of declared canonical app instances
// regardless of start/stop state
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "5 app instances"
func (space *Space) AppInstancesCount() int {
	count := 0
	for _, app := range space.Apps {
		count += int(app.Desire)
	}
	return count
}

// RunningAppInstancesCount returns the count of declared canonical app instances
// which are actively running
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "4 running app instances"
func (space *Space) RunningAppInstancesCount() int {
	count := 0
	for _, app := range space.Apps {
		count += int(app.Actual)
	}
	return count
}

// ServicesCount returns total count of registered services in the space
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (space *Space) ServicesCount() int {
	count := len(space.Services)
	return count
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (space *Space) BillableServicesCount() int {
	count := space.ServicesCount()
	count -= space.SpringCloudServicesCount()
	return count
}

// ServicesCountByServiceLabel returns the number of service instances
// within a space which contain the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (space *Space) ServicesCountByServiceLabel(serviceType string) int {
	count := 0
	for _, service := range space.Services {
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
func (space *Space) ServicesSuiteForPivotalPlatformCount() int {
	count := 0

	count += space.ServicesCountByServiceLabel("p-dataflow-servers") // TODO

	count += space.ServicesCountByServiceLabel("p-mysql")
	count += space.ServicesCountByServiceLabel("p.mysql")
	count += space.ServicesCountByServiceLabel("pivotal-mysql")

	count += space.ServicesCountByServiceLabel("p-redis")
	count += space.ServicesCountByServiceLabel("p.redis")

	count += space.ServicesCountByServiceLabel("p-rabbitmq")
	count += space.ServicesCountByServiceLabel("p.rabbitmq")

	return count
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (space *Space) SpringCloudServicesCount() int {
	count := 0

	// scs 2.x
	count += space.ServicesCountByServiceLabel("p-config-server")
	count += space.ServicesCountByServiceLabel("p-service-registry")
	count += space.ServicesCountByServiceLabel("p-circuit-breaker")

	// scs 3.x
	count += space.ServicesCountByServiceLabel("p.config-server")
	count += space.ServicesCountByServiceLabel("p.service-registry")

	return count
}
