package models

import (
	log "github.com/sirupsen/logrus"
)

// SummaryReporter -
type SummaryReporter interface {
	Reporter
	OrgReports() []OrgReport
}

// func (*s SummaryReport) AppInstancesCount() int {}

// SummaryReport describes an aggregated view
// of multiple OrgReport after a Report Execution run
//
// will probably get rid of this at some point
type SummaryReport struct {
	AppInstancesCount         int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	BillableAppInstancesCount int
	SpringCloudServicesCount  int
	BillableServicesCount     int
	orgs                      []Org
	OrgReports                []OrgReport
}

// NewSummaryReport -
func NewSummaryReport(orgs []Org) SummaryReport {
	var orgReports []OrgReport
	for _, org := range orgs {

		// this really isn't even that helpful, it's mostly
		// here as a little example for myself TODO
		log.WithFields(log.Fields{
			"orgReport": org.Name,
		}).Traceln("processing")

		orgReports = append(orgReports, NewOrgReport(org))
	}

	return SummaryReport{
		orgs:       orgs,
		OrgReports: orgReports,
	}
}
