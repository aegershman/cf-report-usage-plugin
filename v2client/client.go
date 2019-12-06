package v2client

import (
	"github.com/cloudfoundry/cli/plugin"
)

// Client -
type Client struct {
	cli plugin.CliConnection
}

// NewClient -
func NewClient(cli plugin.CliConnection) *Client {
	return &Client{
		cli: cli,
	}
}

type service struct {
	client *Client
}
