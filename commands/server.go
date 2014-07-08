// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server for feeds",
	Long:  `Dagobah will serve all feeds listed in the config file.`,
	Run:   serverRun,
}

func init() {
	serverCmd.Flags().Int("port", 1138, "Port to run Dagobah server on")
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
}

func serverRun(cmd *cobra.Command, args []string) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	templates := loadTemplates("home.html")
	r.HTMLTemplates = templates

	r.GET("/", homeRoute)

	static, err := rice.FindBox("static")
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/static/*filepath", func(c *gin.Context) {
		original := c.Req.URL.Path
		c.Req.URL.Path = c.Params.ByName("filepath")
		fmt.Println(c.Params.ByName("filepath"))
		http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Req)
		c.Req.URL.Path = original
	})
	//r.GET("/feed/:key", feedRoute)
	//r.GET("/post/:key", postRoute)
	port := viper.GetString("port")
	r.Run(":" + port)
}

func loadTemplates(list ...string) *template.Template {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	templates := template.New("")

	for _, x := range list {
		templateString, err := templateBox.String(x)
		if err != nil {
			log.Fatal(err)
		}

		// get file contents as string
		_, err = templates.New(x).Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
	}

	return templates
}

func homeRoute(c *gin.Context) {
	obj := gin.H{"title": "Go Rules"}
	c.HTML(200, "home.html", obj)
}
