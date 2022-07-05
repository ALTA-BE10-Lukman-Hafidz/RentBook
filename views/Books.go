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

func UpdateBook(AksesBooks entity.AksesBooks, ID_User int, statusbook bool) {
	var newBook entity.Books
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan ID buku: ")
	fmt.Scanln(&newBook.ID_Book)

	newBook.ID_Owner = ID_User

	fmt.Print("Masukan judul buku: ")
	scanner.Scan()
	newBook.Title_Book = scanner.Text()

	newBook.Status = statusbook

	AksesBooks.UpdateBookData(newBook, ID_User)
}
