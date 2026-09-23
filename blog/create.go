package blog

import "ecommerce/domain"

func (s *service) Create(b domain.Blog) (*domain.Blog, error) {
	return s.blogRepo.Create(b)
}