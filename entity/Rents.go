package entity

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Rents struct {
	ID_Rent     int
	ID_Renter   int
	ID_Rentbook int
	Rented_at   time.Time `gorm:"autoCreateTime"`
	Return_at   time.Time
	Status_Rent bool
}

type AksesRents struct {
	DB *gorm.DB
}

func (ar *AksesRents) GetAllData() []Rents {
	var daftarRents = []Rents{}
	// err := au.DB.Raw("Select * from Users").Scan(&daftarUsers)
	err := ar.DB.Find(&daftarRents)

	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())

		return nil
	}

	return daftarRents
}

func (ar *AksesRents) TambahRentData(newRent Rents) Rents {
	err := ar.DB.Create(&newRent).Error

	if err != nil {
		log.Println(err)

		return Rents{}
	}
	fmt.Println("Berhasil Data berhasil masuk")
	return newRent
}

//select * from book where id_owner == id_user
func (ar *AksesRents) ReturnBookData(ID_User int, newReturn Rents, ID_Rentbook int) Rents {
	err := ar.DB.Model(&Rents{}).Where("ID_Rentbook = ? AND ID_User = ?", ID_Rentbook, ID_User).Updates(&newReturn)

	if err.Error != nil {
		log.Println(err)
		return Rents{}
	}

	if err.RowsAffected == 0 {
		fmt.Println("Tidak ada buku yang dikembalikan")
		return Rents{}
	}

	fmt.Println("Return buku berhasil")

	return newReturn
}
