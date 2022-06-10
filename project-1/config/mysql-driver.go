package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	dbConnection := os.Getenv("DB_CONNECTIONS") // koneksi ke database
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Welcome")
		fmt.Println("================")
	}
	return db
}
