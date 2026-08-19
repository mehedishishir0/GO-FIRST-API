package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&reqLogin)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	usr, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	util.CreateJWT("secret", util.Payload{
		Sub:          usr.ID,
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		Email:        usr.Email,
		IsShopeOwner: usr.IsShopeOwner,
	})

	accessToken, err := util.CreateJWT(h.cnf.JWTSecret, util.Payload{
		Sub:          usr.ID,
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		Email:        usr.Email,
		IsShopeOwner: usr.IsShopeOwner,
	})

	if err != nil {
		http.Error(w, "Failed to create JWT", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, 201)
}
