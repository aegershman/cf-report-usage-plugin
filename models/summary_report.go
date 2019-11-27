package models

import "bytes"

// SummaryReport holds an aggregated view of multiple OrgReports
type SummaryReport struct {
	orgReportsRef []OrgReport
	orgsRef       []Org
	Report        Report
}

// NewSummaryReport -
func NewSummaryReport(orgs []Org) *SummaryReport {
	var orgReports []OrgReport
	for _, org := range orgs {
		orgReports = append(orgReports, *NewOrgReport(org))
	}

	return &SummaryReport{
		orgsRef:       orgs,
		orgReportsRef: orgReports,
		Report:        Report{},
	}
}

// OrgReports -
func (s *SummaryReport) OrgReports() []OrgReport {
	return s.orgReportsRef
}

// Name -
func (s *SummaryReport) Name() string {
	var name bytes.Buffer
	for _, org := range s.orgReportsRef {
		name.WriteString(org.Name())
	}
	return name.String()
}

// ServicesSuiteForPivotalPlatformCount returns the number of service instances
// part of the "services suite for pivotal platform", e.g. Pivotal's MySQL/Redis/RMQ
//
// see: https://network.pivotal.io/products/pcf-services
// (I know right? It's an intense function name)
func (s *SummaryReport) ServicesSuiteForPivotalPlatformCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.ServicesSuiteForPivotalPlatformCount()
	}
	return count
}

// AppInstancesCount returns the count of declared canonical app instances
// regardless of start/stop state
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "5 app instances"
func (s *SummaryReport) AppInstancesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.AppInstancesCount()
	}
	return count
}

// AppsCount returns the count of unique canonical app guids
// regardless of start/stop state
//
// for example, within a space, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "3 unique apps"
func (s *SummaryReport) AppsCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.AppsCount()
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (s *SummaryReport) BillableAppInstancesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.BillableAppInstancesCount()
	}
	return count
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (s *SummaryReport) BillableServicesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.BillableServicesCount()
	}
	return count
}

// MemoryQuota -
func (s *SummaryReport) MemoryQuota() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.MemoryQuota()
	}
	return count
}

// MemoryUsage -
func (s *SummaryReport) MemoryUsage() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.MemoryUsage()
	}
	return count
}

// RunningAppInstancesCount returns the count of declared canonical app instances
// which are actively running
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "4 running app instances"
func (s *SummaryReport) RunningAppInstancesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.RunningAppInstancesCount()
	}
	return count
}

// RunningAppsCount returns the count of unique canonical app
// guids with at least 1 running app instance
//
// for example, if you have the following result from `cf apps`:
//
// hammerdb-test                   stopped           0/1
// nodejs-web                      started           2/2
// push-test-webhook-switchboard   started           2/2
//
// then you'd have "2 running apps"
func (s *SummaryReport) RunningAppsCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.RunningAppsCount()
	}
	return count
}

// ServicesCount returns total count of registered services
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (s *SummaryReport) ServicesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.ServicesCount()
	}
	return count
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (s *SummaryReport) SpringCloudServicesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.SpringCloudServicesCount()
	}
	return count
}

// StoppedAppInstancesCount -
func (s *SummaryReport) StoppedAppInstancesCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.StoppedAppInstancesCount()
	}
	return count
}

// StoppedAppsCount -
func (s *SummaryReport) StoppedAppsCount() int {
	count := 0
	for _, report := range s.orgReportsRef {
		count += report.StoppedAppsCount()
	}
	return count

}
