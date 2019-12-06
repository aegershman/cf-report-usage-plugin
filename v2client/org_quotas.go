package v2client

import (
	"github.com/aegershman/cf-report-usage-plugin/models"
)

// OrgQuotasService -
type OrgQuotasService service

// GetOrgQuota returns an org's quota. A space quota looks very similar
// but it uses a different (v2) API endpoint, so just to be safe, going to explicitly
// reference this as a way to get quota of an Org
func (o *OrgQuotasService) GetOrgQuota(quotaURL string) (models.OrgQuota, error) {
	quotaJSON, err := o.client.Curl(quotaURL)
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
