package main

import (
	"fmt"
	"os"

	"os/signal"
	"syscall"

	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/lukasjarosch/microservice-structure/internal/transport/grpc"
	"github.com/lukasjarosch/microservice-structure/internal/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	http2 "net/http"
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

	// setup waitgroup with length 2 for http and grpc servers
	var wg sync.WaitGroup
	wg.Add(2)

	// todo: config
	// todo: amqp
	// todo: mysql/mongodb

	service := svc.NewExampleService(config, logger)
	logger.Infow("starting ExampleService", "git.commit", GitCommit, "git.branch", GitBranch, "build.date", BuildTime)

	// http gateway to gRPC server
	httpServer := http.NewServer(logger, config.GrpcPort, config.HttpPort)
	go func() {
		if err := httpServer.Run(); err != http2.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer wg.Done()
	}()

	// setup gRPC transport layer
	grpcServer := grpc.NewServer(logger, service, config.GrpcPort)
	go func() {
		err := grpcServer.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer wg.Done()
	}()

	go func() {
		waitForSignal()
		grpcServer.GracefulShutdown()
		httpServer.GracefulShutdown()
	}()

	wg.Wait()
	logger.Info("shut down")
}

// initLogging initializes a new zap productionLogger and returns the sugared logger
func initLogging(debug bool) *zap.SugaredLogger {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	level := zap.InfoLevel

	if debug {
		level = zap.DebugLevel
		pe.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
	l := zap.New(core)

	return l.Sugar()
}

// wait for SIGINT or SIGTERM
func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
