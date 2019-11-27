package models

// Report -
type Report struct {
	Orgs          []Org
	SummaryReport SummaryReport
}

// NewReport -
func NewReport(orgs []Org) Report {
	return Report{
		Orgs:          orgs,
		SummaryReport: NewSummaryReport(orgs),
	}
}
