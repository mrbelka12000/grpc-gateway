package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mrbelka12000/grpc-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	//"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Gateway struct {
	server *http.Server
	notify chan error
}

func New() *Gateway {
	mux := runtime.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	gateway := &Gateway{
		server: httpServer,
		notify: make(chan error),
	}

	gateway.registerHandlers(mux)

	go gateway.start()

	return gateway
}

func (g *Gateway) start() {
	go func() {
		g.notify <- g.server.ListenAndServe()
		close(g.notify)
	}()
}

func (g *Gateway) Notify() <-chan error {
	return g.notify
}

func (g *Gateway) registerHandlers(mux *runtime.ServeMux) error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, "localhost:50511", opts)
	if err != nil {
		return err
	}
	return nil
}

func (g *Gateway) Shutdown() error {
	return g.server.Shutdown(context.Background())
}
