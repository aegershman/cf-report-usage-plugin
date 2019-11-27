package presenters

import (
	"bytes"
	"fmt"
)

func (p *Presenter) asString() {
	var response bytes.Buffer

	const (
		orgOverviewMsg               = "org %s is consuming %d MB of %d MB\n"
		spaceOverviewMsg             = "\tspace %s is consuming %d MB memory of org quota\n"
		spaceBillableAppInstancesMsg = "\t\tAIs billable: %d\n"
		spaceAppInstancesMsg         = "\t\tAIs canonical: %d (%d running, %d stopped)\n"
		spaceSCSMsg                  = "\t\tSCS instances: %d\n"
		reportSummaryMsg             = "across %d org(s), you have %d billable AIs, %d are canonical AIs (%d running, %d stopped), %d are SCS instances\n"
	)

	for _, orgReport := range p.SummaryReporter.OrgReports() {
		response.WriteString(fmt.Sprintf(orgOverviewMsg, orgReport.Name(), orgReport.MemoryUsage(), orgReport.MemoryQuota()))
		for _, spaceReport := range orgReport.SpaceReports() {
			response.WriteString(fmt.Sprintf(spaceOverviewMsg, spaceReport.Name(), spaceReport.MemoryUsage()))
			response.WriteString(fmt.Sprintf(spaceBillableAppInstancesMsg, spaceReport.BillableAppInstancesCount()))
			response.WriteString(fmt.Sprintf(spaceAppInstancesMsg, spaceReport.AppInstancesCount(), spaceReport.RunningAppInstancesCount(), spaceReport.StoppedAppInstancesCount()))
			response.WriteString(fmt.Sprintf(spaceSCSMsg, spaceReport.SpringCloudServicesCount()))
		}
	}

	response.WriteString(
		fmt.Sprintf(
			reportSummaryMsg,
			len(p.SummaryReporter.OrgReports()),
			p.SummaryReporter.BillableAppInstancesCount(),
			p.SummaryReporter.AppInstancesCount(),
			p.SummaryReporter.RunningAppInstancesCount(),
			p.SummaryReporter.StoppedAppInstancesCount(),
			p.SummaryReporter.SpringCloudServicesCount(),
		),
	)

	fmt.Println(response.String())
}
