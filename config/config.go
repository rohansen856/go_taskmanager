// config/config.go
package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// LoadEnv loads environment variables from a .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, using system environment variables")
	}
}

// InitDB initializes the database connection.
func InitDB() (*sql.DB, error) {
	dbURL := GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/postgres?sslmode=disable")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

// GetEnv retrieves environment variables or returns a default value.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
