package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID_Book    int `gorm:"primaryKey;"`
	ID_Owner   int `gorm:"foreignKey:ID_User"`
	Title_Book string
	Status     bool
	Posted_at  time.Time `gorm:"autoCreateTime"`
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

func (ab *AksesBooks) HapusBuku(ID_Books int) bool {
	postExc := ab.DB.Delete(&Books{}, ID_Books)

	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (ab *AksesBooks) TambahBuku(newBook Books) Books {
	err := ab.DB.Create(&newBook).Error

	if err != nil {
		log.Fatal(err)
		return Books{}
	}

	return newBook
}
