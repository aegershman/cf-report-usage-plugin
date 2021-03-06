package v2client

import "fmt"

// Space -
type Space struct {
	Apps       []App
	GUID       string
	Name       string
	Services   []Service
	SummaryURL string
}

// SpacesService -
type SpacesService service

// GetSpaceAppsAndServicesBySpaceGUID returns the apps and the services from a space
func (s *SpacesService) GetSpaceAppsAndServicesBySpaceGUID(spaceGUID string) ([]App, []Service, error) {
	path := fmt.Sprintf("/v2/spaces/%s/summary", spaceGUID)
	summaryJSON, err := s.client.Curl(path)
	if err != nil {
		return nil, nil, err
	}

	apps := []App{}
	services := []Service{}

	if _, ok := summaryJSON["apps"]; ok {
		for _, a := range summaryJSON["apps"].([]interface{}) {
			theApp := a.(map[string]interface{})
			apps = append(apps,
				App{
					GUID:             theApp["guid"].(string),
					Instances:        int(theApp["instances"].(float64)),
					Memory:           int(theApp["memory"].(float64)),
					Name:             theApp["name"].(string),
					RunningInstances: int(theApp["running_instances"].(float64)),
				})
		}
	}

	if _, ok := summaryJSON["services"]; ok {
		for _, s := range summaryJSON["services"].([]interface{}) {
			theService := s.(map[string]interface{})

			// these properties should exist whether 'service_plan' exists
			// should imply it's a user-provided service
			serviceToAppend := Service{
				GUID: theService["guid"].(string),
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
