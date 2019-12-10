package v2client

// OrgQuota -
// A space quota looks very similar but it uses a different (v2) API endpoint
// just to be safe, going to explicitly reference this as a way to get quota of an Org
type OrgQuota struct {
	AppInstanceLimit        int    `json:"app_instance_limit"`
	AppTaskLimit            int    `json:"app_task_limit"`
	GUID                    string `json:"guid"`
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

// GetOrgQuotaByOrgGUID -
func (o *OrgQuotasService) GetOrgQuotaByOrgGUID(orgGUID string) (OrgQuota, error) {
	org, err := o.client.cfc.GetOrgByGuid(orgGUID)
	if err != nil {
		return OrgQuota{}, err
	}

	quota, err := org.Quota()
	if err != nil {
		return OrgQuota{}, err
	}

	return OrgQuota{
		AppInstanceLimit:        quota.AppInstanceLimit,
		AppTaskLimit:            quota.AppTaskLimit,
		GUID:                    quota.Guid,
		InstanceMemoryLimit:     quota.InstanceMemoryLimit,
		MemoryLimit:             quota.MemoryLimit,
		Name:                    quota.Name,
		TotalPrivateDomains:     quota.TotalPrivateDomains,
		TotalReservedRoutePorts: quota.TotalReservedRoutePorts,
		TotalRoutes:             quota.TotalRoutes,
		TotalServiceKeys:        quota.TotalServiceKeys,
		TotalServices:           quota.TotalServices,
	}, nil
}
