// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"strings"

	"labix.org/v2/mgo/bson"

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
	port := viper.GetString("port")

	r := gin.Default()
	templates := loadTemplates("home.html")
	r.HTMLTemplates = templates

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", homeRoute)
	r.GET("/static/*filepath", staticServe)
	//r.GET("/feed/:key", feedRoute)
	//r.GET("/post/:key", postRoute)
	r.Run(":" + port)
}

func staticServe(c *gin.Context) {
	static, err := rice.FindBox("static")
	if err != nil {
		log.Fatal(err)
	}
	original := c.Req.URL.Path
	c.Req.URL.Path = c.Params.ByName("filepath")
	fmt.Println(c.Params.ByName("filepath"))
	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Req)
	c.Req.URL.Path = original
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

	funcMap := template.FuncMap{
		"html":  ProperHtml,
		"title": func(a string) string { return strings.Title(a) },
	}

	templates.Funcs(funcMap)

	return templates
}

func ProperHtml(text string) template.HTML {
	if strings.Contains(text, "content:encoded>") || strings.Contains(text, "content/:encoded>") {
		text = html.UnescapeString(text)
	}
	return template.HTML(html.UnescapeString(template.HTMLEscapeString(text)))
}

func homeRoute(c *gin.Context) {
	var posts []Itm
	results := Items().Find(bson.M{}).Sort("-date").Limit(20)
	results.All(&posts)

	obj := gin.H{"title": "Go Rules", "posts": posts}
	c.HTML(200, "home.html", obj)
}
