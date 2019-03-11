package grpc

import (
	"context"
	"github.com/lukasjarosch/microservice-structure/pkg"
	"net"
	"google.golang.org/grpc"
	"os"
	"github.com/sirupsen/logrus"
	"os/signal"
)

func RunServer(ctx context.Context, service greeter.HelloServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
	    return err
	}

	// register service
	server := grpc.NewServer()
	greeter.RegisterHelloServer(server, service)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logrus.Info("shutting down ...")
			server.GracefulStop()
			<- ctx.Done()
		}
	}()

	// start the gRPC server
	logrus.Infof("starting gRPC server on :%s ...", port)
	return server.Serve(listen)
}