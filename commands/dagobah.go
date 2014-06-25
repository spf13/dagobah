// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

//var Config *hugolib.Config
var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: "Dagobah is an awesome planet style RSS aggregator",
	Long: `Dagobah provides planet style RSS aggregation. It
is inspired by python planet. It has a simple YAML configuration
and provides it's own webserver.`,
	Run: rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	initConfig()
	fmt.Println(viper.Get("feeds"))
	fmt.Println(viper.GetString("appname"))
}

func init() {
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
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

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
