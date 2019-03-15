package interceptors

import (
	"context"

	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

func LogUnaryInterceptor(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		entry := ctxlogrus.Extract(ctx)
		//entry := withContext(ctx, logger)

		entry.WithFields(log.Fields{
			"method":     info.FullMethod,
			"time.start": startTime.UTC().Format(time.UnixDate),
		}).Info("unary gRPC request")

		defer func(startTime time.Time) {
			entry.WithFields(log.Fields{
				"method":   info.FullMethod,
				"took":     time.Since(startTime),
				"time.end": time.Now().Format(time.UnixDate),
			}).Info("unary gRPC response")
		}(startTime)

		return handler(ctx, req)
	}
}

func LogStreamInterceptor(logger *log.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()
		entry := log.NewEntry(logger)

		entry.WithFields(log.Fields{
			"method":     info.FullMethod,
			"time.start": startTime.UTC().Format(time.UnixDate),
		}).Info("streaming gRPC request")

		defer func(startTime time.Time) {
			entry.WithFields(log.Fields{
				"method":   info.FullMethod,
				"took":     time.Since(startTime),
				"time.end": time.Now().Format(time.UnixDate),
			}).Info("streaming gRPC response")
		}(startTime)

		return handler(srv, ss)
	}
}

func withContext(ctx context.Context, logger *log.Logger) (entry *log.Entry) {
	peerInfo, ok := peer.FromContext(ctx)
	logger = logger.WithContext(ctx).Logger
	if ok {
		if peerInfo.Addr != nil {
			entry = logger.WithFields(log.Fields{
				"peer.address": peerInfo.Addr.String(),
				"peer.network": peerInfo.Addr.Network(),
			})
		}

		if peerInfo.AuthInfo != nil {
			entry = logger.WithFields(log.Fields{
				"peer.auth.type": peerInfo.AuthInfo.AuthType(),
			})
		}

	}

	return entry
}
