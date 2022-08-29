package app

import (
	"fmt"
	"log"
	"net"

	gateway "github.com/mrbelka12000/grpc-gateway/internal/delivery/grpc"
	"github.com/mrbelka12000/grpc-gateway/internal/delivery/http"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"google.golang.org/grpc"
)

func Initialize() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50511))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	srvImp := service.New()
	gatewayImp := gateway.New(srvImp)
	gatewayImp.Register(grpcServer)

	gateway := http.New()
	go grpcServer.Serve(lis)
	fmt.Println("grpc server started")
	fmt.Println("http server started")
	err = <-gateway.Notify()
	if err != nil {
		log.Println(err)
	}
}
