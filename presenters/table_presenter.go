package presenters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// TODO just testing, just goofing off, I know this is wrong, etc...
func (p *Presenter) asTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{
		"Org",
		"Space",
		"MB",
		"cAIs",
		"bAIs",
		"cSIs",
		"bSIs",
	})

	for _, orgStat := range p.Report.OrgStats {
		for _, spaceStat := range orgStat.SpaceStats {
			table.Append([]string{
				orgStat.Name,
				spaceStat.Name,
				strconv.Itoa(spaceStat.ConsumedMemory),
				strconv.Itoa(spaceStat.AppInstancesCount),
				strconv.Itoa(spaceStat.BillableAppInstancesCount),
				strconv.Itoa(spaceStat.ServicesCount),
				strconv.Itoa(spaceStat.BillableServicesCount),
			})
		}

		table.Render()

	}

}
