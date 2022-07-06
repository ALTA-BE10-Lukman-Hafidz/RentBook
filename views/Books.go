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

func UpdateBook(AksesBooks entity.AksesBooks, ID_User int) {
	var newBook entity.Books
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan ID buku: ")
	fmt.Scanln(&newBook.ID_Book)

	fmt.Print("Masukan judul buku: ")
	scanner.Scan()
	newBook.Title_Book = scanner.Text()

	AksesBooks.UpdateBookData(newBook, ID_User)
}

func PostBook(AksesBooks entity.AksesBooks, scanner *bufio.Scanner, ID_User int) {
	var newBook entity.Books

	newBook.ID_Owner = ID_User

	fmt.Print("Masukan Judul Buku: ")
	scanner.Scan()
	newBook.Title_Book = scanner.Text()

	newBook.Status = true

	AksesBooks.TambahBuku(newBook)

	fmt.Println("Berhasil Post Buku")
}

func DeleteBook(AksesBooks entity.AksesBooks, ID_User int) bool {
	var ID_Book int

	fmt.Print("Masukkan ID Buku yang akan dihapus: ")
	fmt.Scanln(&ID_Book)

	if AksesBooks.HapusBuku(ID_Book, ID_User) {
		fmt.Println("Buku berhasil dihapus")
		return true
	} else {
		fmt.Println("Gagal menghapus buku")
		return false
	}

}
