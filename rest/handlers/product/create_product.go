package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) CreateProuct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("JSON DECODE ERROR:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		ImageURL:    newProduct.ImageURL,
		Price:       newProduct.Price,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, 201)
}
