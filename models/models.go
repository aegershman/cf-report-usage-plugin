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
	Name                       string
	DeployedAppsCount          int
	RunningAppsCount           int
	StoppedAppsCount           int
	CanonicalAppInstancesCount int
	DeployedAppInstancesCount  int
	RunningAppInstancesCount   int
	StoppedAppInstancesCount   int
	ServicesCount              int
	ConsumedMemory             int
}

// OrgStats -
type OrgStats struct {
	Name                      string
	MemoryQuota               int
	MemoryUsage               int
	Spaces                    Spaces
	DeployedAppsCount         int
	RunningAppsCount          int
	StoppedAppsCount          int
	DeployedAppInstancesCount int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	ServicesCount             int
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

// InstancesCount -
func (org *Org) InstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.AppInstancesCount()
		SCSCount := space.ServicesCountByServiceLabel("p-spring-cloud-services")
		SCDFCount := space.ServicesCountByServiceLabel("p-dataflow-servers")
		instancesCount += SCSCount + (SCDFCount * 3)
	}
	return instancesCount
}

// RunningAppsCount returns the count of unique canonical app
// guids with at least 1 running app instance across all spaces within the org
func (org *Org) RunningAppsCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningAppsCount()
	}
	return instancesCount
}

// RunningAppInstancesCount returns the count of declared canonical app instances
// which are actively running across all spaces within the org
func (org *Org) RunningAppInstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningAppInstancesCount()
	}
	return instancesCount
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
//
// TODO is this valuable? is this something that should be part of a Space type too?
func (org *Org) AppsCount() int {
	appsCount := 0
	for _, space := range org.Spaces {
		appsCount += len(space.Apps)
	}
	return appsCount
}

// ServicesCount returns total count of registered services in all spaces of the org
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (org *Org) ServicesCount() int {
	servicesCount := 0
	for _, space := range org.Spaces {
		servicesCount += len(space.Services)
	}
	return servicesCount
}

// ConsumedMemory -
func (space *Space) ConsumedMemory() int {
	consumed := 0
	for _, app := range space.Apps {
		consumed += int(app.Actual * app.RAM)
	}
	return consumed
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
	runningAppsCount := 0
	for _, app := range space.Apps {
		if app.Actual > 0 {
			runningAppsCount++
		}
	}
	return runningAppsCount
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
	appInstancesCount := 0
	for _, app := range space.Apps {
		appInstancesCount += int(app.Desire)
	}
	return appInstancesCount
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
	runningAppInstancesCount := 0
	for _, app := range space.Apps {
		runningAppInstancesCount += int(app.Actual)
	}
	return runningAppInstancesCount
}

// ServicesCount returns total count of registered services in the space
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (space *Space) ServicesCount() int {
	servicesCount := len(space.Services)
	return servicesCount
}

// ServicesCountByServiceLabel returns the number of service instances
// within a space which match the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (space *Space) ServicesCountByServiceLabel(serviceType string) int {
	matchingServices := 0
	for _, service := range space.Services {
		if strings.Contains(service.Label, serviceType) {
			matchingServices++
		}
	}
	return matchingServices
}

// Stats -
func (spaces Spaces) Stats(c chan SpaceStats, skipSIcount bool) {
	for _, space := range spaces {
		SCSCount := space.ServicesCountByServiceLabel("p-spring-cloud-services")
		SCDFCount := space.ServicesCountByServiceLabel("p-dataflow-servers")
		lApps := len(space.Apps)
		rApps := space.RunningAppsCount()
		sApps := lApps - rApps
		// "canonical" appInstances are what we can use for setting a quota
		canonicalAppInstances := space.AppInstancesCount()
		// "lAIs" in this context is really "billableAIs", but I don't want to mess
		// with the existing logic before getting a chance to rework this
		lAIs := space.AppInstancesCount()
		lAIs += (SCSCount + (SCDFCount * 3))
		rAIs := space.RunningAppInstancesCount()
		rAIs += (SCSCount + (SCDFCount * 3))
		sAIs := lAIs - rAIs
		siCount := space.ServicesCount()
		siCount -= (SCSCount + SCDFCount)
		rAIConsumedMemory := (space.ConsumedMemory() + (SCSCount * 1024) + (SCDFCount * 3 * 1024))
		if skipSIcount {
			siCount = 0
		}
		c <- SpaceStats{
			Name:                       space.Name,
			DeployedAppsCount:          lApps,
			RunningAppsCount:           rApps,
			StoppedAppsCount:           sApps,
			CanonicalAppInstancesCount: canonicalAppInstances,
			DeployedAppInstancesCount:  lAIs,
			RunningAppInstancesCount:   rAIs,
			StoppedAppInstancesCount:   sAIs,
			ServicesCount:              siCount,
			ConsumedMemory:             rAIConsumedMemory,
		}
	}
	close(c)
}

// Stats -
func (orgs Orgs) Stats(c chan OrgStats) {
	for _, org := range orgs {
		lApps := org.AppsCount()
		rApps := org.RunningAppsCount()
		sApps := lApps - rApps
		lAIs := org.InstancesCount()
		rAIs := org.RunningAppInstancesCount()
		sAIs := lAIs - rAIs
		c <- OrgStats{
			Name:                      org.Name,
			MemoryQuota:               org.MemoryQuota,
			MemoryUsage:               org.MemoryUsage,
			Spaces:                    org.Spaces,
			DeployedAppsCount:         lApps,
			RunningAppsCount:          rApps,
			StoppedAppsCount:          sApps,
			DeployedAppInstancesCount: lAIs,
			RunningAppInstancesCount:  rAIs,
			StoppedAppInstancesCount:  sAIs,
			ServicesCount:             org.ServicesCount(),
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
		spaceBillableAppInstancesMsg  = "\t\t%d billable app instances: %d running, %d stopped\n"
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

			response.WriteString(fmt.Sprintf(spaceCanonicalAppInstancesMsg, spaceState.CanonicalAppInstancesCount))

			response.WriteString(
				fmt.Sprintf(spaceBillableAppInstancesMsg, spaceState.DeployedAppInstancesCount, spaceState.RunningAppInstancesCount, spaceState.StoppedAppInstancesCount))

			response.WriteString(fmt.Sprintf(spaceUniqueAppGuidsMsg, spaceState.DeployedAppsCount, spaceState.RunningAppsCount, spaceState.StoppedAppsCount))

			response.WriteString(fmt.Sprintf(spaceServiceSuiteMsg, spaceState.ServicesCount))

		}
		totalApps += orgStats.DeployedAppsCount
		totalInstances += orgStats.DeployedAppInstancesCount
		totalRunningApps += orgStats.RunningAppsCount
		totalRunningInstances += orgStats.RunningAppInstancesCount
		totalServiceInstances += orgStats.ServicesCount
	}

	response.WriteString(fmt.Sprintf(reportSummaryMsg, totalApps, len(report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	return response.String()
}
