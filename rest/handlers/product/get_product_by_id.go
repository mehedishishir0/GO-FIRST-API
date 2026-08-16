package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	product := database.Get(id)

	if product == nil {
		util.SendError(w, 404, "product not found!")
		return
	}

	util.SendData(w, product, 200)

}
