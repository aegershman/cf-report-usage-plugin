package models

import (
	log "github.com/sirupsen/logrus"
)

// type Reportable interface {
// 	Name() string
// 	MemoryQuota() int
// 	MemoryUsage() int
// 	AppsCount() int
// 	RunningAppsCount() int
// 	StoppedAppsCount() int
// 	AppInstancesCount() int
// 	RunningAppInstancesCount() int
// 	StoppedAppInstancesCount() int
// 	ServicesCount() int
// 	ServicesSuiteForPivotalPlatformCount() int
// }

// AggregateOrgDecorator describes an aggregated view
// of multiple OrgReport after a Report Execution run
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
	OrgReports            []OrgReport
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

	var aggregateOrgDecorator []OrgReport

	aggregateBillableAppInstancesCount := 0
	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0
	aggregateBillableServicesCount := 0

	chOrgReports := make(chan OrgReport, len(r.Orgs))
	go PopulateOrgReports(r.Orgs, chOrgReports)
	for OrgReport := range chOrgReports {

		log.WithFields(log.Fields{
			"org": OrgReport.Name,
		}).Traceln("processing")

		chSpaceReports := make(chan SpaceReport, len(OrgReport.Spaces))
		go PopulateSpaceReports(OrgReport.Spaces, chSpaceReports)
		for spaceReport := range chSpaceReports {

			log.WithFields(log.Fields{
				"org":   OrgReport.Name,
				"space": spaceReport.Name,
			}).Traceln("processing")

			OrgReport.SpaceReport = append(OrgReport.SpaceReport, spaceReport)

		}

		aggregateBillableAppInstancesCount += OrgReport.BillableAppInstancesCount()
		aggregateAppInstancesCount += OrgReport.AppInstancesCount
		aggregateRunningAppInstancesCount += OrgReport.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += OrgReport.StoppedAppInstancesCount
		aggregateSpringCloudServicesCount += OrgReport.SpringCloudServicesCount()
		aggregateBillableServicesCount += OrgReport.BillableServicesCount()

		aggregateOrgDecorator = append(aggregateOrgDecorator, OrgReport)

	}

	r.OrgReports = aggregateOrgDecorator
	r.AggregateOrgDecorator = AggregateOrgDecorator{
		BillableAppInstancesCount: aggregateBillableAppInstancesCount,
		BillableServicesCount:     aggregateBillableServicesCount,
		AppInstancesCount:         aggregateAppInstancesCount,
		RunningAppInstancesCount:  aggregateRunningAppInstancesCount,
		StoppedAppInstancesCount:  aggregateStoppedAppInstancesCount,
		SpringCloudServicesCount:  aggregateSpringCloudServicesCount,
	}

}
