package models

// Space -
type Space struct {
	Name       string
	Apps       []App
	Services   []Service
	SummaryURL string
}
