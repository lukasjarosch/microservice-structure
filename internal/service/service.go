package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	godin "github.com/lukasjarosch/microservice-structure/pkg/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type exampleService struct {
	greeter.HelloServer
	config *config.Config
	log    *logrus.Entry
}

func NewExampleService(config *config.Config, log *logrus.Entry) *godin.Server {

	handler := &exampleService{}

	impl := func(g *grpc.Server) {
		greeter.RegisterHelloServer(g, handler)
	}

	// See pkg/server/server.go for the default options
	return godin.NewServer(
		godin.Name("examle"),
		godin.Implementation(impl),
	)
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	return &greeter.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
