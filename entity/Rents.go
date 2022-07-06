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

	// Users     []Users `gorm:"many2many:User_Rents;"`
	// Books     []Books `gorm:"many2many:User_Books;"`
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
