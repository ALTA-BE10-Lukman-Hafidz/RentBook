package main

import (
	"fmt"
	// "rentbook/config"
	// "rentbook/entity"
)

func main() {
	// conn := config.InitDB()
	var input int

	// AksesUsers := entity.AksesUsers{DB: conn}
	// AksesBooks := entity.AksesBooks{DB: conn}
	// AksesRents := entity.AksesRents{DB: conn}

	for input != 0 {
		fmt.Println("=====Perpus Online=====")
		fmt.Println("1. Read Books")
		fmt.Println("2. Login")
		fmt.Println("3. Register")
		fmt.Print("Masukan input: ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			//Read Books
		case 2:
			//Login
		case 3:
			//Register
		case 0:
			//Exit
			fmt.Println("Terima kasih telah menggunakan program kami")
		}
	}

}
