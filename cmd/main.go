package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrbelka12000/grpc-gateway/internal/delivery/http"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"github.com/mrbelka12000/grpc-gateway/internal/service/repo"
)

func main() {
	godotenv.Load()
	srv := service.New(
		repo.NewUser(),
		repo.NewTracks())

	handler := http.New(
		srv.UserService,
		srv.TracksService)

	fmt.Println("server started")
	http.StartServer(handler)
}
