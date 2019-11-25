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

		for spaceState := range chSpaceStats {

			// TODO just testing, just goofing off, I know this is wrong, etc...

			table := tablewriter.NewWriter(os.Stdout)
			table.SetAlignment(tablewriter.ALIGN_LEFT)

			table.SetHeader([]string{"Name", "MB Consumed", "App Instances"})

			table.Append([]string{spaceState.Name, strconv.Itoa(spaceState.ConsumedMemory), strconv.Itoa(spaceState.AppInstancesCount)})

			table.Render()

		}

	}

}
