package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		byteArrSecret := []byte(m.cnf.JWTSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		expectedSignature := h.Sum(nil)
		signatureBase64 := base64UrlEncode(expectedSignature)

		if signatureBase64 != jwtSignature {
			http.Error(w, "Invalid JWT signature", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})

}

func base64UrlEncode(data []byte) string {

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
