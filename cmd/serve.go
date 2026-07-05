package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serv() {
	mux := http.NewServeMux() // router


  cntrl := func (w http.ResponseWriter, r *http.Request)  {
	
  }

  handler := http.HandlerFunc(cntrl)


    
	mux.Handle("/GET route", middleware.Logger(handler))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProuct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 3001")
	err := http.ListenAndServe(":3001", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
