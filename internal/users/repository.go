package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrDuplicateEmail    = errors.New("a user with this email already exists")
	ErrDuplicateUsername = errors.New("a user with this username already exists")
	ErrDatabaseInternal  = errors.New("an unexpected database error occurred")
)

type UserRepository interface {
	Insert(ctx context.Context, user *User) (*User, error)
}

type PostGresRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *PostGresRepository {
	return &PostGresRepository{
		pool: pool,
	}
}

func (repo *PostGresRepository) Insert(ctx context.Context, user *User) (*User, error) {
	query := `INSERT INTO users (id, username, email, created_at) VALUES ($1, $2, $3, $4)`

	_, err := repo.pool.Exec(ctx, query, user.ID, user.UserName, user.Email, user.CreatedAt)
	if err != nil {
		fmt.Println("ACTUAL DB ERROR:", err)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// 2305 represents unique conflicts error
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {
				case "users_username_key":
					return nil, ErrDuplicateUsername
				case "users_email_key":
					return nil, ErrDuplicateEmail
				}
			}
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseInternal, err)
	}

	createdUser := &User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return createdUser, nil
}
