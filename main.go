package main

import (
	"flag"
	"tree/cmd/client"
	"tree/cmd/server"
)

var clientMode = flag.Bool("c", false, "Enable client mode")

func main() {
	flag.Parse()
	if *clientMode {
		client.RunClient()
	} else {
		server.RunServer()
	}
}
