package presentation

import (
	"github.com/aegershman/cf-report-usage-plugin/internal/report"
)

// Presenter -
type Presenter struct {
	SummaryReport report.SummaryReport `json:"summary_report"`
	Format        string               `json:"format"`
}

// NewPresenter -
func NewPresenter(r report.SummaryReport, format string) Presenter {
	return Presenter{
		SummaryReport: r,
		Format:        format,
	}
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "json":
		p.asJSON()
	case "string":
		p.asString()
	case "table-org-quota": // again, TODO, bleh
		p.asTableOrgQuota()
	case "table":
		p.asTable()
	default:
		// TODO
		// yeah this is kind of awful I know, I'm sorry, I'm still learning,
		// I'll fix this along with much better and earlier error handling on this
		// I'll fix this, I promise
		p.asString()
	}
}
