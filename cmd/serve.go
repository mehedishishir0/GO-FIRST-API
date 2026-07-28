package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serv() {
	cnf := config.GetConfig()

	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	wrappedmux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server is running on port ", addr)
	err := http.ListenAndServe(addr, wrappedmux)

	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
