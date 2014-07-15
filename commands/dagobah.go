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
	fmt.Println(viper.Get("feeds"))
	fmt.Println(viper.GetString("appname"))
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetDefault("feeds", []string{"http://spf13.com/index.xml"})

	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
	RootCmd.PersistentFlags().StringP("dbname", "d", "dagobah", "name of the database")
	RootCmd.PersistentFlags().Int("dbport", 27017, "port to access mongoDB")
	RootCmd.PersistentFlags().String("dbhost", "localhost", "host where mongoDB is")
	RootCmd.PersistentFlags().String("dbusername", "", "username to connect to mongoDB with")
	RootCmd.PersistentFlags().String("dbpassword", "", "password to connect to mongoDB with")

	viper.BindPFlag("dbusername", RootCmd.PersistentFlags().Lookup("dbusername"))
	viper.BindPFlag("dbpassword", RootCmd.PersistentFlags().Lookup("dbpassword"))
	viper.BindPFlag("dbhost", RootCmd.PersistentFlags().Lookup("dbhost"))
	viper.BindPFlag("dbport", RootCmd.PersistentFlags().Lookup("dbport"))
	viper.BindPFlag("dbname", RootCmd.PersistentFlags().Lookup("dbname"))
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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
