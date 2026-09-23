package user

import "ecommerce/domain"

func (svc *service) Create(user domain.User) (*domain.User, error) {
	
	return svc.userRepo.Create(user)

	
}
