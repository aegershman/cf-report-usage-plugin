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

	for _, orgReport := range p.SummaryReport.OrgReports() {
		for _, spaceReport := range orgReport.SpaceReports() {
			table.Append([]string{
				orgReport.Name(),
				spaceReport.Name(),
				strconv.Itoa(spaceReport.BillableAppInstancesCount()),
				strconv.Itoa(spaceReport.AppInstancesCount()),
				strconv.Itoa(spaceReport.StoppedAppInstancesCount()),
				strconv.Itoa(spaceReport.SpringCloudServicesCount()),
			})
		}
	}

	table.SetFooter([]string{
		"-",
		"Total",
		strconv.Itoa(p.SummaryReport.BillableAppInstancesCount()),
		strconv.Itoa(p.SummaryReport.AppInstancesCount()),
		strconv.Itoa(p.SummaryReport.StoppedAppInstancesCount()),
		strconv.Itoa(p.SummaryReport.SpringCloudServicesCount()),
	})

	table.Render()

}
