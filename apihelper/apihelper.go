package apihelper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/aegershman/cf-trueup-plugin/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
)

var (
	// ErrOrgNotFound -
	ErrOrgNotFound = errors.New("organization not found")
)

// Organization -
type Organization struct {
	URL       string
	Name      string
	QuotaURL  string
	SpacesURL string
}

// Space -
type Space struct {
	Name       string
	SummaryURL string
}

// App -
type App struct {
	Actual float64
	Desire float64
	RAM    float64
}

// Service -
type Service struct {
	Label       string
	ServicePlan string
}

// Orgs -
type Orgs []Organization

// Spaces -
type Spaces []Space

// Apps -
type Apps []App

// Services -
type Services []Service

// CFAPIHelper wraps cf curl results
type CFAPIHelper interface {
	GetTarget() string
	GetOrgs() (Orgs, error)
	GetOrg(string) (Organization, error)
	GetQuotaMemoryLimit(string) (float64, error)
	GetOrgMemoryUsage(Organization) (float64, error)
	GetOrgSpaces(string) (Spaces, error)
	GetSpaceAppsAndServices(string) (Apps, Services, error)
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
func (api *APIHelper) GetOrgs() (Orgs, error) {
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
			if name == "system" {
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

// GetOrg -
func (api *APIHelper) GetOrg(name string) (Organization, error) {
	query := fmt.Sprintf("name:%s", name)
	path := fmt.Sprintf("/v2/organizations?q=%s", url.QueryEscape(query))
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

// GetQuotaMemoryLimit returns memory quota (in MB) of a given org
func (api *APIHelper) GetQuotaMemoryLimit(quotaURL string) (float64, error) {
	quotaJSON, err := cfcurl.Curl(api.cli, quotaURL)
	if nil != err {
		return 0, err
	}
	return quotaJSON["entity"].(map[string]interface{})["memory_limit"].(float64), nil
}

// GetOrgMemoryUsage returns amount of memory (in MB) a given org is currently using
func (api *APIHelper) GetOrgMemoryUsage(org Organization) (float64, error) {
	usageJSON, err := cfcurl.Curl(api.cli, org.URL+"/memory_usage")
	if nil != err {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}

// GetOrgSpaces returns the spaces in an org
func (api *APIHelper) GetOrgSpaces(spacesURL string) (Spaces, error) {
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
func (api *APIHelper) GetSpaceAppsAndServices(summaryURL string) (Apps, Services, error) {
	apps := []App{}
	services := []Service{}
	summaryJSON, err := cfcurl.Curl(api.cli, summaryURL)
	if nil != err {
		return nil, nil, err
	}
	if _, ok := summaryJSON["apps"]; ok {
		for _, a := range summaryJSON["apps"].([]interface{}) {
			theApp := a.(map[string]interface{})
			apps = append(apps,
				App{
					Actual: theApp["running_instances"].(float64),
					Desire: theApp["instances"].(float64),
					RAM:    theApp["memory"].(float64),
				})
		}
	}
	if _, ok := summaryJSON["services"]; ok {
		for _, s := range summaryJSON["services"].([]interface{}) {
			theService := s.(map[string]interface{})
			if _, servicePlanExist := theService["service_plan"]; servicePlanExist {
				servicePlan := theService["service_plan"].(map[string]interface{})
				if _, serviceExist := servicePlan["service"]; serviceExist {
					service := servicePlan["service"].(map[string]interface{})
					label := service["label"].(string)
					if strings.EqualFold(label, "p-dataflow") {
						services = append(services,
							Service{
								Label:       "p-dataflow-servers",
								ServicePlan: servicePlan["name"].(string),
							})
					}
					if strings.EqualFold(label, "p-dataflow-analytics") {
						services = append(services,
							Service{
								Label:       "p-redis",
								ServicePlan: "p-dataflow-analytics",
							})
					} else if strings.EqualFold(label, "p-dataflow-relational") {
						services = append(services,
							Service{
								Label:       "p-mysql",
								ServicePlan: "p-dataflow-relational",
							})
					} else if strings.EqualFold(label, "p-dataflow-messaging") {
						services = append(services,
							Service{
								Label:       "p-rabbit",
								ServicePlan: "p-dataflow-messaging",
							})
					} else if strings.Contains(label, "p-circuit-breaker") || strings.Contains(label, "p-config-server") || strings.Contains(label, "p-service-registry") {
						services = append(services,
							Service{
								Label:       "p-spring-cloud-services",
								ServicePlan: label,
							})
					} else if strings.Contains(label, "rabbit") || strings.Contains(label, "redis") || strings.Contains(label, "mysql") {
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
	return apps, services, nil
}
