package product

import "ecommerce/domain"

func (s *service) Create(p domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(p)
}
