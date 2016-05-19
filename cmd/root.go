package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	token        string
	url          string
	pollInterval int
	formatted    bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "docker-deploy-client",
	Short: "Deploy and check status of deploying Docker containers",
	Long: `Client for making requests to a docker-deploy-server to deploy
 Docker containers into one or more machines (swarm cluster) and check the
status of a previous deploy request,`,
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docker-deploy-client.yaml)")
	RootCmd.PersistentFlags().StringP("token", "o", "", "API Token")
	RootCmd.PersistentFlags().StringP("url", "u", "", "docker-deploy-server endpoint")
	RootCmd.PersistentFlags().StringP("poll_interval", "i", "5", "Polling interval for status check")
	RootCmd.PersistentFlags().BoolP("formatted", "f", true, "JSON indented status results")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("poll_interval", RootCmd.PersistentFlags().Lookup("poll_interval"))
	viper.BindPFlag("formatted", RootCmd.PersistentFlags().Lookup("formatted"))
	viper.SetDefault("poll_interval", "0")
	viper.SetDefault("formatted", "true")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".docker-deploy-client") // name of config file (without extension).
	viper.AddConfigPath("$HOME")                 // adding home directory as first search path.
	viper.AddConfigPath(".")                     // adding current.
	viper.AutomaticEnv()                         // read in environment variables that match.

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Cannot find configuration file.\n\nERR: %s\n", err.Error())
		os.Exit(0)
	}
	token = viper.GetString("token")
	url = viper.GetString("url")
	pollInterval = viper.GetInt("poll_interval")
	formatted = viper.GetBool("formatted")
}
