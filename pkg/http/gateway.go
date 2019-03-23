package http

import (
	"context"

	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpcServer "github.com/lukasjarosch/microservice-structure/pkg/grpc"
	"google.golang.org/grpc"
)

type GatewayHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

type Gateway struct {
	Context           context.Context
	Handler           []GatewayHandler
	Options           []runtime.ServeMuxOption
	GrpcServerNetwork grpcServer.Network
}

func NewGateway(ctx context.Context, grpcServerNetwork grpcServer.Network, handler []GatewayHandler, opts []runtime.ServeMuxOption) *Gateway {
	return &Gateway{
		Context:           ctx,
		Handler:           handler,
		GrpcServerNetwork: grpcServerNetwork,
		Options:           opts,
	}
}

func (gw *Gateway) HttpHandler() (http.Handler, error) {
	mux := runtime.NewServeMux(gw.Options...)

	client, err := grpc.DialContext(gw.Context, gw.GrpcServerNetwork.Address(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	for _, handler := range gw.Handler {
		if err := handler(gw.Context, mux, client); err != nil {
			return nil, err
		}
	}

	return mux, nil
}

// GatewayServer is a convenience method which returns a pre-initialized http-gateway server
func GatewayServer(grpcNetwork grpcServer.Network, handler GatewayHandler) (server *Server, err error) {
	server = NewServer()

	gateway := NewGateway(context.Background(),
		grpcNetwork,
		[]GatewayHandler{handler},
		nil)

	gwHandler, err := gateway.HttpHandler()
	if err != nil {
		return nil, err
	}

	server.AddEndpoint(Endpoint{
		Pattern: "/",
		Handler: gwHandler,
	})

	return server, nil
}
