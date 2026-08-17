package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create()
	Get()
	List()
	Delete()
	Update()
}

type productRepo struct {
	productList []Product
}

func NewProductRepo() ProductRepo {
	return productRepo{}
}

func (r productRepo) Create()
func (r productRepo) Get()
func (r productRepo) List()
func (r productRepo) Delete()
func (r productRepo) Update()
