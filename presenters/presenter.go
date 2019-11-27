package presenters

import (
	"github.com/aegershman/cf-usage-report-plugin/models"
)

// Presenter -
type Presenter struct {
	SummaryReporter models.SummaryReporter
	Format          string
}

// NewPresenter -
func NewPresenter(r models.SummaryReporter, format string) Presenter {
	return Presenter{
		SummaryReporter: r,
		Format:          format,
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
	default:
		// TODO
		// yeah this is kind of awful I know, I'm sorry, I'm still learning,
		// I'll fix this along with much better and earlier error handling on this
		// I'll fix this, I promise
		p.asString()
	}
}
