package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser), middleware.Logger))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.Login), middleware.Logger))
}
