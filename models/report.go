package models

import (
	log "github.com/sirupsen/logrus"
)

// NewReport -
func NewReport(orgs Orgs) Report {
	return Report{
		Orgs: orgs,
	}
}

// Execute -
//
// Since "[]SpaceStats" are a property of every individual "OrgStats"
// within "[]OrgStats" (whew), we make sure that "aggregateSpaceStats"
// below only exists within the context of a given "OrgStats".
// Then we aggregate together all the "OrgStats" for the Report
func (r *Report) Execute() {
	chOrgStats := make(chan OrgStats, len(r.Orgs))

	var aggregateOrgStats []OrgStats

	go r.Orgs.Stats(chOrgStats)
	for orgStat := range chOrgStats {
		log.Debugf("processing %s\n", orgStat.Name) // todo just testing
		chSpaceStats := make(chan SpaceStats, len(orgStat.Spaces))
		// TODO make this more dynamic
		go orgStat.Spaces.Stats(chSpaceStats, orgStat.Name == "p-spring-cloud-services")
		for spaceStat := range chSpaceStats {
			log.Debugf("processing %s\n", spaceStat.Name) // todo just testing
			orgStat.SpaceStats = append(orgStat.SpaceStats, spaceStat)
		}
		aggregateOrgStats = append(aggregateOrgStats, orgStat)
	}

	r.OrgStats = aggregateOrgStats

}
