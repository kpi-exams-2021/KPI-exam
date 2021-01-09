//+build wireinject

package client

import "github.com/google/wire"

func InitializeClient(api Api, filename string) (*TreeClient, error) {
	wire.Build(NewClient)
	return &TreeClient{}, nil
}
