package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// Silakan ubah parameter sesuai konfigurasi lokal PostgreSQL-mu
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "exol930103")
	dbname := getEnv("DB_NAME", "inventorydb")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("gagal ping database: %w", err)
	}

	fmt.Println("Berhasil terkoneksi ke database PostgreSQL 🎉")
	return nil
}

// Helper untuk membaca environment variable
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
