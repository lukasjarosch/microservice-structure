package server

import (
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

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
	GRPCOptions        []grpc.ServerOption
	Config           Config
	PrometheusConfig PrometheusConfig
}

// Name will set the service Name and regenerate the ID
func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
		opts.ID = fmt.Sprintf("%s-%s", name, uuid.New().String())
	}
}

// Implementation will set a GRPCImplementation
func Implementation(impl GRPCImplementation) Option {
	return func(opts *Options) {
		opts.GRPCImplementation = impl
	}
}

func PrometheusCopnfig(config PrometheusConfig) Option {
	return func(opts *Options) {
		opts.PrometheusConfig = config
	}
}
