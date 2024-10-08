package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
}

func Load() (*Config, error) {

	cfg := &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://user:password@db:5432/deuna_challenge?sslmode=disable"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}

	log.Printf("Loaded config: %+v", cfg)
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
