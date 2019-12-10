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
	Spaces              []Space
	SpacesURL           string
}

// OrgsService -
type OrgsService service

// GetOrgByName -
func (o *OrgsService) GetOrgByName(name string) (Org, error) {
	org, err := o.client.cfc.GetOrgByName(name)
	if err != nil {
		return Org{}, err
	}
	spacesURL := fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid)
	return Org{
		GUID:                org.Guid,
		Name:                org.Name,
		QuotaDefinitionGUID: org.QuotaDefinitionGuid,
		SpacesURL:           spacesURL,
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
		spacesURL := fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid)
		orgs = append(orgs,
			Org{
				GUID:                org.Guid,
				Name:                org.Name,
				QuotaDefinitionGUID: org.QuotaDefinitionGuid,
				SpacesURL:           spacesURL,
			})
	}
	return orgs, nil
}

// GetOrgSpacesByOrgGUID returns the spaces in an org using the org's GUID
func (o *OrgsService) GetOrgSpacesByOrgGUID(orgGUID string) ([]Space, error) {
	org, err := o.client.cfc.GetOrgByGuid(orgGUID)
	if err != nil {
		return nil, err
	}

	orgSummary, err := org.Summary()
	if err != nil {
		return nil, err
	}

	spaces := []Space{}
	for _, space := range orgSummary.Spaces {
		summaryURL := fmt.Sprintf("/v2/space/%s/summary", space.Guid)
		spaces = append(spaces,
			Space{
				GUID:       space.Guid,
				Name:       space.Name,
				SummaryURL: summaryURL,
			})
	}
	return spaces, nil
}

// GetOrgMemoryUsageByOrgGUID returns amount of memory (in MB) a given org is currently using
func (o *OrgsService) GetOrgMemoryUsageByOrgGUID(orgGUID string) (float64, error) {
	path := fmt.Sprintf("/v2/organizations/%s/memory_usage", orgGUID)
	usageJSON, err := o.client.Curl(path)
	if err != nil {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}
