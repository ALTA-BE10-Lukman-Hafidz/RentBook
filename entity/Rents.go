package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Rents struct {
	ID_Rent   int `gorm:"primaryKey;"`
	ID_Renter int `gorm:"foreignKey:ID_User"`
	ID_Book   int
	Rented_at time.Time `gorm:"autoCreateTime"`
	Return_at time.Time
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
