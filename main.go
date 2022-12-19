package main

import (
	"fmt"
	"tokoku/config"
	"tokoku/employee"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var employeeAuth = employee.EmployeeAuth{DB: conn}

	//Menu Login
	loginMenu := 0
	for loginMenu != 9 {
		fmt.Print("\n===TOKOKU Project===\n\n")
		fmt.Println("1. Login")
		fmt.Println("2. About")
		fmt.Println("9. Exit")
		fmt.Print("Enter an option : ")
		fmt.Scanln(&loginMenu)

		switch loginMenu {
		case 1:
			username, password := "", ""
			fmt.Print("\n--- Login\n")
			fmt.Print("Username : ")
			fmt.Scanln(&username)
			fmt.Print("Password : ")
			fmt.Scanln(&password)

			if username == "admin" && password == "admin" {
				//Menu Utama Admin
				menuAdm := 0
				for menuAdm != 9 {
					fmt.Print("\n--- Administrator Menu\n")
					fmt.Println("1. Add new employee")
					fmt.Println("2. Delete item")
					fmt.Println("3. Delete customer")
					fmt.Println("4. Delete transaction")
					fmt.Println("9. Log out")
					fmt.Print("Enter an option : ")
					fmt.Scanln(&menuAdm)

					switch menuAdm {
					case 1:
						var newEmp employee.Employee
						var tmpS string
						fmt.Print("Inseret username :")
						fmt.Scanln(&tmpS)
						newEmp.SetUsername(tmpS)
						fmt.Print("Insert name ")
						fmt.Scanln(&tmpS)
						newEmp.SetName(tmpS)
						fmt.Print("Insert password ")
						fmt.Scanln(&tmpS)
						newEmp.SetPassword(tmpS)
						res, err := employeeAuth.RegisterEmp(newEmp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Register Success")
						} else {
							fmt.Println("Register Failed")
						}
					case 2:

					case 3:

					case 4:

					default:
						fmt.Println("\nSorry, option doesn't exist")
					}
				}
			} else {
				//Menu Utama Pegawai
				menuEmp := 0
				for menuEmp != 9 {
					fmt.Println("--- Employee Menu\n")
					fmt.Println("1. Insert item")
					fmt.Println("2. Edit item")
					fmt.Println("3. Update quantity")
					fmt.Println("4. Add new customer")
					fmt.Println("5. Add item to cart")
					fmt.Println("6. Show cart")
					fmt.Println("9. Log out")
					fmt.Print("Enter an option : ")
					fmt.Scanln(&menuEmp)

					switch menuEmp {
					case 1:

					case 2:

					case 3:

					case 4:

					case 5:

					case 6:

					default:
						fmt.Println("\nSorry, option doesn't exist")
					}
				}
			}
		case 2:

		default:
			fmt.Println("\nSorry, option doesn't exist")
		}
	}
}
