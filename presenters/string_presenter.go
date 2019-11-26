package presenters

import (
	"bytes"
	"fmt"
)

func (p *Presenter) asString() {
	var response bytes.Buffer

	const (
		orgOverviewMsg               = "org %s is consuming %d MB of %d MB\n"
		spaceOverviewMsg             = "\tspace %s is consuming %d MB memory (%d%%) of org quota\n"
		spaceBillableAppInstancesMsg = "\t\tAIs billable: %d\n"
		spaceAppInstancesMsg         = "\t\tAIs canonical: %d (%d running, %d stopped)\n"
		spaceSCSMsg                  = "\t\tSCS instances: %d\n"
		reportSummaryMsg             = "across %d org(s), you have %d billable AIs, %d are canonical AIs (%d running, %d stopped), %d are SCS instances\n"
	)

	for _, orgStat := range p.Report.OrgDecorators {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgStat.Name, orgStat.MemoryUsage, orgStat.MemoryQuota))
		for _, spaceStat := range orgStat.SpaceStats {
			if orgStat.MemoryQuota > 0 {
				spaceMemoryConsumedPercentage := (100 * spaceStat.ConsumedMemory / orgStat.MemoryQuota)
				response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceStat.Name, spaceStat.ConsumedMemory, spaceMemoryConsumedPercentage))
			}
			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceStat.BillableAppInstancesCount()))
			response.WriteString(fmt.Sprintf(spaceAppInstancesMsg, spaceStat.AppInstancesCount, spaceStat.RunningAppInstancesCount, spaceStat.StoppedAppInstancesCount))
			response.WriteString(fmt.Sprintf(spaceSCSMsg, spaceStat.SpringCloudServicesCount()))
		}
	}

	response.WriteString(
		fmt.Sprintf(
			reportSummaryMsg,
			len(p.Report.Orgs),
			p.Report.AggregateOrgDecorators.BillableAppInstancesCount,
			p.Report.AggregateOrgDecorators.AppInstancesCount,
			p.Report.AggregateOrgDecorators.RunningAppInstancesCount,
			p.Report.AggregateOrgDecorators.StoppedAppInstancesCount,
			p.Report.AggregateOrgDecorators.SpringCloudServicesCount,
		),
	)

	fmt.Println(response.String())
}
