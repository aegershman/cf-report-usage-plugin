package report

import (
	"github.com/aegershman/cf-report-usage-plugin/v2client"
	"github.com/cloudfoundry/cli/plugin"
)

// Reporter -
type Reporter struct {
	orgsRef []v2client.Org
	cli     plugin.CliConnection
}

// NewReporter -
func NewReporter(cli plugin.CliConnection, orgNames []string) {
}
