package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load .env file")
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("error opening the database!")
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln("cannot connect to database")
		return nil, err
	}

	log.Println("connected to postgres sql")
	return DB, nil
}
