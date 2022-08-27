package main

import (
	handler "github.com/mrbelka12000/grpc-gateway/internal/delivery/http"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/internal/service/repo"
)

func main() {
	srv := service.New(
		repo.NewUser(),
		repo.NewTracks())
	handler := handler.New(srv)
	handler.Check()
}
