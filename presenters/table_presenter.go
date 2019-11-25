package presenters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// TODO just testing, just goofing off, I know this is wrong, etc...
func (p *Presenter) asTable() {

	for _, orgStats := range p.Report.OrgStats {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"Space", "MB", "cAIs", "bAIs", "cSIs", "bSIs"})

		for _, spaceStat := range orgStats.SpaceStats {
			table.Append([]string{
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
