package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

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

	return &Database{DB: db}, nil
}
