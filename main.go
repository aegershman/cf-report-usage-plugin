package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-report-usage-plugin/presenters"
	"github.com/aegershman/cf-report-usage-plugin/report"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

// ReportUsageCmd -
type ReportUsageCmd struct{}

type orgNamesFlag struct {
	names []string
}

func (o *orgNamesFlag) String() string {
	return fmt.Sprint(o.names)
}

func (o *orgNamesFlag) Set(value string) error {
	o.names = append(o.names, value)
	return nil
}

// Run -
func (cmd *ReportUsageCmd) Run(cli plugin.CliConnection, args []string) {

	// TODO can't really imagine a situation where this would happen, but
	// I don't know I guess I'll just leave it for now
	if args[0] != "report-usage" {
		return
	}

	var (
		orgNamesFlag orgNamesFlag
		formatFlag   string
		logLevelFlag string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&orgNamesFlag, "o", "")
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

	reporter := report.NewReporter(cli, orgNamesFlag.names)
	summaryReport, err := reporter.GetSummaryReport()
	if err != nil {
		log.Fatalln(err)
	}

	presenter := presenters.NewPresenter(*summaryReport, formatFlag) // todo hacky pointer
	presenter.Render()

}

// GetMetadata -
func (cmd *ReportUsageCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-report-usage-plugin",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 12,
			Build: 1,
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

func main() {
	plugin.Start(new(ReportUsageCmd))
}
