package service

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/nadirbasalamah/books-grpc/book/bookpb"
)

// Service untuk menggunakan
// fungsi dari server gRPC
type Service struct {
	Client bookpb.BookServiceClient
}

// membuat variabel untuk menyimpan id buku
var bookId string

// AddBook untuk menambahkan data buku
func (s *Service) AddBook() {
	// membuat request
	var request bookpb.AddBookRequest = bookpb.AddBookRequest{
		Book: &bookpb.Book{
			Title:  "my book",
			Author: "grpc client",
			IsRead: false,
		},
	}

	// menambahkan data buku
	res, err := s.Client.AddBook(context.Background(), &request)

	// jika terdapat error
	// tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan response dari server
	fmt.Printf("Book added: %v\n", res)
}

// AddBatchBook untuk menambahkan sekumpulan data buku
func (s *Service) AddBatchBook() {
	var requests []*bookpb.AddBatchBookRequest = []*bookpb.AddBatchBookRequest{
		{
			Book: &bookpb.Book{
				Title:  "title one",
				Author: "author one",
				IsRead: false,
			},
		},
		{
			Book: &bookpb.Book{
				Title:  "title two",
				Author: "author two",
				IsRead: false,
			},
		},
	}

	stream, err := s.Client.AddBatchBook(context.Background())

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Printf("Add batch book result: %v\n", res.Message)
}

// GetBooks untuk mendapatkan seluruh data buku
func (s *Service) GetBooks() {
	stream, err := s.Client.GetBooks(context.Background(), &bookpb.GetBooksRequest{})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Println("All Books")
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when streaming: %v", err)
		}

		fmt.Println(res.Data)

		// memasukkan id buku ke dalam variabel bookId
		bookId = res.Data.Id
	}

}

// GetBook untuk mendapatkan data buku
func (s *Service) GetBook() {
	res, err := s.Client.GetBook(context.Background(), &bookpb.GetBookRequest{
		Id: bookId,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Printf("Book data: %v\n", res)
}

// UpdateBook untuk mengubah data buku
func (s *Service) UpdateBook() {
	var request bookpb.UpdateBookRequest = bookpb.UpdateBookRequest{
		Book: &bookpb.Book{
			Id:     bookId,
			Title:  "updated title",
			Author: "updated author",
			IsRead: true,
		},
	}

	res, err := s.Client.UpdateBook(context.Background(), &request)

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Printf("Updated book: %v", res)
}

// DeleteBook untuk menghapus data buku
func (s *Service) DeleteBook() {
	res, err := s.Client.DeleteBook(context.Background(), &bookpb.DeleteBookRequest{
		BookId: bookId,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Printf("Deleted book: %v", res)
}
