package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
)

// Config holds the service configuration
type Config struct {
	LogDebug    bool   `env:"LOG_DEBUG" envDefault:"false"`
	GrpcPort    string `env:"GRPC_PORT" envDefault:"50051"`
	HttpPort    string `env:"HTTP_PORT" envDefault:"8080"`
	MetricsPath string `env:"METRICS_PATH" envDefault:"/metrics"`
}

// NewConfig returns a new Config. The configuration is parsed from environment variables.
// Default values are only set if an environment variable is not set
func NewConfig() *Config {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("unable to parse config: %v", err)
		os.Exit(1)
	}
	return cfg
}
