package apihelper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/aegershman/cf-usage-report-plugin/models"

	"github.com/aegershman/cf-usage-report-plugin/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
)

var (
	// ErrOrgNotFound -
	ErrOrgNotFound = errors.New("organization not found")
)

// CFAPIHelper wraps cf curl results
type CFAPIHelper interface {
	GetTarget() string
	GetOrgs() (models.Orgs, error)
	GetOrg(string) (models.Org, error)
	GetQuotaMemoryLimit(string) (float64, error)
	GetOrgMemoryUsage(models.Org) (float64, error)
	GetOrgSpaces(string) (models.Spaces, error)
	GetSpaceAppsAndServices(string) (models.Apps, models.Services, error)
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
	if nil != err {
		return ""
	}
	apiep, _ := envInfo["routing_endpoint"].(string)
	u, err := url.Parse(apiep)
	if err != nil {
		panic(err)
	}
	host := u.Host
	return host
}

// GetOrgs -
func (api *APIHelper) GetOrgs() (models.Orgs, error) {
	orgsJSON, err := cfcurl.Curl(api.cli, "/v2/organizations")
	if nil != err {
		return nil, err
	}
	pages := int(orgsJSON["total_pages"].(float64))
	orgs := models.Orgs{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			orgsJSON, err = cfcurl.Curl(api.cli, "/v2/organizations?page="+strconv.Itoa(i))
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

// GetOrg -
func (api *APIHelper) GetOrg(name string) (models.Org, error) {
	query := fmt.Sprintf("name:%s", name)
	path := fmt.Sprintf("/v2/organizations?q=%s", url.QueryEscape(query))
	orgsJSON, err := cfcurl.Curl(api.cli, path)
	if nil != err {
		return models.Org{}, err
	}

	results := int(orgsJSON["total_results"].(float64))
	if results == 0 {
		return models.Org{}, ErrOrgNotFound
	}

	orgResource := orgsJSON["resources"].([]interface{})[0]
	org := api.orgResourceToOrg(orgResource)

	return org, nil
}

func (api *APIHelper) orgResourceToOrg(o interface{}) models.Org {
	theOrg := o.(map[string]interface{})
	entity := theOrg["entity"].(map[string]interface{})
	metadata := theOrg["metadata"].(map[string]interface{})
	return models.Org{
		Name:      entity["name"].(string),
		URL:       metadata["url"].(string),
		QuotaURL:  entity["quota_definition_url"].(string),
		SpacesURL: entity["spaces_url"].(string),
	}
}

// GetQuotaMemoryLimit returns memory quota (in MB) of a given org
func (api *APIHelper) GetQuotaMemoryLimit(quotaURL string) (float64, error) {
	quotaJSON, err := cfcurl.Curl(api.cli, quotaURL)
	if nil != err {
		return 0, err
	}
	return quotaJSON["entity"].(map[string]interface{})["memory_limit"].(float64), nil
}

// GetOrgMemoryUsage returns amount of memory (in MB) a given org is currently using
func (api *APIHelper) GetOrgMemoryUsage(org models.Org) (float64, error) {
	usageJSON, err := cfcurl.Curl(api.cli, org.URL+"/memory_usage")
	if nil != err {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}

// GetOrgSpaces returns the spaces in an org
func (api *APIHelper) GetOrgSpaces(spacesURL string) (models.Spaces, error) {
	nextURL := spacesURL
	spaces := models.Spaces{}
	for nextURL != "" {
		spacesJSON, err := cfcurl.Curl(api.cli, nextURL)
		if nil != err {
			return nil, err
		}
		for _, s := range spacesJSON["resources"].([]interface{}) {
			theSpace := s.(map[string]interface{})
			metadata := theSpace["metadata"].(map[string]interface{})
			entity := theSpace["entity"].(map[string]interface{})
			spaces = append(spaces,
				models.Space{
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

// GetSpaceAppsAndServices returns the apps and the services in a space
//
// In reality though, it's just getting the SpaceSummary, right?
//
// This function is problematic because the services it returns are a limited subset
// of the actual services found within a space. We don't want to make those decisions
// here, we'll want to store off the data and make decisions on _how_ and _what_ to render
// somewhere else
//
// Granted, on the other hand, this isn't the "cf go library" so if it makes opinionated
// decisions about what to return it's not the end of the world. But even still, we probably
// don't want to make those decisions here. We'll want to make them in a specific view.
func (api *APIHelper) GetSpaceAppsAndServices(summaryURL string) (models.Apps, models.Services, error) {
	apps := models.Apps{}
	services := models.Services{}
	summaryJSON, err := cfcurl.Curl(api.cli, summaryURL)
	if nil != err {
		return nil, nil, err
	}
	if _, ok := summaryJSON["apps"]; ok {
		for _, a := range summaryJSON["apps"].([]interface{}) {
			theApp := a.(map[string]interface{})
			apps = append(apps,
				models.App{
					RunningInstances: int(theApp["running_instances"].(float64)),
					Instances:        int(theApp["instances"].(float64)),
					Memory:           int(theApp["memory"].(float64)),
				})
		}
	}
	if _, ok := summaryJSON["services"]; ok {
		for _, s := range summaryJSON["services"].([]interface{}) {
			theService := s.(map[string]interface{})
			// TODO I believe us filtering on service plan existing means that
			// user-provided services won't be included in this
			if _, servicePlanExist := theService["service_plan"]; servicePlanExist {
				servicePlan := theService["service_plan"].(map[string]interface{})
				if _, serviceExist := servicePlan["service"]; serviceExist {

					service := servicePlan["service"].(map[string]interface{})
					label := service["label"].(string)

					services = append(services,
						models.Service{
							ServicePlanLabel: label,
							ServicePlanName:  servicePlan["name"].(string),
						})

				}
			}
		}
	}
	return apps, services, nil
}
