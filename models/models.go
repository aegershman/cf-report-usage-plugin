package models

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      []Space
}

type Space struct {
	Name string
	Apps []App
	Services []Service
}

//App representation
type App struct {
	Actual	int
	Desire	int
	RAM     int
}

//Service representation
type Service struct {
	Label    	string
	ServicePlan string
}

type Report struct {
	Orgs []Org
}

func (org *Org) InstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.InstancesCount()
	}
	return instancesCount
}

func (org *Org) RunningAppsCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningAppsCount()
	}
	return instancesCount
}

func (org *Org) RunningInstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningInstancesCount()
	}
	return instancesCount
}

func (org *Org) AppsCount() int {
	appsCount := 0
	for _, space := range org.Spaces {
		appsCount += len(space.Apps)
	}
	return appsCount
}

func (org *Org) ServicesCount() int {
	servicesCount := 0
	for _, space := range org.Spaces {
		servicesCount += len(space.Services)
	}
	return servicesCount
}

func (space *Space) ConsumedMemory() int {
	consumed := 0
	for _, app := range space.Apps {
		consumed += int(app.Actual * app.RAM)
	}
	return consumed
}

func (space *Space) RunningAppsCount() int {
	runningAppsCount := 0
	for _, app := range space.Apps {
		if (app.Actual > 0) {
			runningAppsCount++
		}
	}
	return runningAppsCount
}

func (space *Space) InstancesCount() int {
	instancesCount := 0
	for _, app := range space.Apps {
		instancesCount += int(app.Desire)
	}
	return instancesCount
}

func (space *Space) RunningInstancesCount() int {
	runningInstancesCount := 0
	for _, app := range space.Apps {
		runningInstancesCount += int(app.Actual)
	}
	return runningInstancesCount
}

func (space *Space) ServicesCount() int {
	servicesCount := len(space.Services)
	return servicesCount
}

func (report *Report) String() string {
	var response bytes.Buffer

	totalApps := 0
	totalInstances := 0
	totalRunningApps := 0
	totalRunningInstances := 0
	totalServiceInstances := 0

	for _, org := range report.Orgs {
		response.WriteString(fmt.Sprintf("Org %s is consuming %d MB of %d MB.\n",
			org.Name, org.MemoryUsage, org.MemoryQuota))

		for _, space := range org.Spaces {
			spaceRunningAppsCount := space.RunningAppsCount()
			spaceInstancesCount := space.InstancesCount()
			spaceServiceInstancesCount := space.ServicesCount()
			spaceRunningInstancesCount := space.RunningInstancesCount()
			spaceConsumedMemory := space.ConsumedMemory()

			response.WriteString(
				fmt.Sprintf("\tSpace %s is consuming %d MB memory (%d%%) of org quota.\n",
					space.Name, spaceConsumedMemory, (100 * spaceConsumedMemory / org.MemoryQuota)))
			response.WriteString(
				fmt.Sprintf("\t\t%d apps: %d running %d stopped\n", len(space.Apps),
					spaceRunningAppsCount, len(space.Apps)-spaceRunningAppsCount))
			response.WriteString(
				fmt.Sprintf("\t\t%d app instances: %d running, %d stopped\n", spaceInstancesCount,
					spaceRunningInstancesCount, spaceInstancesCount-spaceRunningInstancesCount))
			response.WriteString(
				fmt.Sprintf("\t\t%d service instances of type Service Suite\n", spaceServiceInstancesCount))
		}

		totalApps += org.AppsCount()
		totalInstances += org.InstancesCount()
		totalRunningApps += org.RunningAppsCount()
		totalRunningInstances += org.RunningInstancesCount()
		totalServiceInstances += org.ServicesCount()
	}

	response.WriteString(
		fmt.Sprintf("You have deployed %d apps across %d org(s), with a total of %d app instances configured. You are currently running %d apps with %d app instances and using %d service instances of type Service Suite.\n",
			totalApps, len(report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	return response.String()
}

func (report *Report) CSV() string {
	var rows = [][]string{}
	var csv bytes.Buffer

	var headers = []string{"OrgName", "SpaceName", "SpaceMemoryUsed", "OrgMemoryQuota", "AppsDeployed", "AppsRunning", "AppInstancesDeployed", "AppInstancesRunning", "ServiceInstanceDeployed"}

	rows = append(rows, headers)

	for _, org := range report.Orgs {
		for _, space := range org.Spaces {
			appsDeployed := len(space.Apps)

			spaceResult := []string{
				org.Name,
				space.Name,
				strconv.Itoa(space.ConsumedMemory()),
				strconv.Itoa(org.MemoryQuota),
				strconv.Itoa(appsDeployed),
				strconv.Itoa(space.RunningAppsCount()),
				strconv.Itoa(space.InstancesCount()),
				strconv.Itoa(space.RunningInstancesCount()),
				strconv.Itoa(space.ServicesCount()),
			}

			rows = append(rows, spaceResult)
		}
	}

	for i := range rows {
		csv.WriteString(strings.Join(rows[i], ", "))
		csv.WriteString("\n")
	}

	return csv.String()
}
