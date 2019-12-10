package v2client

// OrgQuota -
// A space quota looks very similar but it uses a different (v2) API endpoint
// just to be safe, going to explicitly reference this as a way to get quota of an Org
type OrgQuota struct {
	AppInstanceLimit        int `json:"app_instance_limit"`
	AppTaskLimit            int `json:"app_task_limit"`
	GUID                    string
	InstanceMemoryLimit     int    `json:"instance_memory_limit"`
	MemoryLimit             int    `json:"memory_limit"`
	Name                    string `json:"name"`
	TotalPrivateDomains     int    `json:"total_private_domains"`
	TotalReservedRoutePorts int    `json:"total_reserved_route_ports"`
	TotalRoutes             int    `json:"total_routes"`
	TotalServiceKeys        int    `json:"total_service_keys"`
	TotalServices           int    `json:"total_services"`
}

// OrgQuotasService -
type OrgQuotasService service

// GetOrgQuota returns an org's quota. A space quota looks very similar
// but it uses a different (v2) API endpoint, so just to be safe, going to explicitly
// reference this as a way to get quota of an Org
func (o *OrgQuotasService) GetOrgQuota(quotaURL string) (OrgQuota, error) {
	quotaJSON, err := o.client.Curl(quotaURL)
	if err != nil {
		return OrgQuota{}, err
	}

	metadata := quotaJSON["metadata"].(map[string]interface{})
	guid := metadata["guid"].(string)

	quota := quotaJSON["entity"].(map[string]interface{})
	return OrgQuota{
		AppInstanceLimit:        int(quota["app_instance_limit"].(float64)),
		AppTaskLimit:            int(quota["app_task_limit"].(float64)),
		GUID:                    guid,
		InstanceMemoryLimit:     int(quota["instance_memory_limit"].(float64)),
		MemoryLimit:             int(quota["memory_limit"].(float64)),
		Name:                    quota["name"].(string),
		TotalPrivateDomains:     int(quota["total_private_domains"].(float64)),
		TotalReservedRoutePorts: int(quota["total_service_keys"].(float64)),
		TotalRoutes:             int(quota["total_routes"].(float64)),
		TotalServiceKeys:        int(quota["total_service_keys"].(float64)),
		TotalServices:           int(quota["total_services"].(float64)),
	}, nil
}
