package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&reqLogin)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	
	usr := database.Find(reqLogin.Email, reqLogin.Password)

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

	cnf := config.GetConfig()
	
   accessToken, err := util.CreateJWT(cnf.JWTSecret, util.Payload{
		Sub:      usr.ID,
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
