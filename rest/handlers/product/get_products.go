package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Please hit the get method", 400)
		return
	}

	product, err := h.productRepo.List()

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
