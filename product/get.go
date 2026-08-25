package product

import "ecommerce/domain"

func (s *service) Get(id int) (*domain.Product, error) {
	return s.productRepo.Get(id)
}
