package presenters

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func (p *Presenter) asJSON() {
	json, err := json.Marshal(p.SummaryReporter.SummaryReportRef())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(json))
	log.Fatalln("TODO")
}
