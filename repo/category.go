package repo

import (
	"ecommerce/category"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo interface {
	category.CategoryRepo
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(p domain.Category) (*domain.Category, error) {
	query := `
	INSERT INTO categories (
		title,
		description,
		img_url
	)
	VALUES (
		$1,
		$2,
		$3
	)
	RETURNING id
`

	row := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImageURL,
	)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// func (r *categoryRepo) Get(int) (*domain.Category, error) {
// 	var prd domain.Product

// 	query := "SELECT id , title, description, price, img_url FROM products WHERE id = $1 "

// 	err := r.db.Get(&prd, query, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &prd, nil
// }
