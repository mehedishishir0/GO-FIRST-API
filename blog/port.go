package blog

import (
	"ecommerce/domain"
	blogHandler "ecommerce/rest/handlers/blog" 
)

type Service interface {
	blogHandler.Service
}

type BlogRepo interface {
	Create(b domain.Blog) (*domain.Blog, error)
}

