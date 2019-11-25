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

	aggregateBillableAppInstancesCount := 0
	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0
	aggregateBillableServicesCount := 0

	var aggregateOrgStats []OrgStats

	go r.Orgs.Stats(chOrgStats)
	for orgStat := range chOrgStats {

		log.WithFields(log.Fields{
			"org": orgStat.Name,
		}).Traceln("processing")

		chSpaceStats := make(chan SpaceStats, len(orgStat.Spaces))
		go orgStat.Spaces.Stats(chSpaceStats, orgStat.Name == "p-spring-cloud-services") // TODO make this more dynamic
		for spaceStat := range chSpaceStats {

			log.WithFields(log.Fields{
				"org":   orgStat.Name,
				"space": spaceStat.Name,
			}).Traceln("processing")

			orgStat.SpaceStats = append(orgStat.SpaceStats, spaceStat)

		}

		aggregateBillableAppInstancesCount += orgStat.BillableAppInstancesCount
		aggregateAppInstancesCount += orgStat.AppInstancesCount
		aggregateRunningAppInstancesCount += orgStat.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += orgStat.StoppedAppInstancesCount
		aggregateSpringCloudServicesCount += orgStat.SpringCloudServicesCount
		aggregateBillableServicesCount += orgStat.BillableServicesCount

		aggregateOrgStats = append(aggregateOrgStats, orgStat)

	}

	r.OrgStats = aggregateOrgStats
	r.AggregateOrgStats = AggregateOrgStats{
		BillableAppInstancesCount: aggregateBillableAppInstancesCount,
		BillableServicesCount:     aggregateBillableServicesCount,
		AppInstancesCount:         aggregateAppInstancesCount,
		RunningAppInstancesCount:  aggregateRunningAppInstancesCount,
		StoppedAppInstancesCount:  aggregateStoppedAppInstancesCount,
		SpringCloudServicesCount:  aggregateSpringCloudServicesCount,
	}

}
