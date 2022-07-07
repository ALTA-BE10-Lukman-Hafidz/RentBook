package views

import (
	"bufio"
	"fmt"
	"os"
	"rentbook/entity"
)

func ReadBooks(AksesBooks entity.AksesBooks) {
	fmt.Println("Daftar Buku")
	for _, val := range AksesBooks.GetAllData() {
		fmt.Println(val)
	}
}

func AvailableBooks(AksesBooks entity.AksesBooks) {
	fmt.Println("Daftar buku yang dapat dipinjam")
	for _, val := range AksesBooks.GetAvailableBooks() {
		fmt.Println(val)
	}
}

func UpdateBook(AksesBooks entity.AksesBooks, ID_User int) {
	var newBook entity.Books
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan ID buku: ")
	fmt.Scanln(&newBook.ID_Book)

	fmt.Print("Masukan judul buku: ")
	scanner.Scan()
	newBook.Title_Book = scanner.Text()

	status := AksesBooks.UpdateBookData(newBook, ID_User)

	if status {
		fmt.Println("Data buku berhasil diupdate")
	}
}

func PostBook(AksesBooks entity.AksesBooks, scanner *bufio.Scanner, ID_User int) {
	var newBook entity.Books

	newBook.ID_Owner = ID_User

	fmt.Print("Masukan Judul Buku: ")
	scanner.Scan()
	newBook.Title_Book = scanner.Text()

	newBook.Status = true

	status := AksesBooks.TambahBuku(newBook)

	if status {
		fmt.Println("Berhasil Post Buku")
	} else {
		fmt.Println("Gagal Post Buku")
	}
}

func DeleteBook(AksesBooks entity.AksesBooks, ID_User int) {
	var ID_Book int

	fmt.Print("Masukkan ID Buku yang akan dihapus: ")
	fmt.Scanln(&ID_Book)

	status := AksesBooks.HapusBuku(ID_Book, ID_User)
	if status {
		fmt.Println("Buku berhasil dihapus")
	} else {
		fmt.Println("Gagal menghapus buku")
	}
}

func MyBooks(AksesBooks entity.AksesBooks, ID_Owner int) {
	fmt.Println("Daftar Buku Saya")
	for _, val := range AksesBooks.GetMyBooks(ID_Owner) {
		fmt.Println(val)
	}
}
