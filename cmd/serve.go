package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serv() {
   println("ami hoalm serve")

	mux := http.NewServeMux() // router
	
	manager := middleware.NewManager()


	 manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	wrappedmux := manager.WrapMux( mux)

	InitRoutes(mux, manager)

	fmt.Println("Server is running on port 3001")
	err := http.ListenAndServe(":3001", wrappedmux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
