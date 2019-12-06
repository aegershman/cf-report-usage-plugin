package v2client

import (
	"log"
	"net/url"
)

// InfoService -
type InfoService service

// GetTarget -
func (i *InfoService) GetTarget() string {
	envInfo, err := i.client.Curl("/v2/info")
	if err != nil {
		return ""
	}
	apiep, _ := envInfo["routing_endpoint"].(string)
	u, err := url.Parse(apiep)
	if err != nil {
		log.Fatalln(err)
	}
	host := u.Host
	return host
}
