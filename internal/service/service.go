package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure/internal/config"
	api "github.com/lukasjarosch/microservice-structure/internal"
	log "go.uber.org/zap"
	"time"
)

type exampleService struct {
	config *config.Config
	log *log.SugaredLogger
}

func NewExampleService(config *config.Config, log *log.SugaredLogger) api.HelloServer {
	return &exampleService{
		config: config,
		log: log,
	}
}

func (e *exampleService) Greeting(ctx context.Context, request *api.GreetingRequest) (*api.GreetingResponse, error) {
	time.Sleep(2 * time.Second)
	return &api.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *api.FarewellRequest) (*api.FarewellResponse, error) {
	return &api.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
