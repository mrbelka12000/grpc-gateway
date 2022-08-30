package service

import (
	"context"
	"github.com/mrbelka12000/grpc-gateway/proto"
)

type (
	Fetch interface {
		GetInfoByIIN(ctx context.Context, message *proto.Message) (*proto.Response, error)
	}
)
