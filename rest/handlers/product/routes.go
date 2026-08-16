package product

import (
	
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProuct), h.middlewares.AuthMiddleware))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProductByID)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middlewares.AuthMiddleware))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.AuthMiddleware))
}
