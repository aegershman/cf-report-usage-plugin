package apihelper

import (
	"errors"
	"fmt"
	"strings"
	"net/url"
	"strconv"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/cfcurl"
)

var (
	ErrOrgNotFound = errors.New("organization not found")
)

//Organization representation
type Organization struct {
	URL       string
	Name      string
	QuotaURL  string
	SpacesURL string
}

//Space representation
type Space struct {
	Name    	string
	SummaryURL	string
}

//App representation
type App struct {
	Actual	float64
	Desire	float64
	RAM     float64
}

//Service representation
type Service struct {
	Label    	string
	ServicePlan string
}

//CFAPIHelper to wrap cf curl results
type CFAPIHelper interface {
	GetOrgs() ([]Organization, error)
	GetOrg(string) (Organization, error)
	GetQuotaMemoryLimit(string) (float64, error)
	GetOrgMemoryUsage(Organization) (float64, error)
	GetOrgSpaces(string) ([]Space, error)
	GetSpaceAppsAndServices(string) ([]App, []Service, error)
}

//APIHelper implementation
type APIHelper struct {
	cli plugin.CliConnection
}

func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

//GetOrgs returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetOrgs() ([]Organization, error) {
	orgsJSON, err := cfcurl.Curl(api.cli, "/v2/organizations")
	if nil != err {
		return nil, err
	}
	pages := int(orgsJSON["total_pages"].(float64))
	orgs := []Organization{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			orgsJSON, err = cfcurl.Curl(api.cli, "/v2/organizations?page="+strconv.Itoa(i))
		}
		for _, o := range orgsJSON["resources"].([]interface{}) {
			theOrg := o.(map[string]interface{})
			entity := theOrg["entity"].(map[string]interface{})
			name := entity["name"].(string)
			if (name == "system") {
				continue
			}
			metadata := theOrg["metadata"].(map[string]interface{})
			orgs = append(orgs,
				Organization{
					Name:      name,
					URL:       metadata["url"].(string),
					QuotaURL:  entity["quota_definition_url"].(string),
					SpacesURL: entity["spaces_url"].(string),
				})
		}
	}
	return orgs, nil
}

//GetOrg returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetOrg(name string) (Organization, error) {
	query := fmt.Sprintf("name:%s", name)
	path := fmt.Sprintf("/v2/organizations?q=%s&inline-relations-depth=1", url.QueryEscape(query))
	orgsJSON, err := cfcurl.Curl(api.cli, path)
	if nil != err {
		return Organization{}, err
	}

	results := int(orgsJSON["total_results"].(float64))
	if results == 0 {
		return Organization{}, ErrOrgNotFound
	}

	orgResource := orgsJSON["resources"].([]interface{})[0]
	org := api.orgResourceToOrg(orgResource)

	return org, nil
}

func (api *APIHelper) orgResourceToOrg(o interface{}) Organization {
	theOrg := o.(map[string]interface{})
	entity := theOrg["entity"].(map[string]interface{})
	metadata := theOrg["metadata"].(map[string]interface{})
	return Organization{
		Name:      entity["name"].(string),
		URL:       metadata["url"].(string),
		QuotaURL:  entity["quota_definition_url"].(string),
		SpacesURL: entity["spaces_url"].(string),
	}
}

//GetQuotaMemoryLimit retruns the amount of memory (in MB) that the org is allowed
func (api *APIHelper) GetQuotaMemoryLimit(quotaURL string) (float64, error) {
	quotaJSON, err := cfcurl.Curl(api.cli, quotaURL)
	if nil != err {
		return 0, err
	}
	return quotaJSON["entity"].(map[string]interface{})["memory_limit"].(float64), nil
}

//GetOrgMemoryUsage returns the amount of memory (in MB) that the org is consuming
func (api *APIHelper) GetOrgMemoryUsage(org Organization) (float64, error) {
	usageJSON, err := cfcurl.Curl(api.cli, org.URL+"/memory_usage")
	if nil != err {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}

//GetOrgSpaces returns the spaces in an org.
func (api *APIHelper) GetOrgSpaces(spacesURL string) ([]Space, error) {
	nextURL := spacesURL
	spaces := []Space{}
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
				Space{
					Name:    entity["name"].(string),
					SummaryURL: metadata["url"].(string)+"/summary",
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

//GetSpaceAppsAndServices returns the apps and the services in a space
func (api *APIHelper) GetSpaceAppsAndServices(summaryURL string) ([]App, []Service, error) {
	apps := []App{}
	services := []Service{}
	summaryJSON, err := cfcurl.Curl(api.cli, summaryURL)
	if nil != err {
		return nil, nil, err
	}

	_, aok := summaryJSON["apps"]
	_, sok := summaryJSON["services"]

	if(aok) {
		for _, a := range summaryJSON["apps"].([]interface{}) {
			theApp := a.(map[string]interface{})
			apps = append(apps,
				App{
					Actual: theApp["running_instances"].(float64),
					Desire: theApp["instances"].(float64),
					RAM:	theApp["memory"].(float64),
				})
		}
	}
	if(sok) {
		for _, s := range summaryJSON["services"].([]interface{}) {
			theService := s.(map[string]interface{})
			_, servicePlanExist := theService["service_plan"]
			if(servicePlanExist) {
				boundedApps := theService["bound_app_count"].(float64)
				if(boundedApps > 0) {
					servicePlan := theService["service_plan"].(map[string]interface{})
					_, serviceExist := servicePlan["service"]
					if(serviceExist) {
						service := servicePlan["service"].(map[string]interface{})
						label := service["label"].(string)
						if (strings.Contains(label,"rabbit") || strings.Contains(label,"redis") || strings.Contains(label,"mysql")) {
							services = append(services,
								Service{
									Label:       label,
									ServicePlan: servicePlan["name"].(string),
								})
						}
					}
				}
			}
		}
	}
	return apps, services, nil
}
