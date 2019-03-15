package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/lukasjarosch/microservice-structure/internal"
)

func RunServer(ctx context.Context, logger *zap.SugaredLogger, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := internal.RegisterHelloHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Errorw("failed to start HTTP gateway", "err", err)
		return err
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux, // todo: add middleware/interceptors
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Info("shutting down HTTP gateway ...")
			// todo: properly handle
			<-ctx.Done()
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Infof("starting HTTP/JSON gateway on :%s...", httpPort)
	return srv.ListenAndServe()
}
