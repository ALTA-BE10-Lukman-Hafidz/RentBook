package entity

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	ID_User    int
	Name       string
	No_HP      string
	Pass       string
	Email      string
	Deleted_at string
}

type AksesUsers struct {
	DB *gorm.DB
}

func (au *AksesUsers) GetAllData() []Users {
	var daftarUsers = []Users{}
	// err := au.DB.Raw("Select * from Users").Scan(&daftarUsers)
	err := au.DB.Find(&daftarUsers, "Deleted_at = 0")

	if err.Error != nil {
		log.Println(err.Statement.SQL.String())
		return nil

	}

	return daftarUsers
}

func (au *AksesUsers) LoginUsers(No_HP string) Users {
	var daftarUsers = Users{}

	err := au.DB.Find(&daftarUsers, "No_HP = ?", No_HP)

	if err.Error != nil {
		log.Println(err)
		return Users{}
	}

	return daftarUsers
}

func (au *AksesUsers) TambahUser(newuser Users) bool {
	err := au.DB.Create(&newuser)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	return true
}

func (au *AksesUsers) UpdateUser(newprofile Users, ID_User int) bool {
	err := au.DB.Model(&Users{}).Where("ID_User = ?", ID_User).Updates(&newprofile)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	return true
}

func (au *AksesUsers) GetDataUser(ID_User int) Users {
	var daftarUsers = Users{}

	err := au.DB.Select("ID_User", "Name", "Email").Find(&daftarUsers, "ID_User = ? AND Deleted_at = 0", ID_User)

	if err.Error != nil {
		log.Println(err)
		return Users{}
	}
	if err.RowsAffected == 0 {
		log.Println("User tidak ditemukan")
		return Users{}
	}

	return daftarUsers
}

func (au *AksesUsers) DeactivatedAccountData(UserData Users, ID_User int) bool {

	err := au.DB.Model(&Users{}).Where("ID_User = ?", ID_User).Updates(&UserData)

	if err.Error != nil {
		log.Println(err)
		return false
	}

	return true
}

func (au *AksesUsers) GetMyProfileData(ID_User int) Users {
	var daftarUsers = Users{}

	err := au.DB.Find(&daftarUsers, "ID_User = ?", ID_User)

	if err.Error != nil {
		log.Println(err)
		return Users{}
	}
	if err.RowsAffected == 0 {
		log.Println("User tidak ditemukan")
		return Users{}
	}

	return daftarUsers
}
