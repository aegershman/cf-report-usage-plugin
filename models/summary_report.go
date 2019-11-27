package models

// SummaryReporter -
type SummaryReporter interface {
	OrgReports() []OrgReport
	Reporter
}

// SummaryReport holds an aggregated view of multiple OrgReports
type SummaryReport struct {
	orgsRef       []Org
	orgReportsRef []OrgReport
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
	}
}

func (s *SummaryReport) Name() string {
	return "nil"
}

// ServicesSuiteForPivotalPlatformCount returns the number of service instances
// part of the "services suite for pivotal platform", e.g. Pivotal's MySQL/Redis/RMQ
//
// see: https://network.pivotal.io/products/pcf-services
// (I know right? It's an intense function name)
func (s *SummaryReport) ServicesSuiteForPivotalPlatformCount() int {
	return 0
}

func (s *SummaryReport) OrgReports() []OrgReport {
	return s.orgReportsRef
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
	return 0
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
	return 0
}

// BillableAppInstancesCount returns the count of "billable" AIs
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (s *SummaryReport) BillableAppInstancesCount() int {
	return 0
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (s *SummaryReport) BillableServicesCount() int {
	return 0
}

func (s *SummaryReport) MemoryQuota() int {
	return 0
}

func (s *SummaryReport) MemoryUsage() int {
	return 0
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
	return 0
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
	return 0
}

// ServicesCount returns total count of registered services
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (s *SummaryReport) ServicesCount() int {
	return 0
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (s *SummaryReport) SpringCloudServicesCount() int {
	return 0
}

func (s *SummaryReport) StoppedAppInstancesCount() int {
	return 0
}

func (s *SummaryReport) StoppedAppsCount() int {
	return 0
}
