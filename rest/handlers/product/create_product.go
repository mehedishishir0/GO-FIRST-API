package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}


func (h *Handler) CreateProuct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		
		Title: newProduct.Title,
		Description: newProduct.Description,
		ImageURL: newProduct.ImageURL,
		Price: newProduct.Price,
	})
	if err != nil{
		http.Error(w, "Internal Server error ", http.StatusInternalServerError)
	}

	util.SendData(w, createdProduct, 201)
}
