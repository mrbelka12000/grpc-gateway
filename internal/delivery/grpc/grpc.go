package grpc

import (
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	f service.Fetch
}

func New(f service.Fetch) *GrpcServer {
	return &GrpcServer{f: f}
}

func (i *GrpcServer) Register(server *grpc.Server) {
	proto.RegisterRouteGuideServer(server, nil)
}
