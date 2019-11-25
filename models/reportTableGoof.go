package models

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Stringg -
func (report *Report) Stringg() {

	chOrgStats := make(chan OrgStats, len(report.Orgs))

	go report.Orgs.Stats(chOrgStats)
	for orgStats := range chOrgStats {

		chSpaceStats := make(chan SpaceStats, len(orgStats.Spaces))

		// TODO dynamic filtering?
		go orgStats.Spaces.Stats(chSpaceStats, orgStats.Name == "p-spring-cloud-services")

		// TODO just testing, just goofing off, I know this is wrong, etc...
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"Space", "MB", "cAIs", "bAIs", "cSIs", "bSIs"})

		for spaceState := range chSpaceStats {
			table.Append([]string{
				spaceState.Name,
				strconv.Itoa(spaceState.ConsumedMemory),
				strconv.Itoa(spaceState.AppInstancesCount),
				strconv.Itoa(spaceState.BillableAppInstancesCount),
				strconv.Itoa(spaceState.ServicesCount),
				strconv.Itoa(spaceState.BillableServicesCount),
			})
		}

		table.Render()

	}

}
