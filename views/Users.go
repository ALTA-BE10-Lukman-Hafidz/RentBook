package views

import (
	"bufio"
	"fmt"
	"os"
	"rentbook/controller"
	"rentbook/entity"
)

func LoginUser(AksesUsers entity.AksesUsers) controller.DataUserLogin {
	var No_HP, Pass string

	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&No_HP)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&Pass)

	if AksesUsers.LoginUsers(No_HP).Pass == Pass {
		fmt.Println("Login Sukses")
		DataUserLogin := controller.GetLoginData(AksesUsers, No_HP)
		return DataUserLogin

	} else {
		fmt.Println("Login Gagal")
		return controller.DataUserLogin{}
	}
}

func Register(AksesUsers entity.AksesUsers) {
	var newuser entity.Users
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan Nama: ")
	scanner.Scan()
	newuser.Name = scanner.Text()

	fmt.Print("Masukan No_HP: ")
	fmt.Scanln(&newuser.No_HP)

	fmt.Print("Masukan Pass: ")
	fmt.Scanln(&newuser.Pass)

	fmt.Print("Masukan Email: ")
	fmt.Scanln(&newuser.Email)

	status := AksesUsers.TambahUser(newuser)
	if status {
		fmt.Println("Berhasil register akun")
	} else {
		fmt.Println("gagal register akun")
	}
}

func UpdateProfile(AksesUsers entity.AksesUsers, ID_User int) {
	var newprofile entity.Users
	scanner := bufio.NewScanner(os.Stdin)
	newprofile.ID_User = ID_User

	fmt.Print("Masukan Nama: ")
	scanner.Scan()
	newprofile.Name = scanner.Text()

	fmt.Print("Masukan No_HP: ")
	fmt.Scanln(&newprofile.No_HP)

	fmt.Print("Masukan Pass: ")
	fmt.Scanln(&newprofile.Pass)

	fmt.Print("Masukan Email: ")
	fmt.Scanln(&newprofile.Email)

	status := AksesUsers.UpdateUser(newprofile, ID_User)
	if status {
		fmt.Println("Data Telah Diupdate")
	} else {
		fmt.Println("Data Gagal Diupdate")
	}
}

func UserProfile(AksesUsers entity.AksesUsers) {
	var ID_User int

	fmt.Print("Masukkan ID User yang ingin dilihat: ")
	fmt.Scanln(&ID_User)

	fmt.Println(AksesUsers.GetDataUser(ID_User))
}

func DeleteAccount(AksesRents entity.AksesRents, ID_User int) {

	// if AksesRents.MyRentData(ID_User) != 0 {
	// 	fmt.Println("Anda belum mengembalikan buku")
	// } else if AksesBooks.DeleteOwner(ID_Owner) = 0 {

	// }
}
