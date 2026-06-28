package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateProuct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
