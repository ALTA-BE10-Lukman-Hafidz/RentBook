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

	err := ar.DB.Find(&daftarRents)

	if err.Error != nil {
		log.Println(err.Statement.SQL.String())
		return nil
	}

	return daftarRents
}

func (ar *AksesRents) TambahRentData(newRent Rents, Status_Book bool) bool {
	err := ar.DB.Create(&newRent)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	status := ar.DB.Exec("update books b join rents r on r.ID_Rentbook = b.ID_Book and r.ID_Rentbook = ? and r.ID_Renter = ? set b.Status = ?",
		newRent.ID_Rentbook, newRent.ID_Renter, Status_Book)

	if status.Error != nil {
		log.Println(status)
		return false
	}

	if status.RowsAffected == 0 {
		fmt.Println("Buku tidak tersedia")
		return false
	}

	return true
}

func (ar *AksesRents) ReturnBookData(ID_User int, newReturn Rents, Status_Book bool) bool {
	status := ar.DB.Exec("update books b join rents r on r.ID_Rentbook = b.ID_Book and r.ID_Rentbook = ? and r.ID_Renter = ? set b.Status = ?",
		newReturn.ID_Rentbook, ID_User, Status_Book)

	if status.Error != nil {
		log.Println(status)
		return false
	}

	if status.RowsAffected == 0 {
		fmt.Println("Buku sudah dikembalikan")
		return false
	}

	err := ar.DB.Model(&Rents{}).Where("ID_Rentbook = ? AND ID_Renter = ?",
		newReturn.ID_Rentbook, ID_User).Update("Status_Rent", newReturn.Status_Rent)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	if err.RowsAffected == 0 {
		fmt.Println("Tidak ada buku pinjaman pada akun")
		return false
	}

	return true
}

func (ar *AksesRents) MyRentData(ID_Renter int) []Rents {
	var daftarRent = []Rents{}

	err := ar.DB.Where("ID_Renter = ? AND Status_Rent = True", ID_Renter).Find(&daftarRent)

	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())

		return nil
	}

	return daftarRent
}
