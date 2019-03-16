package http

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type server struct {
	logger      *zap.SugaredLogger
	grpcPort    string
	httpPort    string
	context     context.Context
	server      *http.Server
	metricsPath string
}

func NewServer(logger *zap.SugaredLogger, grpcPort, httpPort string, metricsPath string) *server {
	return &server{
		logger:      logger,
		grpcPort:    grpcPort,
		httpPort:    httpPort,
		metricsPath: metricsPath,
	}
}

func (s *server) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	s.context = ctx
	defer cancel()

	mux := http.NewServeMux()

	// gRPC HTTP gateway
	var gwOpts []gwruntime.ServeMuxOption
	gw, err := newGateway(ctx, s.grpcPort, gwOpts)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	// add prometheus
	mux.Handle(s.metricsPath, promhttp.Handler())

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
