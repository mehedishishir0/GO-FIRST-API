package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() []*Product
	Delete(productID int) ( error)
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1

	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if productID == product.ID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() []*Product {
	return r.productList

}
func (r *productRepo) Delete(productID int) ( error) {
	var tempList []*Product

	for _, product := range r.productList {

		if product.ID != productID {
			tempList = append(tempList, product)
		}
	}
	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {

	for i, pd := range r.productList {
		if pd.ID == product.ID {
			r.productList[i] = &product
		}
	}
	return &product, nil
}

func generateProduct(r *productRepo) {

	prd1 := &Product{
		ID:          1,
		Title:       "Product 1",
		Price:       100.0,
		Description: "Description 1",
		ImageURL:    "https://example.com/image1.jpg",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Product 2",
		Price:       200.0,
		Description: "Description 2",
		ImageURL:    "https://example.com/image2.jpg",
	}

	r.productList = append(r.productList, prd1, prd2)

}
