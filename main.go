package main

import (
	"flag"
	"log"
	"tree/ntrees"
)

var clientMode = flag.Bool("c", false, "Enable client mode")

func main() {
	//flag.Parse()
	//if *clientMode {
	//	client.RunClient()
	//} else {
	//	server.RunServer()
	//}

	ntree, err := ntrees.FromFile("ntree.json")
	if err != nil {
		log.Fatal(err)
	}
	ntree.Op()
	if err := ntrees.ToFile(ntree, "ntree2.json"); err != nil {
		log.Fatal(err)
	}
}
