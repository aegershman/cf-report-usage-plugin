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

func (s *SummaryReport) AppInstancesCount() int {
	return 0
}

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

func (s *SummaryReport) RunningAppInstancesCount() int {
	return 0
}

func (s *SummaryReport) RunningAppsCount() int {
	return 0
}

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
