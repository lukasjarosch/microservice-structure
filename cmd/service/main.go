package main

import (
	"os"

	"os/signal"
	"syscall"

	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure/pkg/grpc"
)

// Compile time variables are injected
var (
	GitCommit string
	GitBranch string
	BuildTime string
)

func main() {
	// perpare dependencies
	config := cfg.NewConfig()
	logger := initLogging(config.LogDebug)

	// setup our ExampleService
	service := svc.NewExampleService(config, logger)

	logger.WithFields(logrus.Fields{
		"instance": service.Options.ID,
		"git.commit": GitCommit,
		"git.branch": GitBranch,
		"build": BuildTime,
	}).Infof("starting service: %s", service.Options.Name)

	// goroutines
	go signalHandler(service)
	service.ServeMetrics()

	// and off we go ...
	if err := service.ServeGRPC(); err != nil {
		logger.Fatal(err)
	}
}

// initLogging initializes a new zap productionLogger and returns the sugared logger
func initLogging(debug bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}

// wait for SIGINT or SIGTERM and then call Shutdown()
func signalHandler(service *grpc.Server) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	signal.Notify(sigs, syscall.SIGTERM)
	logrus.Infof("signal: %v", <-sigs)

	service.Shutdown()
}
