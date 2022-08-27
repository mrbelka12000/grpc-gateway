package main

import (
	handler "github.com/mrbelka12000/grpc-gateway/internal/delivery/http"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/internal/service/repo"
)

func main() {
	repo := repo.NewRepo()
	srv := service.New(repo.User, repo.Tracks)
	handler := handler.New(srv.User, srv.Tracks)
	handler.Check()
}
