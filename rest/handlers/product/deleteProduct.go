package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id ", 400)
		return
	}
	database.Delete(pID)

	util.SendData(w, "Successfully deleted prodcut", 201)
}
