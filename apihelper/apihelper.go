package apihelper

import (
	"github.com/cloudfoundry/cli/plugin"
)

// CFAPIHelper wraps cf-curl results, acts as a cf-curl client
type CFAPIHelper interface {
}

// APIHelper -
type APIHelper struct {
	cli plugin.CliConnection
}

// New -
func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}
