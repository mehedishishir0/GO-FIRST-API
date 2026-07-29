package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serv() {
	cnf := config.GetConfig()

	rest.Strat(cnf)

}
