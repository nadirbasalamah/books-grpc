package repository

import (
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
	for index, v := range storage {
		if v.Id == bookId {
			return index, v
		}
	}
	return 0, model.Book{}
}

// GetBooks untuk mendapatkan seluruh data buku
func GetBooks() []model.Book {
	return storage
}

// UpdateBook untuk mengedit data buku
func UpdateBook(bookData model.Book, id string) model.Book {
	index, book := GetBook(id)

	book.Title = bookData.Title
	book.Author = bookData.Author
	book.IsRead = bookData.IsRead

	storage[index] = book

	return book
}

// DeleteBook untuk menghapus data buku
func DeleteBook(id string) bool {
	var afterDeleted []model.Book = []model.Book{}
	for _, v := range storage {
		if id != v.Id {
			afterDeleted = append(afterDeleted, v)
		}
	}

	storage = afterDeleted
	return true
}
