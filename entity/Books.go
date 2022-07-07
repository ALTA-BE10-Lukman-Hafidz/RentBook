package entity

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID_Book    int
	ID_Owner   int
	Title_Book string
	Status     bool
	Posted_at  time.Time `gorm:"autoCreateTime"`
}

type AksesBooks struct {
	DB *gorm.DB
}

func (ab *AksesBooks) GetAllData() []Books {
	var daftarBooks = []Books{}

	err := ab.DB.Find(&daftarBooks)

	if err.Error != nil {
		log.Println(err.Statement.SQL.String())
		return nil
	}

	return daftarBooks
}

func (ab *AksesBooks) GetAvailableBooks() []Books {
	var daftarBooks = []Books{}

	err := ab.DB.Find(&daftarBooks, "Status = true")

	if err.Error != nil {
		log.Println(err.Statement.SQL.String())
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

func (ab *AksesBooks) TambahBuku(newBook Books) bool {
	err := ab.DB.Create(&newBook).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (ab *AksesBooks) UpdateBookData(newBook Books, ID_User int) bool {
	err := ab.DB.Model(&Books{}).Where("ID_Book = ? AND ID_Owner = ?", newBook.ID_Book, ID_User).Updates(&newBook)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	if err.RowsAffected == 0 {
		fmt.Println("Tidak ada buku dengan id tersebut pada akun anda")
		return false
	}

	return true
}

func (ab *AksesBooks) GetMyBooks(ID_User int) []Books {
	var daftarBooks = []Books{}

	err := ab.DB.Find(&daftarBooks, "ID_Owner = ?", ID_User)

	if err.Error != nil {
		log.Println(err.Statement.SQL.String())
		return nil
	}

	if err.RowsAffected == 0 {
		log.Println("Tidak ada buku pada akun anda")
		return nil
	}

	return daftarBooks
}
