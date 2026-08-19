package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	conf           *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		conf:           cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	wrappedmux := manager.WrapMux(mux)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(s.conf.HttpPort)

	fmt.Println("Server is running on port ", addr)
	err := http.ListenAndServe(addr, wrappedmux)

	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
