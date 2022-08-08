package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func CreateConnexion() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	userName := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", userName, pass, dbName))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
