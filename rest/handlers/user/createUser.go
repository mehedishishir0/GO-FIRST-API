package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqUser struct {
	
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopeOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser ReqUser

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)

	if err != nil {
	 util.SendError(w, http.StatusBadGateway, "invalid req body")
		return
	}

	usr, err := h.userRepo.Create(repo.User{
		FirstName: newUser.FirstName,
		Email: newUser.Email,
		Password: newUser.Password,
		IsShopeOwner: newUser.IsShopeOwner,
		
	})
      if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	  }


	util.SendData(w,  usr, 201)
}
