package server

import (
	"fmt"
)

type Network struct {
	Port int
	Host string
}

func (n *Network) Address() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

// Config holds the gRPC server configuration
type Config struct {
	Network Network
}

// PrometheusConfig holds the prometheus HTTP server configuration
type PrometheusConfig struct {
	Network Network
	Endpoint string
	HistogramBuckets []float64
}

