package presenters

import (
	"github.com/aegershman/cf-report-usage-plugin/models"
)

// Presenter -
type Presenter struct {
	SummaryReport models.SummaryReport `json:"summary_report"`
	Format        string               `json:"format"`
}

// NewPresenter -
func NewPresenter(r models.SummaryReport, format string) Presenter {
	return Presenter{
		SummaryReport: r,
		Format:        format,
	}
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "string":
		p.asString()
	case "table":
		p.asTable()
	case "json":
		p.asJSON()
	case "garbage":
		p.asGarbage()
	default:
		// TODO
		// yeah this is kind of awful I know, I'm sorry, I'm still learning,
		// I'll fix this along with much better and earlier error handling on this
		// I'll fix this, I promise
		p.asString()
	}
}
