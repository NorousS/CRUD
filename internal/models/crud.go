package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=todouser password=todopass dbname=todoapp sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr) // присваиваем в глобальную DB
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
