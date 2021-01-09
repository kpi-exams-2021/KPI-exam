package server

import (
	"context"
	"fmt"
	"net/http"
	"tree/cmd/server/handlers"
)

type HttpPortNumber int

type TreeServer struct {
	port HttpPortNumber
	server *http.Server
	rootHandler handlers.TreeHttpHandler
	opHandler handlers.OperationHandler
}

func NewServer(
	rootHandler handlers.TreeHttpHandler,
	opHandler handlers.OperationHandler,
	port HttpPortNumber,
	) *TreeServer {
	return &TreeServer{
		port:    port,
		rootHandler: rootHandler,
		opHandler: opHandler,
	}
}

func (s *TreeServer) Start() error {
	if s.rootHandler == nil || s.opHandler == nil {
		return fmt.Errorf("handler not provided")
	}
	if s.port <= 0 {
		return fmt.Errorf("port is not provided")
	}

	multiplexer := new(http.ServeMux)
	multiplexer.HandleFunc("/", s.rootHandler)
	multiplexer.HandleFunc("/op/", s.opHandler)
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: multiplexer,
	}

	return s.server.ListenAndServe()
}

func (s *TreeServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server is not running")
	}

	return s.server.Shutdown(context.Background())
}
