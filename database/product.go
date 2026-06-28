package database

var ProductList []Product


type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
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

	ProductList = append(ProductList, prd1, prd2)

}
