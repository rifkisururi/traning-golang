package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string dari .env
	connStr := "host=ep-lucky-mud-a1sn3mmx-pooler.ap-southeast-1.aws.neon.tech port=5432 user=neondb_owner password=npg_IuEdFJ37pvWj dbname=neondb sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test koneksi
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database. Running migration...")

	// Baca file migration
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run migrate.go <migration-file>")
	}

	migrationFile := os.Args[1]
	content, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	// Jalankan migration
	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	fmt.Println("Migration completed successfully!")
}
