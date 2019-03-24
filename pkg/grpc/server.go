package grpc

import (
	"fmt"
	"net"
	"net/http"

	"context"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	 "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/lukasjarosch/microservice-structure/pkg/config"
)

type GRPCImplementation func(s *grpc.Server)

// Server is used as wrapper to provide a good starting point to create gRPC based services
type Server struct {
	Options              *Options
	GRPCServer           *grpc.Server
	PrometheusHTTPServer *http.Server
}

// NewServer constructs a new Server with the given options
func NewServer(opts ...Option) *Server {
	options := &Options{
		ID:   fmt.Sprintf("%s-%s", "godin", uuid.New().String()),
		Name: "godin",
		ServerConfig: GrpcServerConfig{
			Network: config.Network{
				Host: "0.0.0.0",
				Port: 50051,
			}},
		PrometheusConfig: PrometheusConfig{
			Network: config.Network{
				Host: "0.0.0.0",
				Port: 9000,
			},
			Endpoint: "/metrics",
			HistogramBuckets: prometheus.ExponentialBuckets(0.005, 1.4, 20),
	},
		GRPCImplementation: func(s *grpc.Server) {},
		UnaryInterceptors: []grpc.UnaryServerInterceptor{
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		},
		StreamingInterceptors: []grpc.StreamServerInterceptor{
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		},
	}

	for _, option := range opts {
		option(options)
	}

	return &Server{
		Options: options,
	}
}

// Serve will serve a gRPC server
func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", s.Options.ServerConfig.Network.Address())
	if err != nil {
		return err
	}

	logrus.Infof("serving gRPC server at %s", s.Options.ServerConfig.Network.Address())
	return s.createGrpcServer().Serve(listener)
}

// ServeMetrics will start up a HTTP server and attach the prometheus handler to it
// The server will ListenAndServe() in a goroutine
func (s *Server) ServeMetrics() {
	s.PrometheusHTTPServer = &http.Server{Addr: s.Options.PrometheusConfig.Network.Address()}

	http.Handle(s.Options.PrometheusConfig.Endpoint, promhttp.Handler())

	logrus.Infof("serving prometheus metrics at http://%s%s",
		s.Options.PrometheusConfig.Network.Address(),
		s.Options.PrometheusConfig.Endpoint)

	go func() {
		if err := s.PrometheusHTTPServer.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Errorf("prometheus http server: ListenAndServe() error: %s", err)
			return
		}
	}()
}

// Shutdown handles the graceful shutdown of the gRPC server and the Prometheus HTTP server
func (s *Server) Shutdown() {
	s.GRPCServer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.PrometheusHTTPServer.Shutdown(ctx); err != nil {
		logrus.Infof("timeout during shutdown of prometheus HTTP server: %v", err)
	}
}

// Setup the gRPC server. Add all interceptors, register with prometheus and return a new grpc.Server
func (s *Server) createGrpcServer() *grpc.Server {
	s.Options.GRPCOptions = append(s.Options.GRPCOptions, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(s.Options.UnaryInterceptors...),
	))

	s.Options.GRPCOptions = append(s.Options.GRPCOptions, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(s.Options.StreamingInterceptors...),
	))

	s.GRPCServer = grpc.NewServer(
		s.Options.GRPCOptions...,
	)

	s.Options.GRPCImplementation(s.GRPCServer)

	// adjust histogram buckets
	grpc_prometheus.EnableHandlingTimeHistogram(
		func(opts *prometheus.HistogramOpts) {
			opts.Buckets = s.Options.PrometheusConfig.HistogramBuckets
		},
	)

	grpc_prometheus.Register(s.GRPCServer)
	return s.GRPCServer
}
