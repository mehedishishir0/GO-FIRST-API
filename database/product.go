package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1

	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if productID == product.ID {
			return &product
		}
	}
	return nil

}

func Update(product Product) {

	for i, pd := range productList {
		if pd.ID == product.ID {
			productList[i] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for i, product := range productList {

		if product.ID != product.ID {
			tempList[i] = product
		}
	}
	productList = tempList
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
