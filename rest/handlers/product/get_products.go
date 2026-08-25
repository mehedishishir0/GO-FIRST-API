package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	limit := reqQuery.Get("limit")
	page := reqQuery.Get("page")
	limitInt, _ := strconv.ParseInt(limit, 10, 32)
	pageInt, _ := strconv.ParseInt(page, 10, 32)

	if limitInt == 0 {
		limitInt = 10
	}
	if pageInt == 0 {
		pageInt = 1
	}

	if r.Method != "GET" {
		http.Error(w, "Please hit the get method", 400)
		return
	}

	product, err := h.svc.List(pageInt, limitInt)

	if product == nil {
		util.SendError(w, 404, "product not found!")
		return
	}

	if err != nil {
		http.Error(w, "internal server error", 400)
		return
	}

	cnt, err := h.svc.Count()

	if err != nil {
		http.Error(w, "internal server error", 400)
		return
	}

	util.SendPage(w, product, pageInt, limitInt, cnt)

}
