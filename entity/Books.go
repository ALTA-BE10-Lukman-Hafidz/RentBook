package entity

import (
	"fmt"
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

func (ab *AksesBooks) HapusBuku(ID_Book int, ID_User int) bool {
	postExc := ab.DB.Delete(&Books{}, "ID_Book = ? AND ID_Owner = ?", ID_Book, ID_User)

	if err := postExc.Error; err != nil {
		log.Println(err)
		return false
	}

	if postExc.RowsAffected == 0 {
		fmt.Println("Tidak ada buku dengan id tersebut pada akun anda")
		return false
	}

	return true
}

func (ab *AksesBooks) TambahBuku(newBook Books) Books {
	err := ab.DB.Create(&newBook).Error

	if err != nil {
		log.Println(err)
		return Books{}
	}

	return newBook
}

func (ab *AksesBooks) UpdateBookData(newBook Books, ID_User int) Books {
	err := ab.DB.Model(&Books{}).Where("ID_Book = ? AND ID_Owner = ?", newBook.ID_Book, ID_User).Updates(&newBook)

	if err.Error != nil {
		log.Println(err)
		return Books{}
	}

	if err.RowsAffected == 0 {
		fmt.Println("Tidak ada buku dengan id tersebut pada akun anda")
		return Books{}
	}

	fmt.Println("Data buku berhasil diupdate")

	return newBook
}
