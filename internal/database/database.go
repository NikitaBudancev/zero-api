package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"zero_api/internal/config"
)

var DB *reform.DB

func ConnectDatabase(cfg *config.DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))
	DB = db
	fmt.Println("Database connection successful")
}

func CountRecords(query string) (int, error) {
	var totalRecords int
	err := DB.QueryRow(query).Scan(&totalRecords)
	return totalRecords, err
}
