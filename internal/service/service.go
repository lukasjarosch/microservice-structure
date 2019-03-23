package service

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	godin "github.com/lukasjarosch/microservice-structure/pkg/server"
	"github.com/prometheus/client_golang/prometheus"
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

	// Deliberately set as many options as possible, just for showing off.
	// Take a look at pkg/server/server.go for the default values
	return godin.NewServer(
		godin.Name("examle"),
		godin.Implementation(impl),
		godin.PrometheusCopnfig(godin.PrometheusConfig{
			Endpoint:         "/metrics",
			HistogramBuckets: prometheus.ExponentialBuckets(0.005, 1.4, 20),
			Network: godin.Network{
				Host: "0.0.0.0",
				Port: 9000,
			},
		}),
	)
}

func (e *exampleService) Greeting(ctx context.Context, request *greeter.GreetingRequest) (*greeter.GreetingResponse, error) {
	return &greeter.GreetingResponse{Greeting: "Hello there, " + request.Name}, nil
}

func (e *exampleService) Farewell(ctx context.Context, request *greeter.FarewellRequest) (*greeter.FarewellResponse, error) {
	return &greeter.FarewellResponse{Farewell: "Goodbye, " + request.Name}, nil
}
