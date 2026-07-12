package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serv() {

	manager := middleware.NewManager()

	manager.Use(middleware.Hudai, middleware.Logger)

	mux := http.NewServeMux() // router

	 InitRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 3001")
	err := http.ListenAndServe(":3001", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
