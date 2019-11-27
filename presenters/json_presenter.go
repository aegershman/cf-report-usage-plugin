package presenters

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (p *Presenter) asJSON() {
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(j))
}
