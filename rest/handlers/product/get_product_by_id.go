package product

import (
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

	product, err := h.svc.Get(id)

	if product == nil {
		util.SendError(w, 404, "product not found!")
		return
	}
 
	if err != nil {
		http.Error(w, "internal server error", 400)
		return
	}


	util.SendData(w, product, 200)

}
