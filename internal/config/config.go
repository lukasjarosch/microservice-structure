package config

import (
	"github.com/caarlos0/env"
	"log"
	"os"
)

type Config struct {
	LogDebug  bool `env:"LOG_DEBUG" envDefault:"false"`
	GrpcPort  string `env:"GRPC_PORT" envDefault:"50051"`
	HttpPort  string `env:"HTTP_GATEWAY_PORT" envDefault:"8080"`
}

func NewConfig() *Config {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("unable to parse config: %v", err)
		os.Exit(1)
	}
	return cfg
}
