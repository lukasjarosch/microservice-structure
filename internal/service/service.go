package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	"errors"
)

type exampleService struct {
	greeter.HelloServer
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("the given name is empty")
)

// NewExampleService returns our business-implementation of the exampleService
func NewExampleService(config *config.Config, logger *logrus.Logger) *exampleService {

	service := &exampleService{
		logger: logger,
		config: config,
	}

	return service
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {

	if request.Name == "" {
		return nil, ErrEmptyName
	}

	return &greeter.GreetingResponse{Greeting: "Hey there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
