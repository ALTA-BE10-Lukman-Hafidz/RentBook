package entity

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Users struct {
	ID_User int `gorm:"primaryKey;"`
	Name    string
	No_HP   string
	Pass    string
	Email   string
	Books   []Books `gorm:"foreignKey:ID_Book"`
}

type AksesUsers struct {
	DB *gorm.DB
}

func (au *AksesUsers) GetAllData() []Users {
	var daftarUsers = []Users{}
	// err := au.DB.Raw("Select * from Users").Scan(&daftarUsers)
	err := au.DB.Find(&daftarUsers)

	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil

	} else {
		fmt.Println("Berhasil register akun")
	}

	return daftarUsers
}

func (au *AksesUsers) TambahUser(newuser Users) Users {
	err := au.DB.Create(&newuser).Error

	if err != nil {
		log.Println(err)

		return Users{}
	}

	return newuser
}

func (au *AksesUsers) LoginUsers(No_HP string) Users {
	var daftarUsers = Users{}
	daftarUsers.No_HP = No_HP
	err := au.DB.Raw("Select * from Users").Scan(&daftarUsers)
	if err != nil {
		log.Println(err.Statement.SQL.String())

		return Users{}
	}

	return daftarUsers
}

func (au *AksesUsers) LoginDataUser(No_HP string, Pass string) Users {
	var daftarUsers = Users{}
	err := au.DB.Where("No. HP= ? AND Pass = ?", No_HP, Pass).Find(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())

	}
	return daftarUsers
}
