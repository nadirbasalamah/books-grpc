package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/nadirbasalamah/books-grpc/database"
	"github.com/nadirbasalamah/books-grpc/model"
)

// membuat penyimpanan lokal dengan menggunakan slice
var storage []model.Book = []model.Book{}

// AddBook untuk menambahkan data buku
func AddBook(bookData model.Book) model.Book {
	var uuid string = uuid.New().String()

	_, err := database.DB.Query("INSERT INTO books (id, title, author, is_read) VALUES (?, ?, ?, ?)",
		uuid,
		bookData.Title,
		bookData.Author,
		bookData.IsRead,
	)

	if err != nil {
		log.Fatalf("Insert data failed: %v", err)
		return model.Book{}
	}

	return bookData

}

// GetBook untuk mendapatkan data buku berdasarkan id
func GetBook(bookId string) (int, model.Book) {
	var book model.Book = model.Book{}

	row, err := database.DB.Query("SELECT * FROM books WHERE id = ?", bookId)

	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return 0, model.Book{}
	}

	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&book.Id, &book.Title, &book.Author, &book.IsRead); err {
		case sql.ErrNoRows:
			log.Printf("Data not found: %v", err)
			return 0, model.Book{}
		case nil:
			log.Println(book)
		default:
			log.Printf("Data cannot be retrieved: %v", err)
			return 0, model.Book{}
		}
	}

	return 1, book

}

// GetBooks untuk mendapatkan seluruh data buku
func GetBooks() []model.Book {
	rows, err := database.DB.Query("SELECT * FROM books")

	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return []model.Book{}
	}

	defer rows.Close()

	var books []model.Book = []model.Book{}

	for rows.Next() {
		var book model.Book = model.Book{}

		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.IsRead)

		if err != nil {
			log.Printf("Data cannot be retrieved: %v", err)
			return []model.Book{}
		}

		books = append(books, book)
	}

	if len(books) == 0 {
		log.Println("Books data not found")
	}

	return books

}

// UpdateBook untuk mengedit data buku
func UpdateBook(bookData model.Book, id string) model.Book {
	_, err := database.DB.Query("UPDATE books SET title=?, author=?, is_read=? WHERE id=?",
		bookData.Title,
		bookData.Author,
		bookData.IsRead,
		bookData.Id,
	)

	if err != nil {
		log.Fatalf("Update data failed: %v", err)
		return model.Book{}
	}

	return bookData

}

// DeleteBook untuk menghapus data buku
func DeleteBook(id string) bool {
	_, err := database.DB.Query("DELETE FROM books WHERE id=?", id)
	if err != nil {
		log.Fatalf("Delete data failed: %v", err)
		return false
	}
	return true
}
