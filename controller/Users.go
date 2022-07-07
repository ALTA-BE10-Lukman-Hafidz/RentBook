package controller

import (
	"rentbook/entity"
)

type DataUserLogin struct {
	ID_User int
	Status  int
}

func GetLoginData(AksesUsers entity.AksesUsers, No_HP string) DataUserLogin {
	var DataUserLogin = DataUserLogin{}

	DataUserLogin.ID_User = AksesUsers.LoginUsers(No_HP).ID_User
	DataUserLogin.Status = 1

	return DataUserLogin
}
