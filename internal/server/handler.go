package server

import (
	"context"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/service"
)

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
