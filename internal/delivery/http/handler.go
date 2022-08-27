package http

import (
	"fmt"

	"github.com/mrbelka12000/grpc-gateway/internal/service"
)

type Handler struct {
	us service.UserService
	ts service.TracksService
}

func New(srv *service.Service) *Handler {
	return &Handler{
		us: srv.UserService,
		ts: srv.TracksService,
	}
}

func (h *Handler) Check() {
	h.us.Login()
	fmt.Println("handler")
}
