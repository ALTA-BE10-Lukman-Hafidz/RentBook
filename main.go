package main

import (
	"bufio"
	"fmt"
	"os"
	"rentbook/config"
	"rentbook/entity"
)

func ReadBooks(AksesBooks entity.AksesBooks) {
	fmt.Println("Daftar Buku")
	for _, val := range AksesBooks.GetAllData() {
		fmt.Println(val)
	}
}

func LoginUser(AksesUsers entity.AksesUsers) int {
	var No_HP, Pass string
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&No_HP)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&Pass)

	if AksesUsers.LoginUsers(No_HP) == Pass {
		fmt.Println("Login Sukses")
		return 1
	} else {
		fmt.Println("Login Gagal")
		return 0
	}
}

func Register(AksesUsers entity.AksesUsers, scanner *bufio.Scanner) {
	var newuser entity.Users

	fmt.Print("Masukan Nama: ")
	scanner.Scan()
	newuser.Name = scanner.Text()

	fmt.Print("Masukan No_HP: ")
	fmt.Scanln(&newuser.No_HP)

	fmt.Print("Masukan Pass: ")
	fmt.Scanln(&newuser.Pass)

	fmt.Print("Masukan Email: ")
	fmt.Scanln(&newuser.Email)

	AksesUsers.TambahUser(newuser)
}

func main() {
	conn := config.InitDB()
	var input, Login int

	scanner := bufio.NewScanner(os.Stdin)
	AksesUsers := entity.AksesUsers{DB: conn}
	AksesBooks := entity.AksesBooks{DB: conn}
	//AksesRents := entity.AksesRents{DB: conn}

	for input != 99 {
		fmt.Println("=====Perpus Online=====")
		fmt.Println("1. Read Books")
		if Login == 0 {
			fmt.Println("2. Login")
			fmt.Println("3. Register")
		} else if Login == 1 {
			fmt.Println("2. Post Books")
			fmt.Println("3. Update Books")
			fmt.Println("4. Delete Books")
		}
		fmt.Println("99. Exit")
		fmt.Print("Masukan input: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			ReadBooks(AksesBooks)

		case 2:
			//Login
			if Login == 0 {
				if LoginUser(AksesUsers) == 1 {
					Login = 1
				}
			} else if Login == 1 {
				// Post Books
			}
		case 3:
			if Login == 0 {
				Register(AksesUsers, scanner)
			} else if Login == 1 {
				// Update Books
			}
		case 4:
			if Login == 0 {
				fmt.Println("Silahkan login terlebih dahulu")
			} else if Login == 1 {
				// Delete Books
			}
		case 99:
			//Exit
			fmt.Println("Terima kasih telah menggunakan program kami")
		}
	}

}
