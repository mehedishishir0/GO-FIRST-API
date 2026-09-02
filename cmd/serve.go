package cmd

import (
	"ecommerce/category"
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	categoryHandler "ecommerce/rest/handlers/category"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"

	"os"
)

func Serv() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)
	cateRepo := repo.NewCategoryRepo(dbCon)

	//domains
	userService := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)
	cateSvc := category.NewService(cateRepo)

	// middlewares
	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHandler := productHandler.NewHandler(middlewares, productSvc)
	userHandler := userHandler.NewHandler(userService, cnf)
    categoryHandler := categoryHandler.NewHandler(middlewares, cateSvc)

	server := rest.NewServer(cnf, productHandler, userHandler, categoryHandler)
	server.Start()
}
