package domain

// model or entity 
type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"name" db:"title"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImageURL    string  `json:"imageUrl" db:"img_url"`
}