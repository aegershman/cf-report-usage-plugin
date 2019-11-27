package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-usage-report-plugin/apihelper"
	"github.com/aegershman/cf-usage-report-plugin/models"
	"github.com/aegershman/cf-usage-report-plugin/presenters"
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
		Name: "cf-usage-report-plugin",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 9,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "usage-report",
				HelpText: "View AIs, SIs and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf usage-report [-o orgName...] --format formatChoice",
					Options: map[string]string{
						"o":         "organization(s) included in report. Flag can be provided multiple times.",
						"format":    "format to print as (options: string,table) (default: table)",
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

	var orgs []models.Org

	if len(userFlags.OrgNames) > 0 {
		for _, orgArg := range userFlags.OrgNames {
			org, err := cmd.getOrg(orgArg)
			if err != nil {
				log.Fatalln(err)
			}
			orgs = append(orgs, org)
		}
	} else {
		orgs, err = cmd.getOrgs()
		if err != nil {
			log.Fatalln(err)
		}
	}

	summaryReport := models.NewSummaryReport(orgs)

	presenter := presenters.NewPresenter(summaryReport, formatFlag)
	presenter.Render()
}

func (cmd *UsageReportCmd) getOrgs() ([]models.Org, error) {
	rawOrgs, err := cmd.apiHelper.GetOrgs()
	if nil != err {
		return nil, err
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

func (cmd *UsageReportCmd) getOrg(name string) (models.Org, error) {
	rawOrg, err := cmd.apiHelper.GetOrg(name)
	if nil != err {
		return models.Org{}, err
	}

	return cmd.getOrgDetails(rawOrg)
}

func (cmd *UsageReportCmd) getOrgDetails(o models.Org) (models.Org, error) {
	usage, err := cmd.apiHelper.GetOrgMemoryUsage(o)
	if nil != err {
		return models.Org{}, err
	}
	quota, err := cmd.apiHelper.GetQuotaMemoryLimit(o.QuotaURL)
	if nil != err {
		return models.Org{}, err
	}
	spaces, err := cmd.getSpaces(o.SpacesURL)
	if nil != err {
		return models.Org{}, err
	}

	return models.Org{
		Name:        o.Name,
		MemoryQuota: int(quota),
		MemoryUsage: int(usage),
		Spaces:      spaces,
	}, nil
}

func (cmd *UsageReportCmd) getSpaces(spaceURL string) ([]models.Space, error) {
	rawSpaces, err := cmd.apiHelper.GetOrgSpaces(spaceURL)
	if nil != err {
		return nil, err
	}
	var spaces = []models.Space{}
	for _, s := range rawSpaces {
		apps, services, err := cmd.getAppsAndServices(s.SummaryURL)
		if nil != err {
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
	if nil != err {
		return nil, nil, err
	}
	return apps, services, nil
}

// Run -
func (cmd *UsageReportCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "usage-report" {
		cmd.apiHelper = apihelper.New(cli)
		cmd.UsageReportCommand(args)
	}
}

func main() {
	plugin.Start(new(UsageReportCmd))
}
