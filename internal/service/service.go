package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
)

type exampleService struct {
	greeter.HelloServer
	config *config.Config
	logger *logrus.Logger
}

// NewExampleService returns our business-implementation of the exampleService
func NewExampleService(config *config.Config, logger *logrus.Logger) *exampleService {

	service := &exampleService{
		logger: logger,
		config: config,
	}

	return service
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	e.logger.Info("ohai")
	return &greeter.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
