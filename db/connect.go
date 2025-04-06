package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	dsn := "postgres://root:admin@localhost:5432/ssk_matrimony?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL DB")
	return db
}
