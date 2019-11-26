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

// Report -
type Report struct {
	Orgs               []Org
	OrgReports         []OrgReport
	AggregateOrgReport AggregateOrgReport
}

// NewReport -
func NewReport(orgs []Org) Report {
	return Report{
		Orgs: orgs,
	}
}

// Execute -
func (r *Report) Execute() {

	var aggregateOrgReport []OrgReport

	aggregateBillableAppInstancesCount := 0
	aggregateAppInstancesCount := 0
	aggregateRunningAppInstancesCount := 0
	aggregateStoppedAppInstancesCount := 0
	aggregateSpringCloudServicesCount := 0
	aggregateBillableServicesCount := 0

	chOrgReports := make(chan OrgReport, len(r.Orgs))
	go PopulateOrgReports(r.Orgs, chOrgReports)
	for orgReport := range chOrgReports {

		log.WithFields(log.Fields{
			"org": orgReport.Name,
		}).Traceln("processing")

		chSpaceReports := make(chan SpaceReport, len(orgReport.Spaces))
		go PopulateSpaceReports(orgReport.Spaces, chSpaceReports)
		for spaceReport := range chSpaceReports {

			log.WithFields(log.Fields{
				"org":   orgReport.Name,
				"space": spaceReport.Name,
			}).Traceln("processing")

			orgReport.SpaceReport = append(orgReport.SpaceReport, spaceReport)

		}

		aggregateBillableAppInstancesCount += orgReport.BillableAppInstancesCount()
		aggregateAppInstancesCount += orgReport.AppInstancesCount
		aggregateRunningAppInstancesCount += orgReport.RunningAppInstancesCount
		aggregateStoppedAppInstancesCount += orgReport.StoppedAppInstancesCount
		aggregateSpringCloudServicesCount += orgReport.SpringCloudServicesCount()
		aggregateBillableServicesCount += orgReport.BillableServicesCount()

		aggregateOrgReport = append(aggregateOrgReport, orgReport)

	}

	r.OrgReports = aggregateOrgReport
	r.AggregateOrgReport = AggregateOrgReport{
		BillableAppInstancesCount: aggregateBillableAppInstancesCount,
		BillableServicesCount:     aggregateBillableServicesCount,
		AppInstancesCount:         aggregateAppInstancesCount,
		RunningAppInstancesCount:  aggregateRunningAppInstancesCount,
		StoppedAppInstancesCount:  aggregateStoppedAppInstancesCount,
		SpringCloudServicesCount:  aggregateSpringCloudServicesCount,
	}

}
