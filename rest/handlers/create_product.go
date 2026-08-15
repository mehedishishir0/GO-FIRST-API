package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProuct(w http.ResponseWriter, r *http.Request) {
	// parse jwt

	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	headerArry := strings.Split(header, " ")
	if len(headerArry) != 2 {
		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
		return
	}

	fmt.Println(headerArry)

	accessToken := headerArry[1]

	tokenParts := strings.Split(accessToken, ".")
	if len(tokenParts) != 3 {
		http.Error(w, "Invalid JWT format", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	jwtSignature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload

	cnf := config.GetConfig()

	byteArrSecret := []byte(cnf.JWTSecret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	expectedSignature := h.Sum(nil)
	signatureBase64 := base64UrlEncode(expectedSignature)

	if signatureBase64 != jwtSignature {
		http.Error(w, "Invalid JWT signature", http.StatusUnauthorized)
		return
	}
	

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)
}

func base64UrlEncode(data []byte) string {

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
