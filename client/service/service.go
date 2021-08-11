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

	// memasukkan id buku ke dalam variabel bookId
	bookId = res.Data.Id

	// menampilkan response dari server
	fmt.Printf("Book added: %v\n", res)
}

// AddBatchBook untuk menambahkan sekumpulan data buku
func (s *Service) AddBatchBook() {
	// membuat beberapa request
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

	// menambahkan beberapa data buku
	stream, err := s.Client.AddBatchBook(context.Background())

	// jika terjadi error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// untuk setiap request
	// kirim melalui stream
	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
	}

	// tutup stream
	// karena sudah digunakan
	res, err := stream.CloseAndRecv()

	// jiak terjadi error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan hasil penambahan data buku
	fmt.Printf("Add batch book result: %v\n", res.Message)
}

// GetBooks untuk mendapatkan seluruh data buku
func (s *Service) GetBooks() {
	// mendapatkan seluruh data buku
	stream, err := s.Client.GetBooks(context.Background(), &bookpb.GetBooksRequest{})

	// jika terjadi error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan seluruh data buku
	fmt.Println("All Books")
	// untuk setiap response dari stream
	for {
		// terima response
		res, err := stream.Recv()

		// jika tidak ada response lagi
		// hentikan eksekusi for loop
		if err == io.EOF {
			break
		}

		// jika terdapat error, tampilkan error
		if err != nil {
			log.Fatalf("Error when streaming: %v", err)
		}

		// menampilkan data buku
		fmt.Println(res.Data)
	}

}

// GetBook untuk mendapatkan data buku
func (s *Service) GetBook() {
	// mendapatkan data buku berdasarkan id
	res, err := s.Client.GetBook(context.Background(), &bookpb.GetBookRequest{
		Id: bookId,
	})

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan data buku
	fmt.Printf("Book data: %v\n", res)
}

// UpdateBook untuk mengubah data buku
func (s *Service) UpdateBook() {
	// membuat request
	var request bookpb.UpdateBookRequest = bookpb.UpdateBookRequest{
		Book: &bookpb.Book{
			Id:     bookId,
			Title:  "updated title",
			Author: "updated author",
			IsRead: true,
		},
	}

	// mengubah data buku
	res, err := s.Client.UpdateBook(context.Background(), &request)

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan data buku yang telah diubah
	fmt.Printf("Updated book: %v\n", res)
}

// DeleteBook untuk menghapus data buku
func (s *Service) DeleteBook() {
	// menghapus data buku berdasarkan id
	res, err := s.Client.DeleteBook(context.Background(), &bookpb.DeleteBookRequest{
		BookId: bookId,
	})

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// menampilkan hasil penambahan data buku
	fmt.Printf("Deleted book: %v\n", res)
}
