package category

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateCategory struct {
	Title       string `json:"name" db:"title"`
	Description string `json:"description" db:"description"`
	ImageURL    string `json:"imageUrl" db:"img_url"`
}

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory ReqCreateCategory

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newCategory)

	if err != nil {
		fmt.Println("JSON DECODE ERROR:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createCategory, err := h.svc.Create(domain.Category{
		Title:       newCategory.Title,
		Description: newCategory.Description,
		ImageURL:    newCategory.ImageURL,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, createCategory, 201)

}
