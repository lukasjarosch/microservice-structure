package config

import (
	"os"
)

type Config struct {
	LogLevel string
	GitCommit string
	GitBranch string
	BuildTime string
	GrpcPort string
	HttpPort string
}

func NewConfig(commit, branch, buildTime string) *Config {
	return &Config{
		GitCommit: commit,
		GitBranch: branch,
		BuildTime: buildTime,
		LogLevel: os.Getenv("LOG_LEVEL"),
		GrpcPort: "50051",
		HttpPort: "8080",
	}
}