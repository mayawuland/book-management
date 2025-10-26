package controller

import (
	"book-management/model"
	"book-management/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// AddBook handles POST requests to add a new book
func AddBook(responseWriter http.ResponseWriter, request *http.Request) {
	var newBook model.Book
	decodeError := json.NewDecoder(request.Body).Decode(&newBook)
	if decodeError != nil {
		http.Error(responseWriter, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	err := service.AddBookService(newBook)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Book added successfully!"})
}

// GetAllBooks handles GET requests to retrieve all books
func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	allBooks, err := service.GetAllBooksService()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(allBooks)
}

// GetBookById handles GET requests to get a specific book by id (?id=)
func GetBookById(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(responseWriter, "Missing book ID in query parameter", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(responseWriter, "Invalid book ID format", http.StatusBadRequest)
		return
	}

	book, err := service.GetBookByIdService(bookID)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(responseWriter, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(book)
}

// UpdateBook handles PUT requests to update an existing book (?id=)
func UpdateBook(responseWriter http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(responseWriter, "Missing book ID in query parameter", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(responseWriter, "Invalid book ID format", http.StatusBadRequest)
		return
	}

	var updatedBook model.Book
	decodeError := json.NewDecoder(request.Body).Decode(&updatedBook)
	if decodeError != nil {
		http.Error(responseWriter, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	updatedBook.ID = bookID

	err = service.UpdateBookService(updatedBook)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(responseWriter, err.Error(), http.StatusNotFound)
		} else {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Book updated successfully!"})
}

// DeleteBook handles DELETE requests to remove a book by query ?id=
func DeleteBook(responseWriter http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(responseWriter, "Missing book ID in query parameter", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(responseWriter, "Invalid book ID format", http.StatusBadRequest)
		return
	}

	err = service.DeleteBookService(bookID)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(responseWriter, err.Error(), http.StatusNotFound)
		} else {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Book deleted successfully!"})
}

// SearchBookByTitle handles GET requests to get a specific book by Title (?title=)
func SearchBookByTitle(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	titleStr := request.URL.Query().Get("title")
	if titleStr == "" {
		http.Error(responseWriter, "Missing book ID in query parameter", http.StatusBadRequest)
		return
	}

	book, err := service.SearchBookByTitleService(titleStr)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(responseWriter, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(book)
}