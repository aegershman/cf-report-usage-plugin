package presenters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (p *Presenter) asGarbage() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{
		"Org",
		"Billable canonical AIs", // doesn't include SCS, etc.
		"Running AIs",
		"AI Limit",
	})

	for _, orgReport := range p.SummaryReport.OrgReports {
		table.Append([]string{
			orgReport.Name,
			strconv.Itoa(orgReport.AppInstancesCount),
			strconv.Itoa(orgReport.RunningAppInstancesCount),
			strconv.Itoa(orgReport.Quota.AppInstanceLimit),
		})
	}

	table.Render()

}
