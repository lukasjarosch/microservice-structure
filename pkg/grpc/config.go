package grpc

import (
	"fmt"
)

type Network struct {
	Port int
	Host string
}

// Address returns the concatenated Host:Port combination as string
func (n *Network) Address() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

// GrpcServerConfig holds the gRPC server configuration
type GrpcServerConfig struct {
	Network Network
}

// PrometheusConfig holds the prometheus HTTP server configuration
type PrometheusConfig struct {
	Network Network
	Endpoint string
	HistogramBuckets []float64
}

