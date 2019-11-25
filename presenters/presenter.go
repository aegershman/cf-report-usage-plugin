package presenters

import (
	m "github.com/aegershman/cf-trueup-plugin/models"
)

// Presenter -
type Presenter struct {
	Report m.Report
	Format string
}

// NewPresenter -
func NewPresenter(r m.Report, format string) Presenter {
	return Presenter{
		Report: r,
		Format: format,
	}
}

// Render -
func (p *Presenter) Render() {
	switch p.Format {
	case "string":
		p.AsString()
	case "table":
		p.AsTable()
	default:
		// TODO
		// yeah this is kind of awful I know, I'm sorry, I'm still learning,
		// I'll fix this along with much better and earlier error handling on this
		// I'll fix this, I promise
		p.AsString()
	}
}
