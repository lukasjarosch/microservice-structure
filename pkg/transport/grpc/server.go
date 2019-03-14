package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	greeter "github.com/lukasjarosch/microservice-structure/pkg"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	ctx                 context.Context
	implementation      greeter.HelloServer
	port                string
	unaryMiddleware     []grpc.UnaryServerInterceptor
	streamingMiddleware []grpc.StreamServerInterceptor
}

// NewServer returns a configured gRPC server
func NewServer(ctx context.Context, service greeter.HelloServer, grpcPort string) *server {
	return &server{
		ctx:            ctx,
		implementation: service,
		port:           grpcPort,
	}
}
// AddUnaryInterceptor registers an unary interceptor
func (s *server) AddUnaryInterceptor(interceptor grpc.UnaryServerInterceptor) {
	s.unaryMiddleware = append(s.unaryMiddleware, interceptor)
}

// AddStreamingInterceptor registers a streaming interceptor
func (s *server) AddStreamingInterceptor(interceptor grpc.StreamServerInterceptor) {
	s.streamingMiddleware = append(s.streamingMiddleware, interceptor)
}

// Run opens the TCP port for the server, registers all interceptors, provides a graceful shutdown handler and then
// serves the gRPC server
func (s *server) Run() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	// create new gRPC server including middleware
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.unaryMiddleware...,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			s.streamingMiddleware...,
		)),
	)

	greeter.RegisterHelloServer(server, s.implementation)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Info("shutting down gRPC server ...")
			server.GracefulStop()
			<-s.ctx.Done()
		}
	}()

	// start the gRPC server
	log.Infof("starting gRPC server on :%s ...", s.port)
	return server.Serve(listen)
}
