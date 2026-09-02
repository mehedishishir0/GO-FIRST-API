package category

import (
	"ecommerce/domain"
	ctgHandler "ecommerce/rest/handlers/category"
)

type Service interface {
	ctgHandler.Service
}

type CategoryRepo interface {
	Create(domain.Category) (*domain.Category, error)
	// Get(int) (*domain.Category, error)
}
