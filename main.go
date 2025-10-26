package main

import (
	"book-management/database"
	"book-management/route"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Book Management API...")

	// Connect to database
	database.ConnectDB()

	// Setup routes
	route.SetupBookRoutes()

	// Run the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
