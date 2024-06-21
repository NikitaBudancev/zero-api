package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	dbURL := fmt.Sprint("postgres://postgres:postgres@localhost:5432/mynewsapp?sslmode=disable")

	m, err := migrate.New(
		"file://internal/migrations",
		dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
