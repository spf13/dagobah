// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: "Dagobah is an awesome planet style RSS aggregator",
	Long: `Dagobah provides planet style RSS aggregation. It
is inspired by python planet. It has a simple YAML configuration
and provides it's own webserver.`,
	Run: rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	go Server()
	go Fetcher()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetDefault("title", "Feeds powered by Dagobah")
	viper.SetDefault("feeds", []string{"http://spf13.com/index.xml"})

	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")

	RootCmd.PersistentFlags().String("mongodb_uri", "mongodb://localhost:27017/", "Uri to connect to mongoDB")
	viper.BindPFlag("mongodb_uri", RootCmd.PersistentFlags().Lookup("mongodb_uri"))
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.AddConfigPath("/etc/dagobah/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.dagobah/") // call multiple times to add many search paths
	viper.ReadInConfig()
}

func addCommands() {
	RootCmd.AddCommand(fetchCmd)
	RootCmd.AddCommand(serverCmd)
}

func Execute() {
	addCommands()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
