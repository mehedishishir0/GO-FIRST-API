package user

import "ecommerce/domain"

func (svc *service) Find(email string, pass string) (*domain.User, error) {
	return svc.userRepo.Find(email, pass)
}
