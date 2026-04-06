package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=postgres password=1109 dbname=inventory sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
	}
	DB = db
	fmt.Println("Подключение к БД успешно!")
}
