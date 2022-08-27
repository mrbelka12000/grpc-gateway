package http

import (
	"fmt"

	"github.com/mrbelka12000/grpc-gateway/internal/service"
)

type Handler struct {
	us service.UserService
	ts service.TracksService
}

func New(us service.UserService, ts service.TracksService) *Handler {
	return &Handler{
		us: us,
		ts: ts,
	}
}

func (h *Handler) Check() {
	h.us.Login()
	fmt.Println("handler")
}
