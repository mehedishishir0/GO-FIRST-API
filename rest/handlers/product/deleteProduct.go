package product

import (
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

	err = h.svc.Delete(pID)

	if err != nil {
		http.Error(w, "Internal server error ", 500)
		return
	}

	util.SendData(w, "Successfully deleted prodcut", 201)
}
