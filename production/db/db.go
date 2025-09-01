package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sql.DB

func DB_connection() {
	godotenv.Load(".env")

	pg_url := os.Getenv("DATABASE_URL")
	if pg_url == "" {
		log.Fatal("DATABASE_URL not found")
	}

	var err error
	Db, err = sql.Open("postgres", pg_url)
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}

	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(100)

	err = Db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}
	fmt.Println("✅ Connected to PostgreSQL")
}
