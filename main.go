package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-report-usage-plugin/apihelper"
	"github.com/aegershman/cf-report-usage-plugin/models"
	"github.com/aegershman/cf-report-usage-plugin/presenters"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

// UsageReportCmd -
type UsageReportCmd struct {
	apiHelper apihelper.CFAPIHelper
}

type flags struct {
	OrgNames []string
}

func (f *flags) String() string {
	return fmt.Sprint(f.OrgNames)
}

func (f *flags) Set(value string) error {
	f.OrgNames = append(f.OrgNames, value)
	return nil
}

// GetMetadata -
func (cmd *UsageReportCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-report-usage-plugin",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 11,
			Build: 2,
		},
		Commands: []plugin.Command{
			{
				Name:     "report-usage",
				HelpText: "View AIs, SIs and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf report-usage [-o orgName...] --format formatChoice",
					Options: map[string]string{
						"o":         "organization(s) included in report. Flag can be provided multiple times.",
						"format":    "format to print as (options: string,table,json) (default: table)",
						"log-level": "(options: info,debug,trace) (default: info)",
					},
				},
			},
		},
	}
}

// UsageReportCommand -
func (cmd *UsageReportCmd) UsageReportCommand(args []string) {
	var (
		userFlags    flags
		formatFlag   string
		logLevelFlag string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&userFlags, "o", "")
	flagss.StringVar(&formatFlag, "format", "table", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	orgs, err := cmd.getOrgs(userFlags.OrgNames)
	if err != nil {
		log.Fatalln(err)
	}

	summaryReport := models.NewSummaryReport(orgs)
	presenter := presenters.NewPresenter(*summaryReport, formatFlag) // todo hacky pointer
	presenter.Render()
}

func (cmd *UsageReportCmd) getOrgs(orgNames []string) ([]models.Org, error) {
	var rawOrgs []models.Org

	if len(orgNames) > 0 {
		for _, orgName := range orgNames {
			rawOrg, err := cmd.apiHelper.GetOrg(orgName)
			if err != nil {
				return nil, err
			}
			rawOrgs = append(rawOrgs, rawOrg)
		}
	} else {
		extraRawOrgs, err := cmd.apiHelper.GetOrgs()
		if err != nil {
			return nil, err
		}
		rawOrgs = extraRawOrgs
	}

	var orgs = []models.Org{}

	for _, o := range rawOrgs {
		orgDetails, err := cmd.getOrgDetails(o)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, orgDetails)
	}
	return orgs, nil
}

func (cmd *UsageReportCmd) getOrgDetails(o models.Org) (models.Org, error) {
	usage, err := cmd.apiHelper.GetOrgMemoryUsage(o)
	if err != nil {
		return models.Org{}, err
	}

	// TODO teeing up to swap out for 'quota' being it's own managed entity
	// for time being, going to simply modify it _here_ to not break anything obvious
	quota, err := cmd.apiHelper.GetOrgQuota(o.QuotaURL)
	if err != nil {
		return models.Org{}, err
	}
	spaces, err := cmd.getSpaces(o.SpacesURL)
	if err != nil {
		return models.Org{}, err
	}

	return models.Org{
		Name:        o.Name,
		MemoryQuota: quota.MemoryLimit,
		MemoryUsage: int(usage),
		Spaces:      spaces,
	}, nil
}

func (cmd *UsageReportCmd) getSpaces(spaceURL string) ([]models.Space, error) {
	rawSpaces, err := cmd.apiHelper.GetOrgSpaces(spaceURL)
	if err != nil {
		return nil, err
	}
	var spaces = []models.Space{}
	for _, s := range rawSpaces {
		apps, services, err := cmd.getAppsAndServices(s.SummaryURL)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces,
			models.Space{
				Name:     s.Name,
				Apps:     apps,
				Services: services,
			},
		)
	}
	return spaces, nil
}

func (cmd *UsageReportCmd) getAppsAndServices(summaryURL string) ([]models.App, []models.Service, error) {
	apps, services, err := cmd.apiHelper.GetSpaceAppsAndServices(summaryURL)
	if err != nil {
		return nil, nil, err
	}
	return apps, services, nil
}

// Run -
func (cmd *UsageReportCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "report-usage" {
		cmd.apiHelper = apihelper.New(cli)
		cmd.UsageReportCommand(args)
	}
}

func main() {
	plugin.Start(new(UsageReportCmd))
}
