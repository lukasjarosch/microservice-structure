package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	godin "github.com/lukasjarosch/microservice-structure/pkg/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type exampleService struct {
	greeter.HelloServer
	config *config.Config
	logger *logrus.Logger
}

// NewExampleService constructs a new Server using the exampleService as
// gRPC handler implementation
func NewExampleService(config *config.Config, logger *logrus.Logger) *godin.Server {

	// setup the business logic with it's dependencies
	handler := &exampleService{
		logger: logger,
		config: config,
	}

	// register handler as implementation
	impl := func(g *grpc.Server) {
		greeter.RegisterHelloServer(g, handler)
	}

	// create new server
	// See pkg/server/server.go for the default options
	server := godin.NewServer(
		godin.Name("examle"),
		godin.Implementation(impl),

		// Override config with env variables configured by our business domain
		godin.GrpcNetworkPort(config.GrpcPort),
		godin.PrometheusNetworkPort(config.PrometheusPort),
		godin.PrometheusEndpoint(config.MetricsEndpoint),
	)
	return server
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	return &greeter.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
