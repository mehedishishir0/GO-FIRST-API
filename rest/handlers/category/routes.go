package category


import (
	
	"net/http"
	middleware "ecommerce/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	
	mux.Handle("POST /categories", manager.With(http.HandlerFunc(h.createCategory), h.middlewares.AuthMiddleware))
}
