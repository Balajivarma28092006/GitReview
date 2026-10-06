package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBSslMode  string `env:"DB_SSLMODE"`
	Port       int    `env:"PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to parse env: %w", err)
	}
	return &cfg, nil
}
