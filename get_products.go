package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Please hit the get method", 400)
		return
	}

	sendData(w, productList, 200)

}