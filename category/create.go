package category

import "ecommerce/domain"

func (s *service) Create(p domain.Category) (*domain.Category, error) {
	return s.categoryRepo.Create(p)
}
