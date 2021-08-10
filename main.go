package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nadirbasalamah/books-grpc/book/bookpb"
	"github.com/nadirbasalamah/books-grpc/database"
	"github.com/nadirbasalamah/books-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// jika kode mengalami crash, nomor line akan ditampilkan
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Book service started")

	// membuat koneksi ke DB
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// membuat gRPC server
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	// melakukan register BookServiceServer
	bookpb.RegisterBookServiceServer(s, &server.Server{})
	// mengaktifkan reflection
	// agar bisa digunakan untuk pengujian dengan evans
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Menunggu hingga dihentikan dengan Ctrl + C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Lakukan block hingga sinyal sudah didapatkan
	<-ch
	fmt.Println("Stopping the server..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
