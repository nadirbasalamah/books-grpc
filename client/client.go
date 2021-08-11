package main

import (
	"fmt"
	"log"

	"github.com/nadirbasalamah/books-grpc/book/bookpb"
	"github.com/nadirbasalamah/books-grpc/client/service"
	"google.golang.org/grpc"
)

// membuat variabel untuk menyimpan service
var clientService service.Service

func main() {
	fmt.Println("Client of Book Service")
	// menyambungkan ke server gRPC
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the book service: %v", err)
	}

	// jika program selesai digunakan
	// tutup koneksi dengan server gRPC
	defer cc.Close()

	// membuat client gRPC
	var client bookpb.BookServiceClient = bookpb.NewBookServiceClient(cc)

	// membuat service
	clientService = service.Service{Client: client}

	// memanggil fungsi showOptions
	showOptions()
}

// showOptions menampilkan pilihan menu di konsol / terminal
func showOptions() {
	var done bool = false

	for {
		// menampilkan daftar menu
		fmt.Println("Choose the options")
		fmt.Println("1. Add New Book")
		fmt.Println("2. Add Many Books")
		fmt.Println("3. Get All Books")
		fmt.Println("4. Get Book By ID")
		fmt.Println("5. Update Book")
		fmt.Println("6. Delete Book")
		fmt.Println("7. Exit")

		var choice int32

		// mengambil inputan nomor menu
		fmt.Scanln(&choice)

		// menentukan operasi
		// berdasarkan inputan user
		switch choice {
		case 1:
			fmt.Println("Adding book..")
			// memanggil fungsi AddBook
			clientService.AddBook()
		case 2:
			fmt.Println("Adding many books..")
			// memanggil fungsi AddBatchBook
			clientService.AddBatchBook()
		case 3:
			fmt.Println("Getting all books..")
			// memanggil fungsi GetBooks
			clientService.GetBooks()
		case 4:
			fmt.Println("Getting book..")
			// memanggil fungsi GetBook
			clientService.GetBook()
		case 5:
			fmt.Println("Updating book..")
			// memanggil fungsi UpdateBook
			clientService.UpdateBook()
		case 6:
			fmt.Println("Deleting book..")
			// memanggil fungsi DeleteBook
			clientService.DeleteBook()
		case 7:
			fmt.Println("Good Bye")
			// ganti done menjadi true
			// agar program dapat dihentikan
			done = true
		default:
			// jika input tidak valid
			// tampilkan pesan berikut
			fmt.Println("Invalid input!")
			fmt.Println("Enter the valid option number")
		}

		// jika done bernilai true
		// hentikan program
		if done {
			break
		}
	}
}
