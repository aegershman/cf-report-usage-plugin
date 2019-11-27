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

	for _, orgReport := range p.Report.SummaryReport.OrgReports() {
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
		strconv.Itoa(p.Report.SummaryReport.BillableAppInstancesCount()),
		strconv.Itoa(p.Report.SummaryReport.AppInstancesCount()),
		strconv.Itoa(p.Report.SummaryReport.StoppedAppInstancesCount()),
		strconv.Itoa(p.Report.SummaryReport.SpringCloudServicesCount()),
	})

	table.Render()

}
