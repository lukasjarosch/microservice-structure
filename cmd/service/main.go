package main

import (
	"os"

	"os/signal"
	"syscall"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/lukasjarosch/microservice-structure/pkg/grpc"
	"github.com/lukasjarosch/microservice-structure/pkg/http"
	"github.com/sirupsen/logrus"
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
		"instance":   service.Options.ID,
		"git.commit": GitCommit,
		"git.branch": GitBranch,
		"build":      BuildTime,
	}).Infof("starting service: %s", service.Options.Name)

	// If you want to have a HTTP/JSON gateway you can easily start it up like this
	gatewayServer, err := http.GatewayServer(service.Options.ServerConfig.Network, greeter.RegisterHelloHandler)
	if err != nil {
		logger.Fatal("failed to start the HTTP gateway server: %v", err)
		os.Exit(-1)
	}
	gatewayServer.ServeHTTP()

	// graceful shutdown using signals (SIGINT and SIGTERM)
	go shutdownHandler(service, gatewayServer)

	// HTTP server providing Prometheus metrics
	service.ServeMetrics()

	// finally: serve the gRPC server in the foreground
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
func shutdownHandler(service *grpc.Server, gateway *http.Server) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	signal.Notify(sigs, syscall.SIGTERM)
	logrus.Infof("signal: %v", <-sigs)

	service.Shutdown()
	gateway.Shutdown()
}
