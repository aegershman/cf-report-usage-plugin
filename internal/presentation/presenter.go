package presentation

import (
	"github.com/aegershman/cf-report-usage-plugin/internal/report"
	log "github.com/sirupsen/logrus"
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
	// TODO better handling of defaults
	if len(p.formats) == 0 {
		p.asTable()
	}

	for _, format := range p.formats {
		switch format {
		case "json":
			p.asJSON()
		case "string":
			p.asString()
		case "table-org-quota": // again, TODO, get rid of this, bleh
			p.asTableOrgQuota()
		case "table":
			p.asTable()
		default:
			log.Debugf("could not identify presentation format %s\n", format)
		}
	}
}
