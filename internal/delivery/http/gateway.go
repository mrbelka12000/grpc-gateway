package http

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Gateway struct {
	server *http.Server
	notify chan error
}

func NewServer() *Gateway {
	mux := runtime.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	gateway := &Gateway{
		server: httpServer,
		notify: make(chan error),
	}

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

func (g *Gateway) RegisterHanlders() error {
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// localhost:12201 is address of grpc server
	// err := proto.RegisterGatewayHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	// if err != nil {
	// 	return err
	// }

	return nil
}
