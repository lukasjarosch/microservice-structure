package http

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Endpoint describes a single gRPC endpoint
type Endpoint struct {
	Network, Addr string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the listen address
	Addr string
	// GRPCServe defines an endpoint of the gRPC service
	GRPCServer Endpoint
	// Mux defines the options passed to the grpc-gateway multiplexer
	Mux []gwruntime.ServeMuxOption
	// Logger
	Logger *zap.SugaredLogger
	// MetricsPath is the endpoint where the prometheus handler handles on
	MetricsPath string
}

// Run will start an HTTP server and blocks while running
// Once the context is cancelled, the server will shut down
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Addr)
	if err != nil {
	    return err
	}

	go func() {
		<-ctx.Done()
		opts.Logger.Debug("closing HTTP connections")
		if err := conn.Close(); err != nil {
			opts.Logger.Errorw("failed to close a client connection to the gRPC server", "err", err)
		}
		opts.Logger.Debug("HTTP connections closed")
	}()

	// setup multiplexer
	mux := http.NewServeMux()
	mux.Handle(opts.MetricsPath, promhttp.Handler())

	// gRPC HTTP gateway
	gw, err := newGateway(ctx, conn, opts.Mux)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)


	srv := &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}

	go func() {
		<- ctx.Done()
		opts.Logger.Info("shutting down HTTP server")
		if err := srv.Shutdown(context.Background()); err != nil {
			opts.Logger.Errorw("failed to shutdown HTTP server", "err", err)
		}
		opts.Logger.Info("HTTP server shut down")
	}()

	opts.Logger.Infof("HTTP server listening on %s", opts.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		opts.Logger.Errorw("failed to listen and serve", "err", err)
		return err
	}
	return nil
}

