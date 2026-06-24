package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	handelCors(w)
	handelOption(w, r)

	if r.Method != "GET" {
		http.Error(w, "Please hit the get method", 400)
		return
	}

	sendData(w, productList, 200)

}

func createProuct(w http.ResponseWriter, r *http.Request) {
	handelCors(w)
	handelOption(w, r)

	if r.Method == "POST" {
		http.Error(w, "Please hit the post method", 400)
	}
	var newProduct Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func handelCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handelOption(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}

}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)

}

func main() {

	mux := http.NewServeMux() // router

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler)) // route

	mux.HandleFunc("GET /about", http.HandlerFunc(aboutHandler))
	mux.HandleFunc("GET /products", http.HandlerFunc(getProducts))
	mux.HandleFunc("POST /create-product", http.HandlerFunc(createProuct))
	fmt.Println("Server is running on port 3001")

	err := http.ListenAndServe(":3001", mux) // Start the server on port 3001

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:       100.0,
		Description: "Description 1",
		ImageURL:    "https://example.com/image1.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:       200.0,
		Description: "Description 2",
		ImageURL:    "https://example.com/image2.jpg",
	}

	productList = append(productList, prd1, prd2)
}
