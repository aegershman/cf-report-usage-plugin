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
		"AIs",
		"Stopped AIs",
		"Billable AIs",
		"Billable SIs",
	})

	for _, orgStat := range p.Report.OrgStats {
		for _, spaceStat := range orgStat.SpaceStats {
			table.Append([]string{
				orgStat.Name,
				spaceStat.Name,
				strconv.Itoa(spaceStat.AppInstancesCount),
				strconv.Itoa(spaceStat.StoppedAppInstancesCount),
				strconv.Itoa(spaceStat.BillableAppInstancesCount()),
				strconv.Itoa(spaceStat.BillableServicesCount()),
			})
		}
	}

	table.SetFooter([]string{
		"-",
		"Total",
		strconv.Itoa(p.Report.AggregateOrgStats.AppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgStats.StoppedAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgStats.BillableAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgStats.BillableServicesCount),
	})

	table.Render()

}
