package service

import (
	"context"

	api "github.com/lukasjarosch/microservice-structure/internal"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	log "go.uber.org/zap"
)

type exampleService struct {
	config *config.Config
	log    *log.SugaredLogger
}

func NewExampleService(config *config.Config, log *log.SugaredLogger) api.HelloServer {
	return &exampleService{
		config: config,
		log:    log,
	}
}

func (e *exampleService) Greeting(ctx context.Context, request *api.GreetingRequest) (*api.GreetingResponse, error) {
	return &api.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *api.FarewellRequest) (*api.FarewellResponse, error) {
	return &api.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
