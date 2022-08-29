package grpc

import (
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/proto"
)

type Implementer struct {
	proto.UnimplementedRouteGuideServer
	f service.Fetch
}

func newImpelementer(f service.Fetch) *Implementer {
	return &Implementer{
		f: f,
	}
}

func (i *Implementer) GetInfoByIIN(iin string) {
}
