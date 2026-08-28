package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := ` INSERT INTO products (
			title,
			price,
			img_url,
			description
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
			RETURNING id 
	`

	row := r.db.QueryRow(query, p.Title, p.Price, p.ImageURL, p.Description)

	fmt.Println(row)

	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := "SELECT id , title, description, price, img_url FROM products WHERE id = $1 "

	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var prdlist []*domain.Product
	fmt.Println(page, limit)

	offset := ((page - 1) * limit) + 1

	query := "SELECT id , title, description, price, img_url FROM products LIMIT $1 OFFSET $2 "

	err := r.db.Select(&prdlist, query, limit, offset)


	if err != nil {
		return nil, err
	}

	return prdlist, err

}

func (r *productRepo) Count() (int64, error) {
	var count int64

	query := "SELECT COUNT(*) FROM products"

	err := r.db.QueryRow(query).Scan(&count)


	if err != nil {
		return 0, err
	}

	return count, err

}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {

	query := "UPDATE  products SET title=$1, description=$2, price=$3, img_url=$4 WHERE id = $5  "

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageURL, p.ID)
	err := row.Err()

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Delete(id int) error {

	query := "DELETE FROM products WHERE id = $1 "

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
