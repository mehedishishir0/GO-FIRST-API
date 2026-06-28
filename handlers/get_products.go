package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Please hit the get method", 400)
		return
	}

	util.SendData(w, database.ProductList, 200)

}
