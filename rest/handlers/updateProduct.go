package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id ", 400)
		return
	}

	var newProduct database.Product

	fmt.Println(newProduct)

	decoder := json.NewDecoder(r.Body)
	

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}
	newProduct.ID = pID

	fmt.Println(newProduct)

	database.Update(newProduct)

	util.SendData(w, "Successfully updated prodcut", 201)
}
