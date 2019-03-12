package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func LogUnaryInterceptor() grpc.UnaryServerInterceptor {
	logger := log.New()
	entry := logger.WithFields(log.Fields{})

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(a()),
	}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.UnaryServerInterceptor(entry, opts...)
}

func LogStreamInterceptor() grpc.StreamServerInterceptor {
	logger := log.New()
	entry := logger.WithFields(log.Fields{})

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(a()),
	}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.StreamServerInterceptor(entry, opts...)
}

func a() grpc_logrus.CodeToLevel {
	return func(code codes.Code) log.Level {
		log.Error(code)
		return log.DebugLevel
	}
}
