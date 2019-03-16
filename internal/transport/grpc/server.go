package grpc

import (
	"context"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	greeter "github.com/lukasjarosch/microservice-structure/internal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	stop           chan os.Signal
	implementation greeter.HelloServer
	port           string
	logger         *zap.SugaredLogger
	server         *grpc.Server
}

// NewServer returns a configured gRPC server
// Interceptors to the server can be added through AddUnaryInterceptor() and AddStreamingInterceptor()
// Interceptors provide a way to inject middleware into the transport layer server (gRPC server).
func NewServer(logger *zap.SugaredLogger, service greeter.HelloServer, grpcPort string) *server {
	return &server{
		implementation: service,
		port:           grpcPort,
		logger:         logger,
	}
}

// Run opens the TCP port for the server, registers all interceptors, provides a graceful shutdown handler and then
// serves the gRPC server
func (s *server) Run() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	grpc_prometheus.EnableHandlingTimeHistogram()

	// create new gRPC server including middleware
	s.server = grpc.NewServer(

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.PayloadUnaryServerInterceptor(s.logger.Desugar(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
				// TODO: this method is used to decide whether to zap request/response payload or not
				// this method is called very often so keep it lightweight
				return true
			}),
			grpc_zap.UnaryServerInterceptor(s.logger.Desugar()),
			grpc_prometheus.UnaryServerInterceptor,
		)),

		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(s.logger.Desugar()),
			grpc_prometheus.StreamServerInterceptor,
		)),
	)

	// register service implementation
	greeter.RegisterHelloServer(s.server, s.implementation)

	s.logger.Infof("starting gRPC server on port %s", s.port)
	return s.server.Serve(listen)
}

func (s *server) GracefulShutdown() {
	s.server.GracefulStop()
}
