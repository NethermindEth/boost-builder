package main

import (
	"io"
	"strings"
	"text/tabwriter"
	"text/template"

	"gopkg.in/urfave/cli.v1"
)

var (
	AppHelpTemplate = `NAME:
   {{.App.Name}} - {{.App.Usage}}
USAGE:
   {{.App.HelpName}} [options] [-help]
   {{if .App.Version}}
VERSION:
   {{.App.Version}}
   {{end}}{{if .FlagGroups}}
{{range .FlagGroups}}{{.Name}} OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
{{end}}{{end}}
`
	AppHelpFlagGroups = []FlagGroup{
		{
			Name: "BUILDER API",
			Flags: []cli.Flag{
				EnableValidatorChecks,
				SecretKey,
				ListenAddr,
				GenesisForkVersion,
				BellatrixForkVersion,
				GenesisValidatorsRoot,
				BeaconEndpoint,
			},
		},
	}
)

// HelpData is a one shot struct to pass to the usage template
type HelpData struct {
	App        interface{}
	FlagGroups []FlagGroup
}

// FlagGroup is a collection of flags belonging to a single topic.
type FlagGroup struct {
	Name  string
	Flags []cli.Flag
}

func init() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.HelpPrinter = func(w io.Writer, tmpl string, data interface{}) {
		printHelp(w, tmpl, HelpData{App: data, FlagGroups: AppHelpFlagGroups})
	}
}

func printHelp(out io.Writer, templ string, data interface{}) {
	funcMap := template.FuncMap{"join": strings.Join}
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	w := tabwriter.NewWriter(out, 38, 8, 2, ' ', 0)
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
	w.Flush()
}
