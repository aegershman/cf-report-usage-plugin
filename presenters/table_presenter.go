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
		strconv.Itoa(p.Report.AggregateOrgDecorators.BillableAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorators.AppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorators.StoppedAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorators.SpringCloudServicesCount),
	})

	table.Render()

}
