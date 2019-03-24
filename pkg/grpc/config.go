package grpc

import (
	common "github.com/lukasjarosch/microservice-structure/pkg/config"
)

// GrpcServerConfig holds the gRPC server configuration
type GrpcServerConfig struct {
	Network common.Network
}

// PrometheusConfig holds the prometheus HTTP server configuration
type PrometheusConfig struct {
	Network          common.Network
	Endpoint         string
	HistogramBuckets []float64
}
