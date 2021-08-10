package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/nadirbasalamah/books-grpc/database"
	"github.com/nadirbasalamah/books-grpc/model"
)

// AddBook untuk menambahkan data buku
func AddBook(bookData model.Book) model.Book {

	// membuat id menggunakan uuid
	var uuid string = uuid.New().String()

	// memasukkan data buku ke dalam database
	_, err := database.DB.Query("INSERT INTO books (id, title, author, is_read) VALUES (?, ?, ?, ?)",
		uuid,
		bookData.Title,
		bookData.Author,
		bookData.IsRead,
	)

	// jika terdapat error, tampilkan pesan error
	if err != nil {
		log.Fatalf("Insert data failed: %v", err)
		return model.Book{}
	}

	// mengembalikan data buku yang dimasukkan
	return bookData

}

// GetBook untuk mendapatkan data buku berdasarkan id
func GetBook(bookId string) (int, model.Book) {

	// membuat variabel book
	// untuk menyimpan data buku berdasarkan id
	var book model.Book = model.Book{}

	// mendapatkan data buku berdasarkan id
	row, err := database.DB.Query("SELECT * FROM books WHERE id = ?", bookId)

	// menampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return 0, model.Book{}
	}

	// Close() akan dipanggil
	// jika data dari database sudah didapatkan
	defer row.Close()

	// untuk setiap baris data
	for row.Next() {
		// masukkan berbagai atribut data buku seperti judul dll.
		// ke dalam variabel book
		switch err := row.Scan(&book.Id, &book.Title, &book.Author, &book.IsRead); err {

		// jika data tidak ditemukan
		// tampilkan pesan error
		case sql.ErrNoRows:
			log.Printf("Data not found: %v", err)
			return 0, model.Book{}

		// jika tidak ada error
		// tampilkan data buku
		case nil:
			log.Println(book)

		// tampilkan error untuk kondisi default
		default:
			log.Printf("Data cannot be retrieved: %v", err)
			return 0, model.Book{}
		}
	}

	// mengembalikan angka 1 sebagai tanda data ditemukan
	// dan data buku yang ditemukan
	return 1, book

}

// GetBooks untuk mendapatkan seluruh data buku
func GetBooks() []model.Book {

	// mendapatkan seluruh data buku
	rows, err := database.DB.Query("SELECT * FROM books")

	// tampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return []model.Book{}
	}

	// close dipanggil jika
	// seluruh data berhasil diambil
	defer rows.Close()

	// membuat variabel books
	// untuk menampung berbagai data buku
	var books []model.Book = []model.Book{}

	// untuk setiap data
	for rows.Next() {
		// membuat variabel book
		// untuk menyimpan sebuah data buku
		var book model.Book = model.Book{}

		// masukkan berbagai atribut data buku seperti judul dll.
		// ke dalam variabel book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.IsRead)

		// jika terdapat error, tampilkan error
		if err != nil {
			log.Printf("Data cannot be retrieved: %v", err)
			return []model.Book{}
		}

		// masukkan data buku ke dalam books
		books = append(books, book)
	}

	// jika jumlah data di dalam books
	// sama dengan 0, maka data kosong
	if len(books) == 0 {
		log.Println("Books data not found")
	}

	// mengembalikan sekumpulan data buku
	return books

}

// UpdateBook untuk mengedit data buku
func UpdateBook(bookData model.Book, id string) model.Book {

	// mengubah data buku berdasarkan id
	_, err := database.DB.Query("UPDATE books SET title=?, author=?, is_read=? WHERE id=?",
		bookData.Title,
		bookData.Author,
		bookData.IsRead,
		bookData.Id,
	)

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Update data failed: %v", err)
		return model.Book{}
	}

	// mengembalikan data buku yang telah diubah
	return bookData

}

// DeleteBook untuk menghapus data buku
func DeleteBook(id string) bool {

	// menghapus data buku berdasarkan id
	_, err := database.DB.Query("DELETE FROM books WHERE id=?", id)

	// jika terdapat error, tampilkan pesan error
	// nilai false dikembalikan
	if err != nil {
		log.Fatalf("Delete data failed: %v", err)
		return false
	}

	// nilai true dikembalikan jika data berhasil dihapus
	return true
}
