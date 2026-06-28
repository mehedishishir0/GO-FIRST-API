package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"log"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

  id, err := strconv.Atoi(productID)

  if err != nil {
	http.Error(w, err.Error(), 400)
	return
  }

for idx, product := range database.ProductList {
 log.Println(idx, product)

 if product.ID == id {
  util.SendData(w, product, 200)
  return
 }
}

http.Error(w, "Product not found", 404)

}
