package models

import (
	"bytes"
	"fmt"
	"strings"
)

// Org -
type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      Spaces
}

// Space -
type Space struct {
	Name     string
	Apps     Apps
	Services Services
}

// App -
type App struct {
	Actual int
	Desire int
	RAM    int
}

// Service -
type Service struct {
	Label       string
	ServicePlan string
}

// SpaceStats is a way to represent the 'business logic'
// of Spaces; we can use it as a way to decorate
// a Space with extra info like billableAIs, etc.
type SpaceStats struct {
	Name                     string
	AppsCount                int
	RunningAppsCount         int
	StoppedAppsCount         int
	AppInstancesCount        int
	RunningAppInstancesCount int
	StoppedAppInstancesCount int
	ServicesCount            int // TODO misnomer
	ConsumedMemory           int

	ServicesSuiteForPivotalPlatformCount int // TODO

	// includes anything which Pivotal deems "billable" as an AI, even if CF
	// considers it a service; e.g., SCS instances (config server, service registry, etc.)
	BillableAppInstancesCount int

	// count of anything which Pivotal deems "billable" as an SI; this might mean
	// subtracting certain services (like SCS) from the count of `cf services`
	BillableServicesCount int
}

// OrgStats -
type OrgStats struct {
	Name                     string
	MemoryQuota              int
	MemoryUsage              int
	Spaces                   Spaces
	AppsCount                int
	RunningAppsCount         int
	StoppedAppsCount         int
	AppInstancesCount        int
	RunningAppInstancesCount int
	StoppedAppInstancesCount int
	ServicesCount            int

	ServicesSuiteForPivotalPlatformCount int // TODO

	// includes anything which Pivotal deems "billable" as an AI, even if CF
	// considers it a service; e.g., SCS instances (config server, service registry, etc.)
	BillableAppInstancesCount int

	// count of anything which Pivotal deems "billable" as an SI; this might mean
	// subtracting certain services (like SCS) from the count of `cf services`
	BillableServicesCount int
}

// Orgs -
type Orgs []Org

// Spaces -
type Spaces []Space

// Apps -
type Apps []App

// Services -
type Services []Service

// Report -
type Report struct {
	Orgs Orgs
}

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

// Stats -
func (orgs Orgs) Stats(c chan OrgStats) {
	for _, org := range orgs {

		totalUniqueApps := org.AppsCount()
		runningUniqueApps := org.RunningAppsCount()
		stoppedUniqueApps := totalUniqueApps - runningUniqueApps

		lAIs := org.AppInstancesCount()
		rAIs := org.RunningAppInstancesCount()
		sAIs := lAIs - rAIs
		c <- OrgStats{
			Name:                     org.Name,
			MemoryQuota:              org.MemoryQuota,
			MemoryUsage:              org.MemoryUsage,
			Spaces:                   org.Spaces,
			AppsCount:                totalUniqueApps,
			RunningAppsCount:         runningUniqueApps,
			StoppedAppsCount:         stoppedUniqueApps,
			AppInstancesCount:        lAIs,
			RunningAppInstancesCount: rAIs,
			StoppedAppInstancesCount: sAIs,
			ServicesCount:            org.ServicesCount(),
		}
	}
	close(c)
}

func (report *Report) String() string {
	var response bytes.Buffer

	totalApps := 0
	totalInstances := 0
	totalRunningApps := 0
	totalRunningInstances := 0
	totalServiceInstances := 0

	const (
		orgOverviewMsg                = "Org %s is consuming %d MB of %d MB.\n"
		spaceOverviewMsg              = "\tSpace %s is consuming %d MB memory (%d%%) of org quota.\n"
		spaceCanonicalAppInstancesMsg = "\t\t%d canonical app instances\n"
		spaceBillableAppInstancesMsg  = "\t\t%d billable app instances (includes AIs and billable SIs, like SCS)\n"
		spaceUniqueAppGuidsMsg        = "\t\t%d unique app_guids: %d running %d stopped\n"
		spaceServiceSuiteMsg          = "\t\t%d service instances of type Service Suite (mysql, redis, rmq)\n"
		reportSummaryMsg              = "[WARNING: THIS REPORT SUMMARY IS MISLEADING AND INCORRECT. IT WILL BE FIXED SOON.] You have deployed %d apps across %d org(s), with a total of %d app instances configured. You are currently running %d apps with %d app instances and using %d service instances of type Service Suite.\n"
	)

	chOrgStats := make(chan OrgStats, len(report.Orgs))

	go report.Orgs.Stats(chOrgStats)
	for orgStats := range chOrgStats {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgStats.Name, orgStats.MemoryUsage, orgStats.MemoryQuota))
		chSpaceStats := make(chan SpaceStats, len(orgStats.Spaces))
		// TODO dynamic filtering?
		go orgStats.Spaces.Stats(chSpaceStats, orgStats.Name == "p-spring-cloud-services")
		for spaceState := range chSpaceStats {

			// handle org having "zero quota", e.g. the org is only allowed to use service instances, not push apps
			if orgStats.MemoryQuota > 0 {
				spaceMemoryConsumedPercentage := (100 * spaceState.ConsumedMemory / orgStats.MemoryQuota)
				response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceState.Name, spaceState.ConsumedMemory, spaceMemoryConsumedPercentage))
			}

			response.WriteString(fmt.Sprintf(spaceCanonicalAppInstancesMsg, spaceState.AppInstancesCount))

			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceState.BillableAppInstancesCount))

			response.WriteString(fmt.Sprintf(spaceUniqueAppGuidsMsg, spaceState.AppsCount, spaceState.RunningAppsCount, spaceState.StoppedAppsCount))

			response.WriteString(fmt.Sprintf(spaceServiceSuiteMsg, spaceState.ServicesSuiteForPivotalPlatformCount))

		}
		totalApps += orgStats.AppsCount
		totalInstances += orgStats.AppInstancesCount
		totalRunningApps += orgStats.RunningAppsCount
		totalRunningInstances += orgStats.RunningAppInstancesCount
		totalServiceInstances += orgStats.ServicesCount
	}

	response.WriteString(fmt.Sprintf(reportSummaryMsg, totalApps, len(report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	return response.String()
}
