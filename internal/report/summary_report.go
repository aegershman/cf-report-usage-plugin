package report

import (
	"github.com/aegershman/cf-report-usage-plugin/internal/v2client"

	"bytes"
)

// SummaryReport holds an aggregated view of multiple OrgReports
// It effectively serves as the entrypoint into aggregating the data
// in preparation for it being presented
type SummaryReport struct {
	OrgReports []OrgReport `json:"org_reports"`
	Report
}

// NewSummaryReport -
func NewSummaryReport(orgs []v2client.Org) *SummaryReport {
	var orgReports []OrgReport
	for _, org := range orgs {
		orgReports = append(orgReports, *NewOrgReport(org))
	}

	self := &SummaryReport{
		OrgReports: orgReports,
	}

	self.Report = Report{
		AppInstancesCount:                    self.appInstancesCount(),
		AppsCount:                            self.appsCount(),
		BillableAppInstancesCount:            self.billableAppInstancesCount(),
		BillableServicesCount:                self.billableServicesCount(),
		MemoryQuota:                          self.memoryQuota(),
		MemoryUsage:                          self.memoryUsage(),
		Name:                                 self.name(),
		RunningAppInstancesCount:             self.runningAppInstancesCount(),
		RunningAppsCount:                     self.runningAppsCount(),
		ServicesCount:                        self.servicesCount(),
		ServicesSuiteForPivotalPlatformCount: self.servicesSuiteForPivotalPlatformCount(),
		SpringCloudServicesCount:             self.springCloudServicesCount(),
		StoppedAppInstancesCount:             self.stoppedAppInstancesCount(),
		StoppedAppsCount:                     self.stoppedAppsCount(),
	}

	return self
}

// Name -
func (s *SummaryReport) name() string {
	var name bytes.Buffer
	for _, org := range s.OrgReports {
		name.WriteString(org.name())
	}
	return name.String()
}

// ServicesSuiteForPivotalPlatformCount returns the number of service instances
// part of the "services suite for pivotal platform", e.g. Pivotal's MySQL/Redis/RMQ
//
// see: https://network.pivotal.io/products/pcf-services
// (I know right? It's an intense function name)
func (s *SummaryReport) servicesSuiteForPivotalPlatformCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.servicesSuiteForPivotalPlatformCount()
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
func (s *SummaryReport) appInstancesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.appInstancesCount()
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
func (s *SummaryReport) appsCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.appsCount()
	}
	return count
}

// BillableAppInstancesCount returns the count of "billable" AIs
//
// This includes anything which Pivotal deems "billable" as an AI, even if CF
// considers it a service; e.g., SCS instances (config server, service registry, etc.)
func (s *SummaryReport) billableAppInstancesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.billableAppInstancesCount()
	}
	return count
}

// BillableServicesCount returns the count of "billable" SIs
//
// This includes anything which Pivotal deems "billable" as an SI; this might mean
// subtracting certain services (like SCS) from the count of `cf services`
func (s *SummaryReport) billableServicesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.billableServicesCount()
	}
	return count
}

// MemoryQuota -
func (s *SummaryReport) memoryQuota() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.memoryQuota()
	}
	return count
}

// MemoryUsage -
func (s *SummaryReport) memoryUsage() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.memoryUsage()
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
func (s *SummaryReport) runningAppInstancesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.runningAppInstancesCount()
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
func (s *SummaryReport) runningAppsCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.runningAppsCount()
	}
	return count
}

// ServicesCount returns total count of registered services
//
// Keep in mind, if a single service ends up creating more service instances
// (or application instances) in a different space (e.g., Spring Cloud Data Flow, etc.)
// those aren't considered in this result. This only counts services registered which
// show up in `cf services`
func (s *SummaryReport) servicesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.servicesCount()
	}
	return count
}

// SpringCloudServicesCount returns the number of service instances
// from "spring cloud services" tile, e.g. config-server/service-registry/circuit-breaker/etc.
//
// see: https://network.pivotal.io/products/p-spring-cloud-services/
func (s *SummaryReport) springCloudServicesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.springCloudServicesCount()
	}
	return count
}

// StoppedAppInstancesCount -
func (s *SummaryReport) stoppedAppInstancesCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.stoppedAppInstancesCount()
	}
	return count
}

// StoppedAppsCount -
func (s *SummaryReport) stoppedAppsCount() int {
	count := 0
	for _, report := range s.OrgReports {
		count += report.stoppedAppsCount()
	}
	return count

}
