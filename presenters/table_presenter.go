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

	for _, orgStat := range p.Report.OrgDecorators {
		for _, spaceStat := range orgStat.SpaceStats {
			table.Append([]string{
				orgStat.Name,
				spaceStat.Name,
				strconv.Itoa(spaceStat.BillableAppInstancesCount()),
				strconv.Itoa(spaceStat.AppInstancesCount),
				strconv.Itoa(spaceStat.StoppedAppInstancesCount),
				strconv.Itoa(spaceStat.SpringCloudServicesCount()),
			})
		}
	}

	table.SetFooter([]string{
		"-",
		"Total",
		strconv.Itoa(p.Report.AggregateOrgDecorator.BillableAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.AppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.StoppedAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.SpringCloudServicesCount),
	})

	table.Render()

}
