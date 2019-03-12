package http

import "context"
import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/lukasjarosch/microservice-structure/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := greeter.RegisterHelloHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logrus.WithError(err).Fatal("failed to start HTTP gateway")
	}

	srv := &http.Server{
		Addr: ":" + httpPort,
		Handler: mux, // todo: add middleware/interceptors
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logrus.Info("shutting down HTTP gateway ...")
			// todo: properly handle
			<- ctx.Done()
		}
		_, cancel := context.WithTimeout(ctx, 5 * time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logrus.Infof("starting HTTP/JSON gateway on :%s...", httpPort)
	return srv.ListenAndServe()
}
