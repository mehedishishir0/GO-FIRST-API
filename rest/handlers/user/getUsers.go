package user

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.List()

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "faild to fetch users")
		return
	}
	
	util.SendData(w, users, http.StatusOK)

}
