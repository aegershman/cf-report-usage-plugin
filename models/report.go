package models

import (
	log "github.com/sirupsen/logrus"
)

// Report -
type Report struct {
	Orgs               []Org
	OrgReports         []OrgReport
	AggregateOrgReport AggregateOrgReport
}

// NewReport -
func NewReport(orgs []Org) Report {
	var orgReports []OrgReport
	for _, org := range orgs {
		orgReports = append(orgReports, NewOrgReport(org))
	}

	return Report{
		Orgs:       orgs,
		OrgReports: orgReports,
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

	for _, orgReport := range r.OrgReports {

		log.WithFields(log.Fields{
			"org": orgReport.Name,
		}).Traceln("processing")

		for _, spaceReport := range orgReport.SpaceReport {

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
