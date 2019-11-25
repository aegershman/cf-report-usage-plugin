package presenters

import (
	"bytes"
	"fmt"
)

func (p *Presenter) asString() {
	var response bytes.Buffer

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

	for _, orgStat := range p.Report.OrgStats {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgStat.Name, orgStat.MemoryUsage, orgStat.MemoryQuota))
		for _, spaceStat := range orgStat.SpaceStats {
			if orgStat.MemoryQuota > 0 {
				spaceMemoryConsumedPercentage := (100 * spaceStat.ConsumedMemory / orgStat.MemoryQuota)
				response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceStat.Name, spaceStat.ConsumedMemory, spaceMemoryConsumedPercentage))
			}
			response.WriteString(fmt.Sprintf(spaceAppInstancesMsg, spaceStat.AppInstancesCount, spaceStat.RunningAppInstancesCount, spaceStat.StoppedAppInstancesCount))
			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceStat.BillableAppInstancesCount))
			response.WriteString(fmt.Sprintf(spaceUniqueAppGuidsMsg, spaceStat.AppsCount, spaceStat.RunningAppsCount, spaceStat.StoppedAppsCount))
			response.WriteString(fmt.Sprintf(spaceServiceMsg, spaceStat.ServicesCount))
			response.WriteString(fmt.Sprintf(spaceServiceSuiteMsg, spaceStat.ServicesSuiteForPivotalPlatformCount))
		}
	}

	// response.WriteString(fmt.Sprintf(reportSummaryMsg, totalApps, len(p.Report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	fmt.Println(response.String())
}
