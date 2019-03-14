package interceptors

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LogUnaryInterceptor() grpc.UnaryServerInterceptor {
	logger := log.New()
	entry := logger.WithFields(log.Fields{})

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.UnaryServerInterceptor(entry)
}

func LogStreamInterceptor() grpc.StreamServerInterceptor {
	logger := log.New()
	entry := logger.WithFields(log.Fields{})

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(entry)

	return grpc_logrus.StreamServerInterceptor(entry)
}
