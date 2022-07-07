package main

import (
	"bufio"
	"fmt"
	"os"
	"rentbook/config"
	"rentbook/entity"
	"rentbook/views"
)

// func Scan(str string) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	str = scanner.Text()
// }

func main() {
	conn := config.InitDB(config.ReadEnv())
	var input, Login, ID_User int
	scanner := bufio.NewScanner(os.Stdin)

	AksesUsers := entity.AksesUsers{DB: conn}
	AksesBooks := entity.AksesBooks{DB: conn}
	AksesRents := entity.AksesRents{DB: conn}

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
			fmt.Println("6. Rent Books")
			fmt.Println("7. Return Books")
			fmt.Println("8. My Books")
			fmt.Println("9. Get User Profile")
			fmt.Println("10. Delete Account")
			fmt.Println("11. My Rent Book")
		}

		fmt.Println("99. Exit")
		fmt.Print("Masukan input: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			views.ReadBooks(AksesBooks)
			// fmt.Println(AksesUsers.GetAllData())

		case 2:
			if Login == 0 {
				res := views.LoginUser(AksesUsers)

				if res.Status == 1 {
					Login = 1
					ID_User = res.ID_User
				}

			} else if Login == 1 {
				views.PostBook(AksesBooks, scanner, ID_User)
			}

		case 3:
			if Login == 0 {
				views.Register(AksesUsers)
			} else if Login == 1 {
				views.UpdateBook(AksesBooks, ID_User)
			}

		case 4:
			if Login == 0 {
				fmt.Println("Silahkan Login terlebih dulu")
			} else if Login == 1 {
				// Delete Books
				views.DeleteBook(AksesBooks, ID_User)
			}

		case 5:
			if Login == 1 {
				views.UpdateProfile(AksesUsers, ID_User)
			}

		case 6:
			views.RentBook(AksesRents, AksesBooks, ID_User)

		case 7:
			//return book
			views.ReturnBook(AksesRents, ID_User)

		case 8:
			//My book
			views.MyBooks(AksesBooks, ID_User)

		case 9:
			//Get user profile
			views.UserProfile(AksesUsers)

		case 10:
			//Hapus Akun
			//views.DeleteAccount(AksesUsers, ID_User)

		case 11:
			//My Rent Book
			views.MyRentBooks(AksesRents, ID_User)

		case 99:
			//Exit
			fmt.Println("Terima kasih telah menggunakan program kami")
		}
	}

}
