package presenters

import (
	m "github.com/aegershman/cf-trueup-plugin/models"
)

// Presenter -
type Presenter struct {
	Report m.Report
}

// NewPresenter -
func NewPresenter(r m.Report) Presenter {
	return Presenter{
		Report: r,
	}
}
