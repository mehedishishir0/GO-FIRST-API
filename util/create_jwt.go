package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub          int `json:"sub"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	IsShopeOwner bool   `json:"isShopOwner"`
}

func CreateJWT( secret string, payload Payload,) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)

	if err != nil {
		// Handle error
		return "", err
	}

	headerBase64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	dataBase64 := base64UrlEncode(byteArrData)

	byteArrSecret := []byte(secret)
	message := headerBase64 + "." + dataBase64

	byteArrMessage := []byte(message)

	signature := hmac.New(sha256.New, byteArrSecret)
	signature.Write(byteArrMessage)
	byteArrSignature := signature.Sum(nil)
	signatureBase64 := base64UrlEncode(byteArrSignature)

	return headerBase64 + "." + dataBase64 + "." + signatureBase64, nil
}

func base64UrlEncode(data []byte) string {

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
