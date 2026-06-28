package main

import (
	"fmt"
	"net/http"
)



var productList []Product

func main() {

	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(createProuct))

	globalRouter := globalRouter(mux)

	fmt.Println("Server is running on port 3001")
	err := http.ListenAndServe(":3001", globalRouter) // Start the server on port 3001

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
