package v2client

import (
	"net/url"
)

// InfoService -
type InfoService service

// GetTarget -
func (i *InfoService) GetTarget() (string, error) {
	envInfo, err := i.client.Curl("/v2/info")
	if err != nil {
		return "", err
	}
	apiep, _ := envInfo["routing_endpoint"].(string)
	u, err := url.Parse(apiep)
	if err != nil {
		return "", err
	}
	host := u.Host
	return host, nil
}
