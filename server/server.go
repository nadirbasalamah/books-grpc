package server

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/nadirbasalamah/books-grpc/book/bookpb"
	"github.com/nadirbasalamah/books-grpc/model"
	"github.com/nadirbasalamah/books-grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server merupakan server gRPC
type Server struct {
}

// fungsi untuk memetakan data dari book model ke bookpb Book
func mapToBookpb(book model.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:     book.Id,
		Author: book.Author,
		Title:  book.Title,
		IsRead: book.IsRead,
	}
}

// GetBook untuk mendapatkan data buku berdasarkan id
func (*Server) GetBook(ctx context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {

	// mengambil inputan id dari request
	var bookId string = req.GetId()

	// memanggil GetBook untuk mendapatkan data buku berdasarkan id
	_, result := service.GetBook(bookId)

	// jika tidak ditemukan, beri pesan error
	if result.Id != bookId {
		return &bookpb.GetBookResponse{Status: false, Data: nil}, errors.New("Data not found!")
	}

	// mengolah hasil dari GetBook
	// untuk data di objek response
	var bookData *bookpb.Book = &bookpb.Book{
		Id:     result.Id,
		Author: result.Author,
		Title:  result.Title,
		IsRead: result.IsRead,
	}

	// memberikan response berupa data buku
	return &bookpb.GetBookResponse{
		Status: true,
		Data:   bookData,
	}, nil

}

// AddBook untuk menambahkan data buku
func (*Server) AddBook(ctx context.Context, req *bookpb.AddBookRequest) (*bookpb.AddBookResponse, error) {

	// menerima request
	var bookRequest *bookpb.Book = req.GetBook()

	// membuat variabel bookData
	// untuk dimasukkan ke storage
	var bookData model.Book = model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	// memanggil AddBook untuk menambahkan data buku
	var insertedBook model.Book = service.AddBook(bookData)

	// mengembalikan response berupa data buku
	// yang telah ditambahkan
	return &bookpb.AddBookResponse{
		Status: true,
		Data:   mapToBookpb(insertedBook),
	}, nil

}

// UpdateBook untuk mengubah data buku
func (*Server) UpdateBook(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {

	// menerima request
	var bookRequest *bookpb.Book = req.GetBook()

	// membuat variabel bookData
	// untuk mengubah data buku
	var bookData model.Book = model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	// memanggil UpdateBook untuk mengubah data buku berdasarkan id
	var updatedBook model.Book = service.UpdateBook(bookData, bookData.Id)

	// mengembalikan response berupa
	// data buku yang telah diubah
	return &bookpb.UpdateBookResponse{
		Status: true,
		Data:   mapToBookpb(updatedBook),
	}, nil
}

// DeleteBook untuk menghapus data buku
func (*Server) DeleteBook(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {

	// mengambil inputan id dari request
	var bookId string = req.GetBookId()

	// memanggil DeleteBook untuk menghapus data buku berdasarkan id
	var result bool = service.DeleteBook(bookId)

	// mengembalikan response
	return &bookpb.DeleteBookResponse{
		Status: result,
	}, nil
}

// GetBooks untuk mendapatkan seluruh data buku
func (*Server) GetBooks(req *bookpb.GetBooksRequest, stream bookpb.BookService_GetBooksServer) error {

	// memanggil GetBooks untuk mendapatkan seluruh data buku
	var books []model.Book = service.GetBooks()

	// melakukan iterasi pada setiap data buku
	for _, book := range books {
		// setiap data buku
		// dikirimkan melalui stream
		stream.Send(&bookpb.GetBooksResponse{
			Status: true,
			Data:   mapToBookpb(book),
		})
	}

	// karena tidak terdapat error
	// kembalian nilai nil / kosong
	return nil
}

// AddBatchBook untuk menambahkan sekumpulan data buku
func (*Server) AddBatchBook(stream bookpb.BookService_AddBatchBookServer) error {
	// untuk setiap request
	for {
		// menerima request
		req, err := stream.Recv()

		// jika request tidak ada
		// tutup stream
		if err == io.EOF {
			return stream.SendAndClose(&bookpb.AddBatchBookResponse{
				Status:  true,
				Message: "All book data inserted successfully!",
			})
		}

		// jika terdapat error di server
		// tampilkan pesan error
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error, insert batch failed: %v", err),
			)
		}

		// mendapatkan request data buku
		var bookData *bookpb.Book = req.GetBook()

		// membuat variabel book
		// untuk dimasukkan ke storage
		var book model.Book = model.Book{
			Id:     bookData.GetId(),
			Title:  bookData.GetTitle(),
			Author: bookData.GetAuthor(),
			IsRead: bookData.GetIsRead(),
		}

		// tambahkan data buku
		service.AddBook(book)

	}
}
