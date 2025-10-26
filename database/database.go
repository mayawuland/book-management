package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:MayaWulandari89@tcp(127.0.0.1:3306)/book_management")

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Database connected successfully!")
}
