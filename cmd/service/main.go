package main

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/lukasjarosch/microservice-structure/pkg/transport/grpc"
	"github.com/lukasjarosch/microservice-structure/pkg/transport/http"
	"github.com/lukasjarosch/microservice-structure/pkg/transport/grpc/interceptors"
)

// Compile time variables are injected
var (
	GitCommit string
	GitBranch string
	BuildTime string
)

func main() {
	config := cfg.NewConfig(GitCommit, GitBranch, BuildTime)

	ctx := context.Background()

	// todo: config
	// todo: amqp
	// todo: mysql/mongodb

	service := svc.NewExampleService(config)

	go func() {
		http.RunServer(ctx, "50051", "8080")
	}()

	// setup gRPC and run
	grpcServer := grpc.NewServer(ctx, service, config.GrpcPort)
	grpcServer.AddUnaryInterceptor(interceptors.LogUnaryInterceptor())
	grpcServer.AddStreamingInterceptor(interceptors.LogStreamInterceptor())

	err := grpcServer.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
