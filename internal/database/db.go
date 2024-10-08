package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

var (
	instance *Database
	once     sync.Once
)

func Init(dataSourceName string) (*Database, error) {
	var db *sql.DB
	var err error

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			return nil, fmt.Errorf("error opening database: %w", err)
		}

		err = db.Ping()
		if err == nil {
			break
		}

		fmt.Printf("Failed to connect to database. Retrying in 5 seconds... (Attempt %d/%d)\n", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	instance = &Database{DB: db}
	return instance, nil
}

func GetDB() *Database {
	if instance == nil {
		panic("Database not initialized. Call Init first.")
	}
	return instance
}
