package apihelper

import (
	"net/url"

	"github.com/aegershman/cf-report-usage-plugin/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

// CFAPIHelper wraps cf-curl results, acts as a cf-curl client
type CFAPIHelper interface {
	GetTarget() string
}

// APIHelper -
type APIHelper struct {
	cli plugin.CliConnection
}

// New -
func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

// GetTarget -
func (api *APIHelper) GetTarget() string {
	envInfo, err := cfcurl.Curl(api.cli, "/v2/info")
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
