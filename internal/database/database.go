package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Balajivarma28092006/GitReview/internal/models/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	pool *pgxpool.Pool
}

// Loads config and creates a new Service and returns that New Service 
func New() (*Service, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName, cfg.DBSslMode)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create a connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping the database: %w", err)
	}	
	log.Println("Connceted to postgresSQL successfully.")
	service := &Service{
		pool: pool,
	}
	return service, nil
}

// for shutting down the connection pool
func (s *Service) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

// accessing the service pool that we can use them in our service modules
func (s *Service) Pool() *pgxpool.Pool {
	return s.pool
}
