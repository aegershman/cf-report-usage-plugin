package report

import (
	"github.com/aegershman/cf-report-usage-plugin/v2client"
	"github.com/cloudfoundry/cli/plugin"
)

// Reporter -
type Reporter struct {
	orgNames []string
	client   *v2client.Client
}

// NewReporter -
func NewReporter(cli plugin.CliConnection, orgNames []string) *Reporter {
	client := v2client.NewClient(cli)

	r := &Reporter{
		client:   client,
		orgNames: orgNames,
	}

	return r
}

// GetSummaryReport -
func (r *Reporter) GetSummaryReport() (*SummaryReport, error) {
	populatedOrgs, err := r.getOrgs()
	if err != nil {
		return &SummaryReport{}, nil
	}

	return NewSummaryReport(populatedOrgs), nil
}

func (r *Reporter) getOrgs() ([]v2client.Org, error) {
	var rawOrgs []v2client.Org

	if len(r.orgNames) > 0 {
		for _, orgName := range r.orgNames {
			rawOrg, err := r.client.Orgs.GetOrg(orgName)
			if err != nil {
				return nil, err
			}
			rawOrgs = append(rawOrgs, rawOrg)
		}
	} else {
		extraRawOrgs, err := r.client.Orgs.GetOrgs()
		if err != nil {
			return nil, err
		}
		rawOrgs = extraRawOrgs
	}

	var orgs = []v2client.Org{}

	for _, o := range rawOrgs {
		orgDetails, err := r.getOrgDetails(o)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, orgDetails)
	}
	return orgs, nil
}

func (r *Reporter) getOrgDetails(o v2client.Org) (v2client.Org, error) {
	usage, err := r.client.Orgs.GetOrgMemoryUsage(o)
	if err != nil {
		return v2client.Org{}, err
	}

	// TODO teeing up to swap out for 'quota' being it's own managed entity
	// for time being, going to simply modify it _here_ to not break anything obvious
	quota, err := r.client.OrgQuotas.GetOrgQuota(o.QuotaURL)
	if err != nil {
		return v2client.Org{}, err
	}
	spaces, err := r.getSpaces(o.SpacesURL)
	if err != nil {
		return v2client.Org{}, err
	}

	return v2client.Org{
		Name:        o.Name,
		MemoryQuota: quota.MemoryLimit,
		MemoryUsage: int(usage),
		Spaces:      spaces,
	}, nil
}

func (r *Reporter) getSpaces(spaceURL string) ([]v2client.Space, error) {
	rawSpaces, err := r.client.Orgs.GetOrgSpaces(spaceURL)
	if err != nil {
		return nil, err
	}
	var spaces = []v2client.Space{}
	for _, s := range rawSpaces {
		apps, services, err := r.getAppsAndServices(s.SummaryURL)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces,
			v2client.Space{
				Name:     s.Name,
				Apps:     apps,
				Services: services,
			},
		)
	}
	return spaces, nil
}

func (r *Reporter) getAppsAndServices(summaryURL string) ([]v2client.App, []v2client.Service, error) {
	apps, services, err := r.client.Spaces.GetSpaceAppsAndServices(summaryURL)
	if err != nil {
		return nil, nil, err
	}
	return apps, services, nil
}
