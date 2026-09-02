package category

type service struct{
	categoryRepo CategoryRepo

}

func NewService(ct CategoryRepo) Service{
	return &service{
		categoryRepo: ct,
	}
}