package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serv() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//domains
	userService := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := productHandler.NewHandler(middlewares, productRepo)

	userHandler := userHandler.NewHandler(userService, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
