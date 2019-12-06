package models

// Org -
type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      []Space
	QuotaURL    string
	Quota       OrgQuota // TODO https://github.com/aegershman/cf-report-usage-plugin/issues/90
	SpacesURL   string
	URL         string
}
