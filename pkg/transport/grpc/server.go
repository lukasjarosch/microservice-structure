package grpc

import (
	"context"
	"github.com/lukasjarosch/microservice-structure/pkg"
	"net"
	"google.golang.org/grpc"
	"os"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/lukasjarosch/microservice-structure/pkg/transport/grpc/interceptors"
)

func RunServer(ctx context.Context, service greeter.HelloServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
	    return err
	}

	// create new gRPC server including middleware
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptors.LogUnaryInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			interceptors.LogStreamInterceptor(),
		)),
	)

	greeter.RegisterHelloServer(server, service)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Info("shutting down gRPC server ...")
			server.GracefulStop()
			<- ctx.Done()
		}
	}()

	// start the gRPC server
	log.Infof("starting gRPC server on :%s ...", port)
	return server.Serve(listen)
}