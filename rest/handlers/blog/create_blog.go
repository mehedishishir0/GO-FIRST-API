package blog

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateBlog

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" || req.Author == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	createBlog, err := h.svc.Create(domain.Blog{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, createBlog, http.StatusCreated)
}
