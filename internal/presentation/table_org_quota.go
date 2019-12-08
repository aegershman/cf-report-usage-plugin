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
		"AI Quota",
		"Deployed AIs",
		"Running AIs",
		"Stopped AIs",
	})

	for _, orgReport := range p.SummaryReport.OrgReports {
		table.Append([]string{
			orgReport.Name,
			strconv.Itoa(orgReport.OrgQuota.AppInstanceLimit),
			strconv.Itoa(orgReport.AppInstancesCount),
			strconv.Itoa(orgReport.RunningAppInstancesCount),
			strconv.Itoa(orgReport.StoppedAppInstancesCount),
		})
	}

	table.SetFooter([]string{
		"Total",
		"-",
		strconv.Itoa(p.SummaryReport.AppInstancesCount),
		strconv.Itoa(p.SummaryReport.RunningAppInstancesCount),
		strconv.Itoa(p.SummaryReport.StoppedAppInstancesCount),
	})

	table.Render()

}
