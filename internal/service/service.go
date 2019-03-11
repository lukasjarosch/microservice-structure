package service

import (
	"context"

	api "github.com/lukasjarosch/microservice-structure/pkg"
	"github.com/sirupsen/logrus"
)

type exampleService struct{}

func NewExampleService() api.HelloServer {
	return &exampleService{}
}

func (e *exampleService) Greeting(ctx context.Context, request *api.GreetingRequest) (*api.GreetingResponse, error) {
	logrus.Infof("call to Greeting() with: %v", request)
	return &api.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}
