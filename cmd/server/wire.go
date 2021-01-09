//+build wireinject

package server

import (
	"github.com/google/wire"
	"tree/cmd/server/handlers"
)

func InitializeServer(port HttpPortNumber, filename string) (*TreeServer, error) {
	wire.Build(NewServer, handlers.TreeHandler, handlers.OperationHttpHandler)
	return &TreeServer{}, nil
}
