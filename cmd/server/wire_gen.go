// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"tree/cmd/server/handlers"
)

// Injectors from wire.go:

func InitializeServer(port HttpPortNumber, filename string) (*TreeServer, error) {
	treeHttpHandler := handlers.TreeHandler(filename)
	operationHandler := handlers.OperationHttpHandler(filename)
	treeServer := NewServer(treeHttpHandler, operationHandler, port)
	return treeServer, nil
}
