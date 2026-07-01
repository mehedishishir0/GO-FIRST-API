package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serv() {

	manager := middleware.NewManager()



	mux := http.NewServeMux() // router

  m :=	manager.With(middleware.Hudai, middleware.Logger)

	mux.Handle("GET /mehedi", manager.With() )

	mux.Handle("GET /test", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProuct))) 
	mux.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 3001")
	err := http.ListenAndServe(":3001", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
