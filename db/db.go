package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB unreachable:", err)
	}
}
