package main

import (
	"flag"
	"github.com/kentaro/evergreen"
)

var port = flag.String("port", "9662", "Port to be listened to")
var cookieKey = flag.String("cookie-key", "evergreen", "Name for cookie storage")

func main() {
	flag.Parse()

	options := map[string]string{
		"port":       *port,
		"cookie_key": *cookieKey,
	}

	server := &evergreen.Server{}
	server.Run(options)
}
