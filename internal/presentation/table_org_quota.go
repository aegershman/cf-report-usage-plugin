package presentation

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// TODO this is ugly as sin, I know, it's temporary
// I just need to get this out for working on a project
func (p *Presenter) asTableOrgQuota() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{
		"Org",
		"Billable AIs",
		"AIs",
		"Stopped AIs",
		"Apps",
		"SCS",
	})

	for _, orgReport := range p.SummaryReport.OrgReports {
		table.Append([]string{
			orgReport.Name,
			strconv.Itoa(orgReport.BillableAppInstancesCount),
			strconv.Itoa(orgReport.AppInstancesCount),
			strconv.Itoa(orgReport.StoppedAppInstancesCount),
			strconv.Itoa(orgReport.AppsCount),
			strconv.Itoa(orgReport.SpringCloudServicesCount),
		})
	}

	table.SetFooter([]string{
		"Total",
		strconv.Itoa(p.SummaryReport.BillableAppInstancesCount),
		strconv.Itoa(p.SummaryReport.AppInstancesCount),
		strconv.Itoa(p.SummaryReport.StoppedAppInstancesCount),
		strconv.Itoa(p.SummaryReport.AppsCount),
		strconv.Itoa(p.SummaryReport.SpringCloudServicesCount),
	})

	table.Render()

}
