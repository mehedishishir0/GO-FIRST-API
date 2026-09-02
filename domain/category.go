package domain

// model or entity 
type Category struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"name" db:"title"`
	Description string  `json:"description" db:"description"`
	ImageURL    string  `json:"imageUrl" db:"img_url"`
}