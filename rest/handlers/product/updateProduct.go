package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id ", 400)
		return
	}

	var newProduct ReqUpdateProduct

	fmt.Println(newProduct)

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

 _, err =	h.productRepo.Update(repo.Product{
		ID:          pID,
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImageURL:  newProduct.ImageURL,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError , "internel server error")
	}

	util.SendData(w, "Successfully updated prodcut", 201)
}
