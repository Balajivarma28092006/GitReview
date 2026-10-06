package posts

import "github.com/jackc/pgx/v5/pgxpool"

type PostRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		pool: pool,
	}
}
