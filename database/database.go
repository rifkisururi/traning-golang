package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDatabase membuat koneksi ke PostgreSQL
func ConnectDatabase(connectionString string) error {
	var err error

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	// Test koneksi
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("gagal ping database: %w", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Printf("[database] Berhasil terkoneksi ke PostgreSQL")

	return nil
}

// CloseDatabase menutup koneksi database
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
