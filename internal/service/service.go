package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure/internal/config"
	api "github.com/lukasjarosch/microservice-structure/pkg"
)

type exampleService struct {
	config *config.Config
}

func NewExampleService(config *config.Config) api.HelloServer {
	return &exampleService{
		config: config,
	}
}

func (e *exampleService) Greeting(ctx context.Context, request *api.GreetingRequest) (*api.GreetingResponse, error) {
	return &api.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *api.FarewellRequest) (*api.FarewellResponse, error) {
	return &api.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
