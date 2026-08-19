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

	util.SendData(w, h.productRepo.List(), 200)

}
