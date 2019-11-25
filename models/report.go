package models

// NewReport -
func NewReport(orgs Orgs) Report {
	return Report{
		Orgs: orgs,
	}
}

// Execute -
func (r *Report) Execute() {
	chOrgStats := make(chan OrgStats, len(r.Orgs))

	// vars we'll use to track the aggregation of Org/Space Stats as they
	// come in off the channels
	var (
		aggregatedOrgStats   []OrgStats
		aggregatedSpaceStats []SpaceStats
	)

	go r.Orgs.Stats(chOrgStats)
	for orgStat := range chOrgStats {
		chSpaceStats := make(chan SpaceStats, len(orgStat.Spaces))
		go orgStat.Spaces.Stats(chSpaceStats, orgStat.Name == "p-spring-cloud-services")
		for spaceStat := range chSpaceStats {
			aggregatedSpaceStats = append(aggregatedSpaceStats, spaceStat)
		}
		aggregatedOrgStats = append(aggregatedOrgStats, orgStat)
	}

	r.orgStats = aggregatedOrgStats
	r.spaceStats = aggregatedSpaceStats

}
