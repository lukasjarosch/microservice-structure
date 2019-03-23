package main

import (
	"os"

	"os/signal"
	"syscall"

	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure/pkg/server"
)

// Compile time variables are injected
var (
	GitCommit string
	GitBranch string
	BuildTime string
)

func main() {
	config := cfg.NewConfig()
	logger := initLogging(config.LogDebug)

	service := svc.NewExampleService(config, logger)

	go signalHandler(service)

	service.ServeMetrics()

	if err := service.ServerGRPC(); err != nil {
		logger.Fatal(err)
	}
}

// initLogging initializes a new zap productionLogger and returns the sugared logger
func initLogging(debug bool) *logrus.Entry {
	logger := logrus.New()

	return logrus.NewEntry(logger)
}

// wait for SIGINT or SIGTERM and then call Shutdown()
func signalHandler(service *server.Server) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	signal.Notify(sigs, syscall.SIGTERM)
	logrus.Infof("signal: %v", <-sigs)

	service.Shutdown()
}
