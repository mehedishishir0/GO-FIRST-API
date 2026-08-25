package product

func (s *service) Delete(id int) error {
	return s.productRepo.Delete(id)
}
