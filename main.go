package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-report-usage-plugin/presenters"
	"github.com/aegershman/cf-report-usage-plugin/report"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

type reportUsageCmd struct{}

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

var (
	globalOrgNamesFlag orgNamesFlag
	globalFormatFlag   string
	globalLogLevelFlag string
)

func (cmd *reportUsageCmd) parseFlags(args []string) error {
	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&globalOrgNamesFlag, "o", "")
	flagss.StringVar(&globalFormatFlag, "format", "table", "")
	flagss.StringVar(&globalLogLevelFlag, "log-level", "info", "")

	if err := flagss.Parse(args[1:]); err != nil {
		return err
	}

	logLevel, err := log.ParseLevel(globalLogLevelFlag)
	if err != nil {
		return err
	}

	log.SetLevel(logLevel)

	return nil
}

// reportUsageCommand is the "real" main entrypoint into program execution
func (cmd *reportUsageCmd) reportUsageCommand(cli plugin.CliConnection, args []string) {
	if err := cmd.parseFlags(args); err != nil {
		log.Fatalln(err)
	}

	reporter := report.NewReporter(cli, globalOrgNamesFlag.names)
	summaryReport, err := reporter.GetSummaryReport()
	if err != nil {
		log.Fatalln(err)
	}

	presenter := presenters.NewPresenter(*summaryReport, globalFormatFlag) // todo hacky pointer
	presenter.Render()
}

// Run -
func (cmd *reportUsageCmd) Run(cli plugin.CliConnection, args []string) {
	switch args[0] {
	case "report-usage":
		cmd.reportUsageCommand(cli, args)
	default:
		log.Infoln("did you know plugin commands can still get ran when uninstalling a plugin? interesting, right?")
		return
	}
}

// GetMetadata -
func (cmd *reportUsageCmd) GetMetadata() plugin.PluginMetadata {
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
	plugin.Start(new(reportUsageCmd))
}
