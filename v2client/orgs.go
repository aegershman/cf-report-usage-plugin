package v2client

import (
	"strconv"

	"github.com/aegershman/cf-report-usage-plugin/models"
)

// OrgsService -
type OrgsService service

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
