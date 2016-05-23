package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/composer22/docker-deploy-client/client"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a Docker image to one or more machines.",
	Long: `Given a Docker image and environment, send a request to a
docker-deploy-server to start containers in one or more machines (swarm cluster).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Repository name is mandatory.")
			os.Exit(0)
		}
		r := args[0]
		if cmd.Flag("env").Value.String() == "" {
			fmt.Println("Environment is mandatory.")
			os.Exit(0)
		}
		e := cmd.Flag("env").Value.String()
		t := cmd.Flag("image-tag").Value.String()
		p, err := strconv.ParseBool(cmd.Flag("poll-status").Value.String())
		if err != nil {
			client.PrintErr(err.Error())
			os.Exit(0)
		}

		result, err := client.New(token, url).Deploy(r, t, e)
		if err != nil {
			client.PrintErr(err.Error())
			os.Exit(0)
		}
		dr := printDeployResult(result)
		if p {
			statusGetAndPrint(dr.DeployID)
		}
	},
	Example: `docker-deploy-server deploy hello-world -e dev
docker-deploy-server deploy hello-world -e dev --poll-status=false
docker-deploy-server deploy hello-world -e dev -t 1.0.0-131
docker-deploy-server deploy -e dev hello-world
docker-deploy-server deploy -e dev --poll-status=false hello-world
docker-deploy-server deploy -e dev -t 1.0.0-131 hello-world
`,
}

type DeployResult struct {
	DeployID string `json:"deployID"`
}

func printDeployResult(result string) *DeployResult {
	var dr DeployResult
	b := []byte(result)
	if err := json.Unmarshal(b, &dr); err != nil {
		client.PrintErr(err.Error())
		os.Exit(0)
	}

	if formatted == false {
		fmt.Println(result)
		return &dr
	}
	var err error
	b, err = json.MarshalIndent(dr, "", "\t")
	if err != nil {
		client.PrintErr(err.Error())
		os.Exit(0)
	}
	fmt.Println(string(b))
	return &dr
}

func init() {
	RootCmd.AddCommand(deployCmd)
	deployCmd.SetUsageTemplate(deployUsageTemplate())
	deployCmd.Flags().StringP("env", "e", "", "Targeted environment to deploy (ex dev, qa etc.)")
	deployCmd.Flags().StringP("image-tag", "t", "latest", "Docker image tag")
	deployCmd.Flags().BoolP("poll-status", "p", true, "Immediatly poll the status after the deploy")

}

// Override help template.
func deployUsageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags] REPO"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
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
