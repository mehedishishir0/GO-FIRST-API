package category

import "ecommerce/domain"

type Service interface {
	Create(domain.Category) (*domain.Category, error)
	Get(int) (*domain.Category, error)
}
