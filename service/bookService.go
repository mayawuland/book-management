package service

import (
	"book-management/database"
	"book-management/model"
	"database/sql"
	"errors"
)

// AddBookService inserts a new book into the database
func AddBookService(newBook model.Book) error {
	if newBook.Title == "" {
		return errors.New("book title cannot be empty")
	}

	insertQuery := "INSERT INTO books (title, author, year) VALUES (?, ?, ?)"
	_, err := database.DB.Exec(insertQuery, newBook.Title, newBook.Author, newBook.Year)
	return err
}

// GetAllBooksService retrieves all books from the database
func GetAllBooksService() ([]model.Book, error) {
	rows, err := database.DB.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allBooks []model.Book
	for rows.Next() {
		var book model.Book
		if scanError := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); scanError != nil {
			return nil, scanError
		}
		allBooks = append(allBooks, book)
	}

	return allBooks, nil
}

// GetBookByIdService retrieves specific book from the database by id
func GetBookByIdService(bookID int)(*model.Book, error) {
	query := "SELECT id, title, author, year FROM books WHERE id=?"

	var book model.Book
	err := database.DB.QueryRow(query, bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		
		return nil, err
	}

	return &book, nil
}

// UpdateBookService updates an existing book record by ID
func UpdateBookService(updatedBook model.Book) error {
	updateQuery := "UPDATE books SET title=?, author=?, year=? WHERE id=?"
	updateResult, err := database.DB.Exec(updateQuery, updatedBook.Title, updatedBook.Author, updatedBook.Year, updatedBook.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := updateResult.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

// DeleteBookService removes a book from the database by ID
func DeleteBookService(bookID int) error {
	deleteQuery := "DELETE FROM books WHERE id=?"
	deleteResult, err := database.DB.Exec(deleteQuery, bookID)
	if err != nil {
		return err
	}

	rowsAffected, _ := deleteResult.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

// SearchBookByTitleService retrieves specific book from the database by title
func SearchBookByTitleService(bookTitle string)(*model.Book, error) {
	query := "SELECT id, title, author, year FROM books WHERE title=?"

	var book model.Book
	err := database.DB.QueryRow(query, bookTitle).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		
		return nil, err
	}

	return &book, nil
}