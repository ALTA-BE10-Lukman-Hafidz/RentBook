package entity

import (
	"log"

	"gorm.io/gorm"
)

type Books struct {
	ID_Book    int `gorm:"primaryKey;"`
	ID_Owner   int `gorm:"foreignKey:ID_User"`
	Title_Book string
	Status     bool
	Posted_at  string
	Rents      []Rents
}

type AksesBooks struct {
	DB *gorm.DB
}

func (ab *AksesBooks) GetAllData() []Books {
	var daftarBooks = []Books{}
	// err := au.DB.Raw("Select * from Users").Scan(&daftarUsers)
	err := ab.DB.Find(&daftarBooks)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarBooks
}
