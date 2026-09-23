package repo

import (
	"ecommerce/blog"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type BlogRepo interface {
	blog.BlogRepo //embedding
}

type blogRepo struct {
	db *sqlx.DB
}

func NewBlogRepo(db *sqlx.DB) BlogRepo {
	return &blogRepo{db: db}
}

func (r *blogRepo) Create(b domain.Blog) (*domain.Blog, error) {
	query := `
		INSERT INTO blogs (title, content, author)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	row := r.db.QueryRow(query, b.Title, b.Content, b.Author)
	err := row.Scan(&b.ID)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
