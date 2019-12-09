package presentation

import (
	"github.com/aegershman/cf-report-usage-plugin/internal/report"
)

// Presenter -
type Presenter struct {
	SummaryReport report.SummaryReport `json:"summary_report"`
	formats       []string
}

// NewPresenter -
func NewPresenter(r report.SummaryReport, formats ...string) Presenter {
	return Presenter{
		SummaryReport: r,
		formats:       formats,
	}
}

// Render -
func (p *Presenter) Render() {
	for _, format := range p.formats {
		switch format {
		case "json":
			p.asJSON()
			fallthrough
		case "string":
			p.asString()
			fallthrough
		case "table-org-quota": // again, TODO, get rid of this, bleh
			p.asTableOrgQuota()
			fallthrough
		case "table":
			p.asTable()
			fallthrough
		default:
			p.asString()
		}
	}
}
