package config

import (
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
}

func Load() (*Config, error) {
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgresql://user:password@localhost:5432/deuna_challenge?sslmode=disable"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
