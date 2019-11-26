package models

import (
	log "github.com/sirupsen/logrus"
)

// AggregateOrgDecorator describes an aggregated view
// of multiple OrgDecorator after a Report Execution run
type AggregateOrgDecorator struct {
	AppInstancesCount         int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	BillableAppInstancesCount int
	SpringCloudServicesCount  int
	BillableServicesCount     int
}

// Report -
// TODO consider breaking into "pre-init" and "post-init" structs,
// e.g. "reportPlan" and "report"? Possibly makes it clearer that you're
// supposed to "execute" the reportPlan to get it to generate the data?
type Report struct {
	Orgs                  []Org
	OrgDecorators         []OrgDecorator
	AggregateOrgDecorator AggregateOrgDecorator
}

// NewReport -
func NewReport(orgs []Org) Report {
	return Report{
		Orgs: orgs,
	}
}

// Execute -
//
// Since "[]SpaceStats" are a property of every individual "OrgDecorator"
// within "[]OrgDecorator" (whew), we make sure that "aggregateSpaceStats"
// below only exists within the context of a given "OrgDecorator".
// Then we aggregate together all the "OrgDecorator" for the Report
func (r *Report) Execute() {

	var aggregateOrgDecorator []OrgDecorator

	aggregateBillableAppInstancesCount := 0
	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0
	aggregateBillableServicesCount := 0

	chOrgStats := make(chan OrgDecorator, len(r.Orgs))
	go NewOrgsStats(r.Orgs, chOrgStats)
	for orgStat := range chOrgStats {

		log.WithFields(log.Fields{
			"org": orgStat.Name,
		}).Traceln("processing")

		chSpaceStats := make(chan SpaceStats, len(orgStat.Spaces))
		go NewSpacesStats(orgStat.Spaces, chSpaceStats)
		for spaceStat := range chSpaceStats {

			log.WithFields(log.Fields{
				"org":   orgStat.Name,
				"space": spaceStat.Name,
			}).Traceln("processing")

			orgStat.SpaceStats = append(orgStat.SpaceStats, spaceStat)

		}

		aggregateBillableAppInstancesCount += orgStat.BillableAppInstancesCount()
		aggregateAppInstancesCount += orgStat.AppInstancesCount
		aggregateRunningAppInstancesCount += orgStat.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += orgStat.StoppedAppInstancesCount
		aggregateSpringCloudServicesCount += orgStat.SpringCloudServicesCount()
		aggregateBillableServicesCount += orgStat.BillableServicesCount()

		aggregateOrgDecorator = append(aggregateOrgDecorator, orgStat)

	}

	r.OrgDecorators = aggregateOrgDecorator
	r.AggregateOrgDecorator = AggregateOrgDecorator{
		BillableAppInstancesCount: aggregateBillableAppInstancesCount,
		BillableServicesCount:     aggregateBillableServicesCount,
		AppInstancesCount:         aggregateAppInstancesCount,
		RunningAppInstancesCount:  aggregateRunningAppInstancesCount,
		StoppedAppInstancesCount:  aggregateStoppedAppInstancesCount,
		SpringCloudServicesCount:  aggregateSpringCloudServicesCount,
	}

}
