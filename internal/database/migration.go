package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL, migrationsPath string) error {
	log.Printf("Starting migrations. DB URL: %s, Migrations Path: %s", dbURL, migrationsPath)

	// Verificar que el directorio existe
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", migrationsPath)
	}

	// Listar archivos en el directorio de migraciones
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}
	for _, file := range files {
		log.Printf("Found migration file: %s", file.Name())
	}

	// Construir la URL del source para migrate
	sourceURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))
	log.Printf("Using source URL: %s", sourceURL)

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get current migration version: %w", err)
	}
	log.Printf("Current migration version: %d, Dirty: %v", version, dirty)

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}
