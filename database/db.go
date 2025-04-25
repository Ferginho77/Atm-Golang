package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/goland"

	DB, err = sql.Open("mysql", dsn) 
	if err != nil {
		panic(fmt.Sprintf("Error opening DB connection: %v", err))
	}

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error pinging DB: %v", err))
	}

	fmt.Println("Database connected!")
}
