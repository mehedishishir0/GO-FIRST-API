package product

func (h *service) Count() (int64, error) {
	return h.productRepo.Count()
}

