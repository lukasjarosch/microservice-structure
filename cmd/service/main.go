package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"os/signal"
	"syscall"

	"sync"

	cfg "github.com/lukasjarosch/microservice-structure/internal/config"
	svc "github.com/lukasjarosch/microservice-structure/internal/service"
	"github.com/lukasjarosch/microservice-structure/internal/transport/grpc"
	"github.com/lukasjarosch/microservice-structure/internal/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	// setup service implementation (gRPC server handler)
	service := svc.NewExampleService(config, logger)
	logger.Infow("starting ExampleService", "git.commit", GitCommit, "git.branch", GitBranch, "build.date", BuildTime)

	// setup http gateway including prometheus metrics
	httpCancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		defer wg.Done()
		http.Run(httpCancelCtx, http.Options{
			Logger:      logger,
			MetricsPath: config.MetricsPath,
			Addr:        fmt.Sprintf(":%s", config.HttpPort),
			GRPCServer: http.Endpoint{
				Network: "tcp",
				Addr:    fmt.Sprintf(":%s", config.GrpcPort),
			},
		})
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

	// handle signals gracefully
	go func() {
		waitForSignal()
		grpcServer.GracefulShutdown()
		httpCancelCtx.Done()
	}()

	wg.Wait()
	os.Exit(0)
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
