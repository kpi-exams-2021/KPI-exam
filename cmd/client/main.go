package client

import "log"

func RunClient() {
	client, err := InitializeClient("http://localhost:8000/", "client.json")
	if err != nil {
		log.Fatal(err)
	}

	client.MakeOpPostTree()
}
