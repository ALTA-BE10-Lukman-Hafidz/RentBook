package views

import (
	"fmt"
	"rentbook/entity"
)

func ReadBooks(AksesBooks entity.AksesBooks) {
	fmt.Println("Daftar Buku")
	for _, val := range AksesBooks.GetAllData() {
		fmt.Println(val)
	}
}
