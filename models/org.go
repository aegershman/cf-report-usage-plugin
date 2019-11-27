package models

// Org -
type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      []Space
	QuotaURL    string
	SpacesURL   string
	URL         string
}
