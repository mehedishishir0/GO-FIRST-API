package product

import "ecommerce/domain"

func (s *service) Update(p domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(p)
}
