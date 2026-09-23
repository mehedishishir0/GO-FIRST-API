package user

import "ecommerce/domain"

func (svc *service) List() ([]*domain.User, error) {
	return svc.userRepo.List()
}
