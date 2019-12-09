package presentation

import (
	"github.com/aegershman/cf-report-usage-plugin/internal/report"
	log "github.com/sirupsen/logrus"
)

const (
	DEFAULT_PRESENTER_FORMAT = "table"
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
		p.render(format)
	}
}

func (p *Presenter) render(format string) {
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
		p.render(DEFAULT_PRESENTER_FORMAT)
	}

}
