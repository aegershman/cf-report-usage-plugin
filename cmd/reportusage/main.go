package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-report-usage-plugin/internal/presentation"
	"github.com/aegershman/cf-report-usage-plugin/internal/report"
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

type formatFlag struct {
	formats []string
}

func (f *formatFlag) String() string {
	return fmt.Sprint(f.formats)
}

func (f *formatFlag) Set(value string) error {
	f.formats = append(f.formats, value)
	return nil
}

// reportUsageCommand is the "real" main entrypoint into program execution
func (cmd *reportUsageCmd) reportUsageCommand(cli plugin.CliConnection, args []string) {

	var (
		orgNamesFlag orgNamesFlag
		formatFlag   formatFlag
		logLevelFlag string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&orgNamesFlag, "o", "")
	flagss.Var(&formatFlag, "format", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")
	if err := flagss.Parse(args[1:]); err != nil {
		log.Fatalln(err)
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	reporter := report.NewClient(cli)
	summaryReport, err := reporter.GetSummaryReportByOrgNames(orgNamesFlag.names...)
	if err != nil {
		log.Fatalln(err)
	}

	presenter := presentation.NewPresenter(*summaryReport, formatFlag.formats...) // todo hacky pointer
	presenter.Render()
}

// Run -
func (cmd *reportUsageCmd) Run(cli plugin.CliConnection, args []string) {
	switch args[0] {
	case "report-usage":
		cmd.reportUsageCommand(cli, args)
	default:
		log.Debugln("did you know plugin commands can still get ran when uninstalling a plugin? interesting, right?")
		return
	}
}

// GetMetadata -
func (cmd *reportUsageCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-report-usage-plugin",
		Version: plugin.VersionType{
			Major: 3,
			Minor: 2,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "report-usage",
				HelpText: "View AIs, SIs and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf report-usage [-o orgName...] [--format formatChoice...]",
					Options: map[string]string{
						"o":         "organization(s) included in report. Flag can be provided multiple times.",
						"format":    "format to print as (options: string,table,json). Flag can be provided multiple times.",
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
