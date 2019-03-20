package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	log "go.uber.org/zap"
)

type exampleService struct {
	config *config.Config
	log    *log.SugaredLogger
}

func NewExampleService(config *config.Config, log *log.SugaredLogger) greeter.HelloServer {
	return &exampleService{
		config: config,
		log:    log,
	}
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	return &greeter.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
