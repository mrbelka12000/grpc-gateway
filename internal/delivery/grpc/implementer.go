package grpc

import (
	"context"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/proto"
)

type Implementer struct {
	f service.Fetch
	proto.UnimplementedGatewayServer
}

func newImpelementer(f service.Fetch) *Implementer {
	return &Implementer{
		f: f,
	}
}

func (i *Implementer) GetInfoByIIN(ctx context.Context, message *proto.Message) (*proto.Response, error) {
	return i.f.GetInfoByIIN(ctx, message)
}
