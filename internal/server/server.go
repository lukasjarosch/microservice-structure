package server

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	"github.com/lukasjarosch/microservice-structure/internal/service"
	godin "github.com/lukasjarosch/microservice-structure/pkg/grpc"
	"github.com/lukasjarosch/microservice-structure/pkg/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// exampleServiceServer is a wrapper to hold all of our services stuff.
// Anything related to the transport-layer can be set up here
type exampleServiceServer struct {
	GRPC        *godin.Server
	HTTPGateway *http.Server
}

// NewExampleServiceServer constructs a new Server using the exampleService as
// gRPC handler implementation.
func NewExampleServiceServer(config *config.Config, logger *logrus.Logger) *exampleServiceServer {

	// setup the business logic with it's dependencies
	svc := service.NewExampleService(config, logger)
	// wrap our business logic in the transport handler
	handler := NewExampleServiceHandler(svc)

	// attach our business logic to the gRPC server
	impl := func(g *grpc.Server) {
		greeter.RegisterHelloServer(g, handler)
	}

	// create the actual gRPC server
	// See pkg/server/server.go for the default options
	server := godin.NewServer(
		godin.Name("examle"),
		godin.Implementation(impl),

		// Override config with env variables configured by our business domain
		godin.GrpcNetworkPort(config.GrpcPort),
		godin.PrometheusNetworkPort(config.PrometheusPort),
		godin.PrometheusEndpoint(config.MetricsEndpoint),
	)

	// register HTTP gateway
	gatewayServer, err := http.GatewayServer(server.Options.ServerConfig.Network, greeter.RegisterHelloHandler)
	if err != nil {
		logger.Fatalf("failed to setup the HTTP gateway server: %v", err)
	}

	return &exampleServiceServer{
		GRPC:        server,
		HTTPGateway: gatewayServer,
	}
}

// exampleServiceHandler is the transport-layer wrapper of our business-logic in the server package
// Everything concerning requests/responses belongs in here. Only conversion (business-model <-> protobuf) should happen here actually.
type exampleServiceHandler struct {
	implementation *service.ExampleService
}

func NewExampleServiceHandler(implementation *service.ExampleService) *exampleServiceHandler {
	return &exampleServiceHandler{
		implementation: implementation,
	}
}

func (e *exampleServiceHandler) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	greeting, err := e.implementation.Greeting(request.Name)
	if err != nil {
		return nil, err
	}

	return &greeter.GreetingResponse{Greeting: greeting}, nil
}

func (e *exampleServiceHandler) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	farewell, err := e.implementation.Farewell(request.Name)
	if err != nil {
		return nil, err
	}

	return &greeter.FarewellResponse{Farewell: farewell}, nil
}
