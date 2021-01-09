package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func RunServer() {
	server, err := InitializeServer(8000, "server.json")
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return
	}

	go func() {
		fmt.Println("Starting tree server...")
		if err = server.Start(); err == http.ErrServerClosed {
			_ = fmt.Errorf(err.Error())
			return
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<- sigChannel

	if err = server.Stop(); err != nil && err != http.ErrServerClosed {
		_ = fmt.Errorf(err.Error())
	}
}
