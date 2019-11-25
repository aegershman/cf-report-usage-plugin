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

	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateBillableAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0

	var aggregateOrgStats []OrgStats

	go r.Orgs.Stats(chOrgStats)
	for orgStat := range chOrgStats {
		log.Tracef("processing %s\n", orgStat.Name) // todo just testing
		chSpaceStats := make(chan SpaceStats, len(orgStat.Spaces))
		go orgStat.Spaces.Stats(chSpaceStats, orgStat.Name == "p-spring-cloud-services") // TODO make this more dynamic
		for spaceStat := range chSpaceStats {
			log.Tracef("processing %s\n", spaceStat.Name) // todo just testing
			orgStat.SpaceStats = append(orgStat.SpaceStats, spaceStat)
		}

		aggregateAppInstancesCount += orgStat.AppInstancesCount
		aggregateRunningAppInstancesCount += orgStat.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += orgStat.StoppedAppInstancesCount
		aggregateBillableAppInstancesCount += orgStat.BillableAppInstancesCount
		aggregateSpringCloudServicesCount += orgStat.SpringCloudServicesCount

		aggregateOrgStats = append(aggregateOrgStats, orgStat)
	}

	r.OrgStats = aggregateOrgStats
	r.AggregateOrgStats = AggregateOrgStats{
		AppInstancesCount:         aggregateAppInstancesCount,
		RunningAppInstancesCount:  aggregateRunningAppInstancesCount,
		StoppedAppInstancesCount:  aggregateStoppedAppInstancesCount,
		BillableAppInstancesCount: aggregateBillableAppInstancesCount,
		SpringCloudServicesCount:  aggregateSpringCloudServicesCount,
	}

}
