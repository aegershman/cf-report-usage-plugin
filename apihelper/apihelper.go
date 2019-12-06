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
