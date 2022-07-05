package entity

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Setting struct {
	DBUser     string
	DBPassword string
}

type Books struct {
	ID_Book    int `gorm:"primaryKey;"`
	ID_Owner   int `gorm:"foreignKey:ID_User"`
	Title_Book string
	Status     bool
	Posted_at  string
}

type AksesBooks struct {
	DB *gorm.DB
}

func ReadFromEnv() Setting {
	res := Setting{}
	err := godotenv.Load("local.env")
	if err != nil {
		return Setting{}
	}
	res.DBPassword = os.Getenv("DBPass")
	res.DBUser = os.Getenv("DBUser")
	return res
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
