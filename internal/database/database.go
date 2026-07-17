package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@/users")
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := createUsersTableIfNotExists(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createUsersTableIfNotExists(db *sql.DB) error {
	query := `	
	CREATE TABLE IF NOT EXISTS users (
		id CHAR(36) PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		biography TEXT NOT NULL);`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("creating users table: %w", err)
	}

	return nil
}
