package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lukasjarosch/microservice-structure/internal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	logger   *zap.SugaredLogger
	grpcPort string
	httpPort string
	context  context.Context
	server *http.Server
}

func NewServer(logger *zap.SugaredLogger, grpcPort, httpPort string) *server {
	return &server{
		logger:   logger,
		grpcPort: grpcPort,
		httpPort: httpPort,
	}
}

func (s *server) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	s.context = ctx
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := internal.RegisterHelloHandlerFromEndpoint(ctx, mux, "localhost:"+s.grpcPort, opts); err != nil {
		s.logger.Errorw("failed to start HTTP gateway", "err", err)
		return err
	}

	s.server = &http.Server{
		Addr:    ":" + s.httpPort,
		Handler: mux, // todo: add middleware/interceptors
	}

	s.logger.Infof("starting HTTP/JSON gateway on port %s", s.httpPort)
	return s.server.ListenAndServe()
}

func (s *server) GracefulShutdown() {
	s.server.Shutdown(s.context)
}
