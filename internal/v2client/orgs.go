package v2client

import (
	"errors"
	"fmt"
	"strconv"
)

// Org -
type Org struct {
	MemoryQuota int
	MemoryUsage int
	Name        string
	QuotaURL    string
	Spaces      []Space
	SpacesURL   string
	URL         string
}

var (
	// ErrOrgNotFound -
	ErrOrgNotFound = errors.New("organization not found")
)

// OrgsService -
type OrgsService service

// GetOrg -
func (o *OrgsService) GetOrg(name string) (Org, error) {
	org, err := o.client.cfc.GetOrgByName(name)
	if err != nil {
		return Org{}, err
	}

	quotaURL := fmt.Sprintf("/v2/quota_definitions/%s", org.QuotaDefinitionGuid)
	spacesURL := fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid)
	url := fmt.Sprintf("/v2/organizations/%s", org.Guid)

	return Org{
		Name:      org.Name,
		QuotaURL:  quotaURL,
		SpacesURL: spacesURL,
		URL:       url,
	}, nil
}

// GetOrgs -
func (o *OrgsService) GetOrgs() ([]Org, error) {
	orgsJSON, err := o.client.Curl("/v2/organizations")
	if err != nil {
		return nil, err
	}
	pages := int(orgsJSON["total_pages"].(float64))
	orgs := []Org{}
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
				Org{
					Name:      name,
					URL:       metadata["url"].(string),
					QuotaURL:  entity["quota_definition_url"].(string),
					SpacesURL: entity["spaces_url"].(string),
				})
		}
	}
	return orgs, nil
}

// GetOrgSpaces returns the spaces in an org
func (o *OrgsService) GetOrgSpaces(spacesURL string) ([]Space, error) {
	nextURL := spacesURL
	spaces := []Space{}
	for nextURL != "" {
		spacesJSON, err := o.client.Curl(nextURL)
		if err != nil {
			return nil, err
		}
		for _, s := range spacesJSON["resources"].([]interface{}) {
			theSpace := s.(map[string]interface{})
			metadata := theSpace["metadata"].(map[string]interface{})
			entity := theSpace["entity"].(map[string]interface{})
			spaces = append(spaces,
				Space{
					Name:       entity["name"].(string),
					SummaryURL: metadata["url"].(string) + "/summary",
				})
		}
		if next, ok := spacesJSON["next_url"].(string); ok {
			nextURL = next
		} else {
			nextURL = ""
		}
	}
	return spaces, nil
}

// GetOrgMemoryUsage returns amount of memory (in MB) a given org is currently using
func (o *OrgsService) GetOrgMemoryUsage(org Org) (float64, error) {
	usageJSON, err := o.client.Curl(org.URL + "/memory_usage")
	if err != nil {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}
