package views

import (
	"fmt"
	"rentbook/entity"
	"time"
)

func RentBook(AksesRents entity.AksesRents, ID_User int) {
	var newRent entity.Rents
	var total int
	rented := time.Now()

	fmt.Println("Masukan jumlah buku yang ingin dipinjam : ")
	fmt.Scanln(&total)

	newRent.ID_Renter = ID_User
	newRent.Return_at = rented.AddDate(0, 0, 7)
	newRent.Status_Rent = true

	for i := 0; i < total; i++ {
		fmt.Println("Masukan Id buku : ")
		fmt.Scanln(&newRent.ID_Rentbook)
		AksesRents.TambahRentData(newRent)
	}
}
