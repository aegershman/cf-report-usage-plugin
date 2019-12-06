package apihelper

import (
	"errors"
	"net/url"

	"github.com/aegershman/cf-report-usage-plugin/models"

	"github.com/aegershman/cf-report-usage-plugin/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

var (
	// ErrOrgNotFound -
	ErrOrgNotFound = errors.New("organization not found")
)

// CFAPIHelper wraps cf-curl results, acts as a cf-curl client
type CFAPIHelper interface {
	GetTarget() string
	GetOrgQuota(string) (models.OrgQuota, error)
	GetOrgMemoryUsage(models.Org) (float64, error)
	GetOrgSpaces(string) ([]models.Space, error)
	GetSpaceAppsAndServices(string) ([]models.App, []models.Service, error)
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
	if err != nil {
		return ""
	}
	apiep, _ := envInfo["routing_endpoint"].(string)
	u, err := url.Parse(apiep)
	if err != nil {
		log.Fatalln(err)
	}
	host := u.Host
	return host
}

// GetOrgQuota returns an org's quota. A space quota looks very similar
// but it uses a different (v2) API endpoint, so just to be safe, going to explicitly
// reference this as a way to get quota of an Org
func (api *APIHelper) GetOrgQuota(quotaURL string) (models.OrgQuota, error) {
	quotaJSON, err := cfcurl.Curl(api.cli, quotaURL)
	if err != nil {
		return models.OrgQuota{}, err
	}

	quota := quotaJSON["entity"].(map[string]interface{})
	return models.OrgQuota{
		Name:                    quota["name"].(string),
		TotalServices:           int(quota["total_services"].(float64)),
		TotalRoutes:             int(quota["total_routes"].(float64)),
		TotalPrivateDomains:     int(quota["total_private_domains"].(float64)),
		MemoryLimit:             int(quota["memory_limit"].(float64)),
		InstanceMemoryLimit:     int(quota["instance_memory_limit"].(float64)),
		AppInstanceLimit:        int(quota["app_instance_limit"].(float64)),
		AppTaskLimit:            int(quota["app_task_limit"].(float64)),
		TotalServiceKeys:        int(quota["total_service_keys"].(float64)),
		TotalReservedRoutePorts: int(quota["total_service_keys"].(float64)),
	}, nil
}

// GetOrgMemoryUsage returns amount of memory (in MB) a given org is currently using
func (api *APIHelper) GetOrgMemoryUsage(org models.Org) (float64, error) {
	usageJSON, err := cfcurl.Curl(api.cli, org.URL+"/memory_usage")
	if err != nil {
		return 0, err
	}
	return usageJSON["memory_usage_in_mb"].(float64), nil
}

// GetOrgSpaces returns the spaces in an org
func (api *APIHelper) GetOrgSpaces(spacesURL string) ([]models.Space, error) {
	nextURL := spacesURL
	spaces := []models.Space{}
	for nextURL != "" {
		spacesJSON, err := cfcurl.Curl(api.cli, nextURL)
		if err != nil {
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

// GetSpaceAppsAndServices returns the apps and the services from a space's /summary endpoint
func (api *APIHelper) GetSpaceAppsAndServices(summaryURL string) ([]models.App, []models.Service, error) {
	summaryJSON, err := cfcurl.Curl(api.cli, summaryURL)
	if err != nil {
		return nil, nil, err
	}

	apps := []models.App{}
	services := []models.Service{}

	if _, ok := summaryJSON["apps"]; ok {
		for _, a := range summaryJSON["apps"].([]interface{}) {
			theApp := a.(map[string]interface{})
			apps = append(apps,
				models.App{
					Name:             theApp["name"].(string),
					RunningInstances: int(theApp["running_instances"].(float64)),
					Instances:        int(theApp["instances"].(float64)),
					Memory:           int(theApp["memory"].(float64)),
				})
		}
	}

	if _, ok := summaryJSON["services"]; ok {
		for _, s := range summaryJSON["services"].([]interface{}) {
			theService := s.(map[string]interface{})

			// these properties should exist whether 'service_plan' exists
			// should imply it's a user-provided service
			serviceToAppend := models.Service{
				Name: theService["name"].(string),
				Type: theService["type"].(string),
			}

			// believe this only occurs with "type: managed_service_instance"
			if _, serviceBrokerNameExists := theService["service_broker_name"]; serviceBrokerNameExists {
				serviceBrokerName := theService["service_broker_name"].(string)
				serviceToAppend.ServiceBrokerName = serviceBrokerName
			}

			// believe this only occurs with "type: managed_service_instance"
			if _, servicePlanExists := theService["service_plan"]; servicePlanExists {
				servicePlan := theService["service_plan"].(map[string]interface{})
				if _, serviceExists := servicePlan["service"]; serviceExists {
					service := servicePlan["service"].(map[string]interface{})
					label := service["label"].(string)
					servicePlanName := servicePlan["name"].(string)

					serviceToAppend.ServicePlanLabel = label
					serviceToAppend.ServicePlanName = servicePlanName
				}
			}

			services = append(services, serviceToAppend)

		}
	}

	return apps, services, nil

}
