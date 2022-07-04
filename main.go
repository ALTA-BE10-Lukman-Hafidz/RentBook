package main

import (
	"bufio"
	"fmt"
	"os"
	"rentbook/config"
	"rentbook/entity"
)

func main() {
	conn := config.InitDB()
	var input int
	scanner := bufio.NewScanner(os.Stdin)
	AksesUsers := entity.AksesUsers{DB: conn}
	// AksesBooks := entity.AksesBooks{DB: conn}
	// AksesRents := entity.AksesRents{DB: conn}

	for input != 99 {
		fmt.Println("=====Perpus Online=====")
		fmt.Println("1. Read Books")
		fmt.Println("2. Login")
		fmt.Println("3. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukan input: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			//Read Books
			//fitur

		case 2:
			//Login

		case 3:
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

		case 99:
			//Exit
			fmt.Println("Terima kasih telah menggunakan program kami")
		}
	}
}
