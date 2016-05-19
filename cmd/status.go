package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/composer22/docker-deploy-client/client"
	"github.com/composer22/docker-deploy-server/db"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve status of a previous deploy",
	Long: `Retrieve the status from a previous deploy to a docker-deploy-server
using the deploy ID returned from the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Repository name is mandatory.")
			os.Exit(0)
		}
		statusGetAndPrint(args[0])
	},
	Example: `docker-deploy-server status "DC8D9C2E-8161-4FC0-937F-4CA7037970D5"`,
}

func statusGetAndPrint(deployID string) {
	for {
		result, err := client.New(token, url).Status(deployID)
		if err != nil {
			client.PrintErr(err.Error())
			os.Exit(0)
		}
		printStatusResult(result)
		if pollInterval <= 0 {
			return
		}
		time.Sleep(time.Duration(pollInterval) * time.Second)
	}
}

func printStatusResult(result string) {
	if formatted == false {
		fmt.Println(result)
		return
	}

	var dr db.DeployStatus
	b := []byte(result)
	if err := json.Unmarshal(b, &dr); err != nil {
		client.PrintErr(err.Error())
		os.Exit(0)
	}
	var err error
	b, err = json.MarshalIndent(dr, "", "\t")
	if err != nil {
		client.PrintErr(err.Error())
		os.Exit(0)
	}
	fmt.Println(string(b))
}

func init() {
	RootCmd.AddCommand(statusCmd)
	statusCmd.SetUsageTemplate(statusUsageTemplate())
}

// Override help template.
func statusUsageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags] DEPLOY-ID"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

}
