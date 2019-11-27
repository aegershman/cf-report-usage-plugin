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

func (s *SummaryReport) BillableAppInstancesCount() int {
	return 0
}

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

func (s *SummaryReport) SpringCloudServicesCount() int {
	return 0
}

func (s *SummaryReport) StoppedAppInstancesCount() int {
	return 0
}

func (s *SummaryReport) StoppedAppsCount() int {
	return 0
}
