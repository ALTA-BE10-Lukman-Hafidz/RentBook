package main

import (
	"fmt"
	"rentbook/config"
	"rentbook/entity"
	"rentbook/views"
)

// func Scan(str string) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	str = scanner.Text()
// }

func PostBook() {

}

func UpdateBook() {

}

func DeleteBook() {

}

func main() {
	conn := config.InitDB(config.ReadEnv())
	var input, Login, ID_User int

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
			fmt.Println("5. Update Profile")
		}

		fmt.Println("99. Exit")
		fmt.Print("Masukan input: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			views.ReadBooks(AksesBooks)

		case 2:
			if Login == 0 {
				res := views.LoginUser(AksesUsers)
				if res.Status == 1 {
					Login = 1
					ID_User = res.ID_User
					fmt.Println(ID_User)
				}

			} else if Login == 1 {
				PostBook()
			}

		case 3:
			if Login == 0 {
				views.Register(AksesUsers)
			} else if Login == 1 {
				UpdateBook()
			}

		case 4:
			if Login == 0 {
				fmt.Println("Silahkan Login terlebih dulu")
			} else if Login == 1 {
				DeleteBook()
			}

		case 5:
			if Login == 1 {
				DeleteBook()
			}

		case 99:
			//Exit
			fmt.Println("Terima kasih telah menggunakan program kami")
		}
	}

}
