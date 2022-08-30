package grpc

import (
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	im proto.GatewayServer
}

func New(f service.Fetch) *GrpcServer {
	return &GrpcServer{
		im: newImpelementer(f),
	}
}

func (i *GrpcServer) Register(server *grpc.Server) {
	proto.RegisterGatewayServer(server, i.im)
}
