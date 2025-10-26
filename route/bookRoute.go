package route

import (
	"book-management/controller"
	"net/http"
)

// SetupBookRoutes registers all book-related routes
func SetupBookRoutes() {
	http.HandleFunc("/books/add", controller.AddBook)
	http.HandleFunc("/books", controller.GetAllBooks)
	http.HandleFunc("/books/getbyid", controller.GetBookById)
	http.HandleFunc("/books/update", controller.UpdateBook)
	http.HandleFunc("/books/delete", controller.DeleteBook)
	http.HandleFunc("/books/searchbytitle", controller.SearchBookByTitle)
}
