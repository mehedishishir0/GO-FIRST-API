package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
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

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)
	
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)

	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
