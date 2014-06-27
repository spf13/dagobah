// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"os"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch feeds",
	Long:  `Dagobah will fetch all feeds listed in the config file.`,
	Run:   fetchRun,
}

type ConfigFeed struct {
	Name, Url string
}

type Config struct {
	Feeds []ConfigFeed
	Port  int
}

func init() {
	fetchCmd.Flags().Int("rsstimeout", 5, "Timeout (in min) for RSS retrival")
	viper.BindPFlag("rsstimeout", fetchCmd.Flags().Lookup("rsstimeout"))
}

func fetchRun(cmd *cobra.Command, args []string) {
	Fetcher()
}

func Fetcher() {
	var config Config
	err := viper.GetAllIntoStruct(&config)
	if err != nil {
		fmt.Println(err)
	}

	for _, feed := range config.Feeds {
		go PollFeed(feed.Url)
	}
}

func PollFeed(uri string) {
	timeout := viper.GetInt("RSSTimeout")
	if timeout < 1 {
		timeout = 1
	}
	feed := rss.New(timeout, true, chanHandler, itemHandler)

	for {
		if err := feed.Fetch(uri, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
			return
		}

		fmt.Printf("Sleeping for %d seconds on %s\n", feed.SecondsTillUpdate(), uri)
		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}
