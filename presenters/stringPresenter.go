package presenters

import (
	"bytes"
	"fmt"

	m "github.com/aegershman/cf-trueup-plugin/models"
)

// AsString -
func (p *Presenter) AsString() {
	var response bytes.Buffer

	totalApps := 0
	totalInstances := 0
	totalRunningApps := 0
	totalRunningInstances := 0
	totalServiceInstances := 0

	const (
		orgOverviewMsg               = "Org %s is consuming %d MB of %d MB.\n"
		spaceOverviewMsg             = "\tSpace %s is consuming %d MB memory (%d%%) of org quota.\n"
		spaceAppInstancesMsg         = "\t\t%d app instances: %d running %d stopped\n"
		spaceBillableAppInstancesMsg = "\t\t%d billable app instances (includes AIs and billable SIs, like SCS)\n"
		spaceUniqueAppGuidsMsg       = "\t\t%d unique app_guids: %d running %d stopped\n"
		spaceServiceMsg              = "\t\t%d service instances total\n"
		spaceServiceSuiteMsg         = "\t\t%d service instances of type Service Suite (mysql, redis, rmq)\n"
		reportSummaryMsg             = "[WARNING: THIS REPORT SUMMARY IS MISLEADING AND INCORRECT. IT WILL BE FIXED SOON.] You have deployed %d apps across %d org(s), with a total of %d app instances configured. You are currently running %d apps with %d app instances and using %d service instances of type Service Suite.\n"
	)

	chOrgStats := make(chan m.OrgStats, len(p.Report.Orgs))

	go p.Report.Orgs.Stats(chOrgStats)
	for orgStats := range chOrgStats {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgStats.Name, orgStats.MemoryUsage, orgStats.MemoryQuota))
		chSpaceStats := make(chan m.SpaceStats, len(orgStats.Spaces))
		// TODO dynamic filtering?
		go orgStats.Spaces.Stats(chSpaceStats, orgStats.Name == "p-spring-cloud-services")
		for spaceState := range chSpaceStats {

			// handle org having "zero quota", e.g. the org is only allowed to use service instances, not push apps
			if orgStats.MemoryQuota > 0 {
				spaceMemoryConsumedPercentage := (100 * spaceState.ConsumedMemory / orgStats.MemoryQuota)
				response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceState.Name, spaceState.ConsumedMemory, spaceMemoryConsumedPercentage))
			}

			response.WriteString(fmt.Sprintf(spaceAppInstancesMsg, spaceState.AppInstancesCount, spaceState.RunningAppInstancesCount, spaceState.StoppedAppInstancesCount))

			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceState.BillableAppInstancesCount))

			response.WriteString(fmt.Sprintf(spaceUniqueAppGuidsMsg, spaceState.AppsCount, spaceState.RunningAppsCount, spaceState.StoppedAppsCount))

			response.WriteString(fmt.Sprintf(spaceServiceMsg, spaceState.ServicesCount))

			response.WriteString(fmt.Sprintf(spaceServiceSuiteMsg, spaceState.ServicesSuiteForPivotalPlatformCount))

		}
		totalApps += orgStats.AppsCount
		totalInstances += orgStats.AppInstancesCount
		totalRunningApps += orgStats.RunningAppsCount
		totalRunningInstances += orgStats.RunningAppInstancesCount
		totalServiceInstances += orgStats.ServicesCount
	}

	response.WriteString(fmt.Sprintf(reportSummaryMsg, totalApps, len(p.Report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	fmt.Println(response.String())
}
