package product

import "ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64,error)
	Delete(int) error
	Update(domain.Product) (*domain.Product, error)
}

