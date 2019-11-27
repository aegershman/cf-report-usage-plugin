package presenters

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func (p *Presenter) asJSON() {
	j, err := json.Marshal(p.SummaryReport)
	if err != nil {
		log.Fatalln(err)
	}

	log.Infoln(string(j))
}
