package presenters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (p *Presenter) asTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{
		"Org",
		"Space",
		"Billable AIs",
		"AIs",
		"Stopped AIs",
		"SCS",
	})

	for _, OrgReport := range p.Report.OrgReports {
		for _, SpaceReport := range OrgReport.SpaceReport {
			table.Append([]string{
				OrgReport.Name,
				SpaceReport.Name,
				strconv.Itoa(SpaceReport.BillableAppInstancesCount()),
				strconv.Itoa(SpaceReport.AppInstancesCount),
				strconv.Itoa(SpaceReport.StoppedAppInstancesCount),
				strconv.Itoa(SpaceReport.SpringCloudServicesCount()),
			})
		}
	}

	table.SetFooter([]string{
		"-",
		"Total",
		strconv.Itoa(p.Report.AggregateOrgReport.BillableAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgReport.AppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgReport.StoppedAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgReport.SpringCloudServicesCount),
	})

	table.Render()

}
