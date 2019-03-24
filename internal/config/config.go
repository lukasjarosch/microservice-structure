package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure/pkg/config"
)

// ServerConfig holds the service configuration
type Config struct {
	LogDebug        bool   `envconfig:"LOG_DEBUG" default:"false"`
	GrpcPort        int    `envconfig:"GRPC_PORT" default:"50051"`
	PrometheusPort  int    `envconfig:"PROMETHEUS_PORT" default:"9000"`
	MetricsEndpoint string `envconfig:"METRICS_ENDPOINT" default:"/metrics"`
	MongoDB config.MongoDBConfiguration
}

// NewConfig returns a new ServerConfig. The configuration is parsed from environment variables.
// Default values are only set if an environment variable is not set
func NewConfig() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
	    logrus.WithError(err).Fatal("unable to process configuration")
	}

	return &cfg
}
