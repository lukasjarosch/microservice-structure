package main

import (
	"context"
	"fmt"
	"os"

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
	config := cfg.NewConfig(GitCommit, GitBranch, BuildTime)
	logger := initLogging(config.LogDebug)

	// internal context to handle graceful shutdowns
	ctx := context.Background()

	// todo: config
	// todo: amqp
	// todo: mysql/mongodb

	service := svc.NewExampleService(config, logger)

	logger.Debug("testing")

	go func() {
		http.RunServer(ctx, logger, config.GrpcPort, config.HttpPort)
	}()

	// setup gRPC transport layer
	grpcServer := grpc.NewServer(ctx, logger, service, config.GrpcPort)

	err := grpcServer.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
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
