package grpc

import (
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	im protoInterface
}

// temp inteface
type protoInterface interface {
	GetInfoByIIN(iin string)
}

func New(f service.Fetch) *GrpcServer {
	return &GrpcServer{im: newImpelementer(f)}
}

func (i *GrpcServer) Register(server *grpc.Server) {
	proto.RegisterRouteGuideServer(server, nil)
}
