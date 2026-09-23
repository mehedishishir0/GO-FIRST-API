package blog

import "ecommerce/domain"

type Service interface {
	Create(b domain.Blog) (*domain.Blog, error)
}
