package v2client

import (
	"fmt"
)

// Org -
type Org struct {
	GUID                string
	MemoryQuota         int
	MemoryUsage         int
	Name                string
	QuotaDefinitionGUID string
	QuotaURL            string
	Spaces              []Space
	SpacesURL           string
	URL                 string
}

// OrgsService -
type OrgsService service

// GetOrgByName -
func (o *OrgsService) GetOrgByName(name string) (Org, error) {
	org, err := o.client.cfc.GetOrgByName(name)
	if err != nil {
		return Org{}, err
	}

	quotaURL := fmt.Sprintf("/v2/quota_definitions/%s", org.QuotaDefinitionGuid)
	spacesURL := fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid)
	url := fmt.Sprintf("/v2/organizations/%s", org.Guid)

	return Org{
		GUID:                org.Guid,
		Name:                org.Name,
		QuotaDefinitionGUID: org.QuotaDefinitionGuid,
		QuotaURL:            quotaURL,
		SpacesURL:           spacesURL,
		URL:                 url,
	}, nil
}

// GetOrgs -
func (o *OrgsService) GetOrgs() ([]Org, error) {
	listedOrgs, err := o.client.cfc.ListOrgs()
	if err != nil {
		return nil, err
	}

	orgs := []Org{}
	for _, org := range listedOrgs {
		quotaURL := fmt.Sprintf("/v2/quota_definitions/%s", org.QuotaDefinitionGuid)
		spacesURL := fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid)
		url := fmt.Sprintf("/v2/organizations/%s", org.Guid)
		orgs = append(orgs,
			Org{
				GUID:                org.Guid,
				Name:                org.Name,
				QuotaDefinitionGUID: org.QuotaDefinitionGuid,
				QuotaURL:            quotaURL,
				SpacesURL:           spacesURL,
				URL:                 url,
			})
	}

	return orgs, nil
}

// GetOrgSpacesByOrgGUID returns the spaces in an org using the org's GUID
func (o *OrgsService) GetOrgSpacesByOrgGUID(orgGUID string) ([]Space, error) {
	nextURL := fmt.Sprintf("/v2/organizations/%s/spaces", orgGUID)
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
					GUID:       metadata["guid"].(string),
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
