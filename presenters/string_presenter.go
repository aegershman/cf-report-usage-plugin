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

	for _, orgDecorator := range p.Report.OrgDecorators {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgDecorator.Name, orgDecorator.MemoryUsage, orgDecorator.MemoryQuota))
		for _, spaceDecorator := range orgDecorator.SpaceDecorator {
			if orgDecorator.MemoryQuota > 0 {
				spaceMemoryConsumedPercentage := (100 * spaceDecorator.ConsumedMemory / orgDecorator.MemoryQuota)
				response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceDecorator.Name, spaceDecorator.ConsumedMemory, spaceMemoryConsumedPercentage))
			}
			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceDecorator.BillableAppInstancesCount()))
			response.WriteString(fmt.Sprintf(spaceAppInstancesMsg, spaceDecorator.AppInstancesCount, spaceDecorator.RunningAppInstancesCount, spaceDecorator.StoppedAppInstancesCount))
			response.WriteString(fmt.Sprintf(spaceSCSMsg, spaceDecorator.SpringCloudServicesCount()))
		}
	}

	response.WriteString(
		fmt.Sprintf(
			reportSummaryMsg,
			len(p.Report.Orgs),
			p.Report.AggregateOrgDecorator.BillableAppInstancesCount,
			p.Report.AggregateOrgDecorator.AppInstancesCount,
			p.Report.AggregateOrgDecorator.RunningAppInstancesCount,
			p.Report.AggregateOrgDecorator.StoppedAppInstancesCount,
			p.Report.AggregateOrgDecorator.SpringCloudServicesCount,
		),
	)

	fmt.Println(response.String())
}
