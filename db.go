package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func init_db() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))

}
