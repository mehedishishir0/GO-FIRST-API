package blog

type service struct {
 blogRepo BlogRepo
}

func NewSercvice(br BlogRepo) Service {
	return &service{
		blogRepo: br,
	}
}