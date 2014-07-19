![alt tag](https://raw.githubusercontent.com/spf13/dagobah/master/commands/static/images/logo.png)

Dagobah is an awesome RSS feed aggregator &amp; reader written in go inspired by planet


## Installing

Simply download the appropriate executable for your platform from the [releases page](https://github.com/spf13/dagobah/releases).

Dagobah depends on [MongoDB](http://mongodb.org). Please make sure to install that prior to running Dagobah.


## Adding Feeds

Dagobah is configured via a simple config file.
By default Dagobah expects this to be either at /etc/dagobah/config.yaml or ~/.dagobah/config.yaml
You can provide your own location with `--config`.

An example file is:

    title: "Go Feeds Powered by Dagobah"
    feeds:
        - "http://spf13.com/index.xml"
        - "http://dave.cheney.net/feed"
        - "http://www.goinggo.net/feeds/posts/default"
        - "http://blog.labix.org/feed"
        - "http://blog.golang.org/feed.atom"

## Running Dagobah

Dagobah will run a high performance web server and fetch and update the feeds in the background. Simply run dagobah to start it up.

./dagobah

## Building From Source

Dagobah depends on [go.rice](https://github.com/GeertJohan/go.rice) to embed the static files and templates
into the binary.

    go get github.com/spf13/dagobah.git
    cd $GOPATH/src/github.com/spf13/dagobah/commands
    rice embed-go
    cd ..
    go build


**If you want to work with the templates please make sure to `rice clean` so that rice will not load from the embedded files.**


The Dagobah logo is based on the Go mascot designed by Renée French and copyrighted under the Creative Commons Attribution 3.0 license.
