package v2client

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// Client -
type Client struct {
	cli plugin.CliConnection
}

type service struct {
	client *Client
}

// NewClient -
func NewClient(cli plugin.CliConnection) *Client {
	return &Client{
		cli: cli,
	}
}

// Curl -
func (c *Client) Curl(path string) (map[string]interface{}, error) {
	output, err := c.cli.CliCommandWithoutTerminalOutput("curl", path)
	if err != nil {
		return nil, err
	}

	return parseOutput(output)
}

func parseOutput(output []string) (map[string]interface{}, error) {
	if nil == output || 0 == len(output) {
		return nil, errors.New("CF API returned no output")
	}

	data := strings.Join(output, "\n")

	if 0 == len(data) || "" == data {
		return nil, errors.New("Failed to join output")
	}

	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	return f.(map[string]interface{}), err
}
