package models

import (
	log "github.com/sirupsen/logrus"
)

type Reportable interface {
	Name() string
	MemoryQuota() int
	MemoryUsage() int
	AppsCount() int
	RunningAppsCount() int
	StoppedAppsCount() int
	AppInstancesCount() int
	RunningAppInstancesCount() int
	StoppedAppInstancesCount() int
	ServicesCount() int
	ServicesSuiteForPivotalPlatformCount() int
}

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
func (r *Report) Execute() {

	var aggregateOrgDecorator []OrgDecorator

	aggregateBillableAppInstancesCount := 0
	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0
	aggregateBillableServicesCount := 0

	chOrgDecorators := make(chan OrgDecorator, len(r.Orgs))
	go PopulateOrgDecorators(r.Orgs, chOrgDecorators)
	for orgDecorator := range chOrgDecorators {

		log.WithFields(log.Fields{
			"org": orgDecorator.Name,
		}).Traceln("processing")

		chSpaceDecorators := make(chan SpaceDecorator, len(orgDecorator.Spaces))
		go PopulateSpaceDecorators(orgDecorator.Spaces, chSpaceDecorators)
		for spaceDecorator := range chSpaceDecorators {

			log.WithFields(log.Fields{
				"org":   orgDecorator.Name,
				"space": spaceDecorator.Name,
			}).Traceln("processing")

			orgDecorator.SpaceDecorator = append(orgDecorator.SpaceDecorator, spaceDecorator)

		}

		aggregateBillableAppInstancesCount += orgDecorator.BillableAppInstancesCount()
		aggregateAppInstancesCount += orgDecorator.AppInstancesCount
		aggregateRunningAppInstancesCount += orgDecorator.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += orgDecorator.StoppedAppInstancesCount
		aggregateSpringCloudServicesCount += orgDecorator.SpringCloudServicesCount()
		aggregateBillableServicesCount += orgDecorator.BillableServicesCount()

		aggregateOrgDecorator = append(aggregateOrgDecorator, orgDecorator)

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
