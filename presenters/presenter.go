package presenters

import (
	"github.com/aegershman/cf-usage-report-plugin/models"
)

// Presenter -
type Presenter struct {
	Report models.Report
	Format string
}

// NewPresenter -
func NewPresenter(r models.Report, format string) Presenter {
	return Presenter{
		Report: r,
		Format: format,
	}
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "string":
		p.asString()
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
