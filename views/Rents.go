package views

import (
	"fmt"
	"rentbook/entity"
	"time"
)

func RentBook(AksesRents entity.AksesRents, AksesBooks entity.AksesBooks, ID_User int) {
	var newRent entity.Rents
	var total int
	rented := time.Now()

	fmt.Println("Masukan jumlah buku yang ingin dipinjam : ")
	fmt.Scanln(&total)

	newRent.ID_Renter = ID_User
	newRent.Return_at = rented.AddDate(0, 0, 7)
	newRent.Status_Rent = true
	Status_Book := false

	for i := 0; i < total; i++ {
		fmt.Println("Masukan Id buku : ")
		fmt.Scanln(&newRent.ID_Rentbook)
		AksesRents.TambahRentData(newRent, Status_Book)
	}
}

func ReturnBook(AksesRents entity.AksesRents, ID_User int) {
	var newReturn entity.Rents
	var total int

	Status_Book := true
	fmt.Print("Masukkan jumlah buku yang ingin dikembalikkan : ")
	fmt.Scanln(&total)

	for i := 0; i < total; i++ {
		fmt.Print("Masukkan Id buku : ")
		fmt.Scanln(&newReturn.ID_Rentbook)
		newReturn.Status_Rent = false
		AksesRents.ReturnBookData(ID_User, newReturn, Status_Book)
	}
}

func MyRentBooks(AksesRents entity.AksesRents, ID_Renter int) {
	fmt.Println("Buku Pinjaman Saya")
	for _, val := range AksesRents.MyRentData(ID_Renter) {
		fmt.Println(val)
	}
}
