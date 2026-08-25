package product

import (
	"ecommerce/domain"
)

func (s *service) List(page, limit int64) ([]*domain.Product, error) {
	
	return s.productRepo.List(page, limit)
}
