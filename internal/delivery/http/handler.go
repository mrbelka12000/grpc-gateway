package http

import (
	"fmt"
	"github.com/mrbelka12000/grpc-gateway/internal/models"
	"github.com/mrbelka12000/grpc-gateway/internal/service"
	"net/http"
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

func (h *Handler) CallBack(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	ip := r.RemoteAddr
	loginResp := &models.LoginResp{
		Code:  values.Get("code"),
		State: values.Get("state"),
	}
	err := h.us.SaveLoginCode(ip, loginResp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("code = " + h.us.GetLoginCode(ip))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	url, err := h.us.GetLoginUrl(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, url, 301)
}
