package main

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/employee"
	"tokoku/item"
	"tokoku/transaction"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var employeeAuth = employee.EmployeeAuth{DB: conn}
	var itemAuth = item.ItemAuth{DB: conn}
	var custAuth = customer.CustAuth{DB: conn}
	var transAuth = transaction.TransAuth{DB: conn}
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
					fmt.Print("\n--- Administrator Menu\n\n")
					fmt.Println("1. Add new employee")
					fmt.Println("2. Delete employee")
					fmt.Println("3. Delete item")
					fmt.Println("4. Delete customer")
					fmt.Println("5. Delete transaction")
					fmt.Println("9. Log out")
					fmt.Print("Enter an option : ")
					fmt.Scanln(&menuAdm)

					switch menuAdm {
					case 1:
						var newEmp employee.Employee
						var tmpS string
						fmt.Print("\n---Employee Register\n")
						fmt.Print("Insert username : ")
						fmt.Scanln(&tmpS)
						newEmp.SetUsername(tmpS)
						fmt.Print("Insert name : ")
						scanner.Scan()
						tmpS = scanner.Text()
						newEmp.SetName(tmpS)
						fmt.Print("Insert password : ")
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
						employeeAuth.ShowEmps()
						var username string
						fmt.Print("Insert username : ")
						fmt.Scanln(&username)
						res, err := employeeAuth.DeleteEmp(username)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete Employee success")
						} else {
							fmt.Println("Delete Employee failed")
						}
					case 3:
						itemAuth.ShowItems()
						var idItem int
						fmt.Print("Insert id item : ")
						fmt.Scanln(&idItem)
						res, err := itemAuth.DeleteItem(idItem)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete item success")
						} else {
							fmt.Println("Delete item failed")
						}

					case 4:
						custAuth.ShowCust()
						var noHp string
						fmt.Print("Insert Phone Number : ")
						fmt.Scanln(&noHp)
						res, err := custAuth.DeleteCust(noHp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete Customer success")
						} else {
							fmt.Println("Delete Customer  failed")
						}
					case 5:

					case 9:
						continue
					default:
						fmt.Println("\nSorry, option doesn't exist")
					}
				}
			} else {
				emp, err := employeeAuth.Login(username, password)
				if err != nil {
					fmt.Println(err.Error())
				}

				//Menu Utama Pegawai
				if emp.GetName() != "" {
					menuEmp := 0
					for menuEmp != 9 {
						fmt.Print("\n--- Employee Menu\n")
						fmt.Print("\nWelcome back, ", emp.GetName(), "\n")
						fmt.Println("1. Insert item")
						fmt.Println("2. Show All items")
						fmt.Println("3. Edit item")
						fmt.Println("4. Update quantity")
						fmt.Println("5. Add new customer")
						fmt.Println("6. Transaction")
						fmt.Println("9. Log out")
						fmt.Print("Enter an option : ")
						fmt.Scanln(&menuEmp)

						switch menuEmp {
						case 1:
							var newItem item.Item
							fmt.Print("Insert item name : ")
							scanner.Scan()
							tmps := scanner.Text()
							newItem.SetItemName(tmps)
							fmt.Print("Insert quantity : ")
							tmpi := 0
							fmt.Scanln(&tmpi)
							newItem.SetQuantity(tmpi)
							newItem.SetIdEmployee(emp.GetId())
							res, err := itemAuth.InsertItem(newItem)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Insert item Success")

							} else {
								fmt.Println("Insert item failed")
							}
						case 2:
							fmt.Println("List Items")
							itemAuth.ShowItems()
						case 3:
							itemAuth.ShowItems()
							var idItem int
							var newName string
							fmt.Print("Insert Id item : ")
							fmt.Scanln(&idItem)
							fmt.Print("Insert New Name  : ")
							fmt.Scanln(&newName)
							res, err := itemAuth.EditItems(idItem, newName)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Update Name Success")

							} else {
								fmt.Println("Update Name failed")
							}

						case 4:
							itemAuth.ShowItems()
							var idItem, qty int
							fmt.Print("Insert Id item : ")
							fmt.Scanln(&idItem)
							fmt.Print("Insert New Quantity : ")
							fmt.Scanln(&qty)
							res, err := itemAuth.UpdateQty(idItem, qty)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Update Quantity Success")

							} else {
								fmt.Println("Update Quantity failed")
							}
						case 5:
							var newCust customer.Customer
							tmps := ""
							fmt.Print("Insert Phone number : ")
							fmt.Scanln(&tmps)
							newCust.SetNohp(tmps)
							fmt.Print("Insert Customer Name : ")
							scanner.Scan()
							tmps = scanner.Text()
							newCust.SetName(tmps)
							newCust.SetIdEmployee(emp.GetId())
							res, err := custAuth.InsertCust(newCust)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Insert New Customer Success")

							} else {
								fmt.Println("Insert New Customer failed")
							}
						case 6:
							//Menu Transaksi
							transMenu := 0
							cart := map[int]*item.Item{}
							items := itemAuth.ItemList()

							for transMenu != 9 {
								fmt.Println("\n--- Transaction Menu")
								fmt.Println("1. Show all items")
								fmt.Println("2. Add item to cart")
								fmt.Println("3. Show cart")
								fmt.Println("4. Checkout")
								fmt.Println("9. Back")
								fmt.Print("Enter an option : ")
								fmt.Scanln(&transMenu)
								switch transMenu {
								case 1:
									fmt.Print("\n--- All items\n")
									itemAuth.ShowItems()
								case 2:
									fmt.Print("\n--- Add item to cart\n")
									var itemCart item.Item
									idItem, qty := 0, 0
									for _, v := range items {
										if v.GetQuantity() > 0 {
											fmt.Print("Item ID : ", v.GetIdItem(), "\n")
											fmt.Print("Item name : ", v.GetItemName(), "\n")
											fmt.Print("Quantity available : ", v.GetQuantity(), "\n\n")
										}
									}
									fmt.Print("Insert ID item : ")
									fmt.Scanln(&idItem)
									fmt.Print("Insert quantity : ")
									fmt.Scanln(&qty)
									if qty > items[idItem].GetQuantity() {
										fmt.Println("\nThe quantity is higher than the maximum available")
										continue
									}

									if thisItem, ok := cart[idItem]; ok {
										thisItem.SetIdItem(idItem)
										thisItem.SetIdEmployee(items[idItem].GetIDEmployee())
										thisItem.SetItemName(items[idItem].GetItemName())
										thisItem.SetQuantity(thisItem.GetQuantity() + qty)
										cart[idItem] = thisItem
									} else {
										itemCart.SetIdItem(idItem)
										itemCart.SetIdEmployee(items[idItem].GetIDEmployee())
										itemCart.SetItemName(items[idItem].GetItemName())
										itemCart.SetQuantity(qty)
										cart[idItem] = &itemCart
									}

									items[idItem].SetQuantity(items[idItem].GetQuantity() - qty)
								case 3:
									fmt.Println("\n--- Cart")
									fmt.Print("Item ID\t\tItem Name\t\tQuantity\n")
									for _, v := range cart {
										fmt.Print(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t", v.GetQuantity(), "\n")
									}
								case 4:
									fmt.Println("\n--- Cart")
									fmt.Print("Item ID\t\tItem Name\t\tQuantity\n")
									for _, v := range cart {
										fmt.Print(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t", v.GetQuantity(), "\n")
									}
									tmp, noHp := "", ""
									fmt.Print("Are you sure want to finish your cart? (y/n) : ")
									fmt.Scanln(&tmp)

									if tmp == "y" {
										fmt.Print("Insert customer phone number : ")
										fmt.Scanln(&noHp)
										res, err := transAuth.Checkout(emp.GetId(), noHp, cart)
										if err != nil {
											fmt.Println(err.Error())
										}
										if res {
											fmt.Println("\nCheckout Success")

										} else {
											fmt.Println("\nCheckout failed")
										}
										for _, v := range cart {
											itemAuth.UpdateQty(v.GetIdItem(), items[v.GetIdItem()].GetQuantity())
										}
									} else if tmp == "n" {
										continue
									} else {
										fmt.Println("Wrong Input")
									}
								case 9:
									continue
								default:
									fmt.Println("\nSorry, option doesn't exist")
								}
							}
						case 9:
							continue
						default:
							fmt.Println("\nSorry, option doesn't exist")
						}
					}
				}
			}
		case 2:

		case 9:
			continue
		default:
			fmt.Println("\nSorry, option doesn't exist")
		}
	}
}
