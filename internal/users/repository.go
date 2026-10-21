package users

import (
	"context"
	"errors"
	"fmt"
	"uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrDuplicateEmail    = errors.New("a user with this email already exists")
	ErrDuplicateUsername = errors.New("a user with this username already exists")
	ErrDatabaseInternal  = errors.New("an unexpected database error occurred")
	ErrUserNotFound      = errors.New("user not found")
)

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type PostGresRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *PostGresRepository {
	return &PostGresRepository{
		pool: pool,
	}
}

func (repo *PostGresRepository) Insert(ctx context.Context, user *User) error {
	query := `INSERT INTO users (id, username, email, created_at) VALUES ($1, $2, $3, $4)`

	_, err := repo.pool.Exec(ctx, query, user.ID, user.UserName, user.Email, user.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// 2305 represents unique conflicts error
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {
				case "users_username_key":
					return ErrDuplicateUsername
				case "users_email_key":
					return ErrDuplicateEmail
				}
			}
		}
		return fmt.Errorf("%w: %v", ErrDatabaseInternal, err)
	}
	return nil
}

const userColumns = `id, username, email, password_hash, created_at`

func (repo *PostGresRepository) getOne(ctx context.Context, query string, arg any) (*User, error) {
	var u User
	err := repo.pool.QueryRow(ctx, query, arg).Scan(&u.ID, &u.UserName, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return &u, nil
}

func (repo *PostGresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	return repo.getOne(ctx, `SELECT `+userColumns+` FROM users WHERE email = $1`, email)
}

func (repo *PostGresRepository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return repo.getOne(ctx, `SELECT `+userColumns+` FROM users WHERE id = $1`, id)
}
