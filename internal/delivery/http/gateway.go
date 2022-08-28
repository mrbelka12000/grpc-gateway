package http

import "net/http"

func StartServer(h *Handler) {

	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/callback", h.CallBack)
	http.ListenAndServe(":8080", nil)
}
