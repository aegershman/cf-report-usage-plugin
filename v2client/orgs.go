package v2client

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/aegershman/cf-report-usage-plugin/models"
)

var (
	// ErrOrgNotFound -
	ErrOrgNotFound = errors.New("organization not found")
)

// OrgsService -
type OrgsService service

// GetOrg -
func (o *OrgsService) GetOrg(name string) (models.Org, error) {
	query := fmt.Sprintf("name:%s", name)
	path := fmt.Sprintf("/v2/organizations?q=%s", url.QueryEscape(query))
	orgsJSON, err := o.client.Curl(path)
	if err != nil {
		return models.Org{}, err
	}

	results := int(orgsJSON["total_results"].(float64))
	if results == 0 {
		return models.Org{}, ErrOrgNotFound
	}

	orgResource := orgsJSON["resources"].([]interface{})[0]
	theOrg := orgResource.(map[string]interface{})
	entity := theOrg["entity"].(map[string]interface{})
	metadata := theOrg["metadata"].(map[string]interface{})

	return models.Org{
		Name:      entity["name"].(string),
		URL:       metadata["url"].(string),
		QuotaURL:  entity["quota_definition_url"].(string),
		SpacesURL: entity["spaces_url"].(string),
	}, nil
}

// GetOrgs -
func (o *OrgsService) GetOrgs() ([]models.Org, error) {
	orgsJSON, err := o.client.Curl("/v2/organizations")
	if err != nil {
		return nil, err
	}
	pages := int(orgsJSON["total_pages"].(float64))
	orgs := []models.Org{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			orgsJSON, err = o.client.Curl("/v2/organizations?page=" + strconv.Itoa(i))
		}
		for _, o := range orgsJSON["resources"].([]interface{}) {
			theOrg := o.(map[string]interface{})
			entity := theOrg["entity"].(map[string]interface{})
			name := entity["name"].(string)
			if name == "system" {
				continue
			}
			metadata := theOrg["metadata"].(map[string]interface{})
			orgs = append(orgs,
				models.Org{
					Name:      name,
					URL:       metadata["url"].(string),
					QuotaURL:  entity["quota_definition_url"].(string),
					SpacesURL: entity["spaces_url"].(string),
				})
		}
	}
	return orgs, nil
}
