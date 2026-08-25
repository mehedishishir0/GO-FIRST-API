package product

type service struct {
	productRepo ProductRepo
}

func NewService(pr ProductRepo) Service {
	return &service{
		productRepo: pr,
	}
}
