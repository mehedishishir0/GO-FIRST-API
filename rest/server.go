package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Strat(cnf config.Config) {
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
