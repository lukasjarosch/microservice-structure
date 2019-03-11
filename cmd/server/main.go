package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lukasjarosch/microservice-structure/pkg/grpc"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/lukasjarosch/microservice-structure/pkg/http"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// Run runs the gRPC server including the HTTP gateway
func Run() error {
	ctx := context.Background()

	// todo: config
	// todo: amqp
	// todo: mysql/mongodb

	service := svc.NewExampleService()

	go http.RunServer(ctx, "50051", "8080")
	return grpc.RunServer(ctx, service, "50051")
}
