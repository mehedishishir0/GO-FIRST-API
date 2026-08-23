package user

import (
	"ecommerce/config"
)

type Handler struct {
	cnf *config.Config
	svc Service
}

func NewHandler(srv Service, cnf *config.Config) *Handler {
	return &Handler{
		cnf: cnf,
		svc: srv,
	}
}
