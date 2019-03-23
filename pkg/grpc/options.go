package grpc

import (
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"strings"
)

// Option defines a settable option of the gRPC server
type Option func(*Options)

// Options define all settable options of a server
type Options struct {

	// plain service name
	Name string

	// unique service instance id (generated when setting Name)
	ID   string

	// gRPC interceptors
	UnaryInterceptors     []grpc.UnaryServerInterceptor
	StreamingInterceptors []grpc.StreamServerInterceptor

	// gRPC implementation (your business logic)
	GRPCImplementation GRPCImplementation

	// configuration
	GRPCOptions      []grpc.ServerOption
	ServerConfig     GrpcServerConfig
	PrometheusConfig PrometheusConfig
}

// Name will set the service Name and regenerate the ID
func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
		opts.ID = fmt.Sprintf("%s-%s", name, uuid.New().String())
	}
}

// AddUnaryInterceptor adds a grpc UnaryServerInterceptor to the interceptor-chain
func AddUnaryInterceptor(interceptor grpc.UnaryServerInterceptor) Option {
	return func(opts *Options) {
		opts.UnaryInterceptors = append(opts.UnaryInterceptors, interceptor)
	}
}

// AddStreamingInterceptor adds a grpc StreamServerInterceptor to the interceptor-chain
func AddStreamingInterceptor(interceptor grpc.StreamServerInterceptor) Option {
	return func(opts *Options) {
		opts.StreamingInterceptors = append(opts.StreamingInterceptors, interceptor)
	}
}

// Implementation will set a GRPCImplementation
func Implementation(impl GRPCImplementation) Option {
	return func(opts *Options) {
		opts.GRPCImplementation = impl
	}
}

// GrpcNetworkPort sets the port on which the gRPC server will listen on
func GrpcNetworkPort(port int) Option {
	return func(opts *Options) {
		opts.ServerConfig.Network.Port = port
	}
}

// PrometheusEndpoint sets the endpoint on which the prometheus metrics will be exposed
func PrometheusEndpoint(endpoint string) Option {
	return func(opts *Options) {
		if !strings.HasPrefix(endpoint, "/") {
			endpoint = "/" + endpoint
		}
		opts.PrometheusConfig.Endpoint = endpoint
	}
}

// PrometheusNetworkPort sets the port of the HTTP server on which prometheus is running
func PrometheusNetworkPort(port int) Option {
	return func(opts *Options) {
		opts.PrometheusConfig.Network.Port = port
	}
}

// PrometheusHistogramBuckets set the bucket list which is used for the default latency histogram
// The prometheus library provides convenient methods to create these buckets.
func PrometheusHistogramBuckets(buckets []float64) Option {
	return func(opts *Options) {
		opts.PrometheusConfig.HistogramBuckets = buckets
	}
}
