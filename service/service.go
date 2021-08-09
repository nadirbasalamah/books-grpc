package service

import (
	"github.com/nadirbasalamah/books-grpc/model"
	"github.com/nadirbasalamah/books-grpc/repository"
)

// AddBook untuk menambahkan data buku
func AddBook(bookData model.Book) model.Book {
	return repository.AddBook(bookData)
}

// GetBook untuk mendapatkan data buku berdasarkan id
func GetBook(bookId string) (int, model.Book) {
	return repository.GetBook(bookId)
}

// GetBooks untuk mendapatkan seluruh data buku
func GetBooks() []model.Book {
	return repository.GetBooks()
}

// UpdateBook untuk mengedit data buku
func UpdateBook(bookData model.Book, id string) model.Book {
	return repository.UpdateBook(bookData, id)
}

// DeleteBook untuk menghapus data buku
func DeleteBook(id string) bool {
	return repository.DeleteBook(id)
}
