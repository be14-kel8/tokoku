package customer

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Customer struct {
	noHp       string
	idEmployee int
	name       string
}

func (c *Customer) SetNohp(newNohp string) {
	c.noHp = newNohp
}
func (c *Customer) SetIdEmployee(newIdEmployee int) {
	c.idEmployee = newIdEmployee
}
func (c *Customer) SetName(newName string) {
	c.name = newName
}

// method get

func (c *Customer) GetNohp() string {
	return c.noHp
}

func (c *Customer) GetIdEmployee() int {
	return c.idEmployee
}

func (c *Customer) GetName() string {
	return c.name
}

type CustAuth struct {
	DB *sql.DB
}

func (ca *CustAuth) InsertCust(newCustomer Customer) (bool, error) {
	InsertQry, err := ca.DB.Prepare("INSERT INTO customers values (?,?,?)")
	if err != nil {
		return false, errors.New("Insert query customers error")
	}

	res, err := InsertQry.Exec(newCustomer.GetNohp(), newCustomer.GetIdEmployee(), newCustomer.GetName())
	if err != nil {
		return false, errors.New("Insert query not match")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("Error after insert")
	}

	if affectedRows <= 0 {
		return false, errors.New(" 0 affected rows")
	}
	return true, nil
}

func (ca *CustAuth) ShowCust() {
	rows, err := ca.DB.Query("SELECT * FROM customers")
	if err != nil {
		errors.New("Error select query")
	}
	defer rows.Close()

	tmpId := 0
	tmpNohp, tmpName := "", ""
	var cust Customer
	var custs []Customer
	for rows.Next() {
		err := rows.Scan(&tmpNohp, &tmpId, &tmpName)
		if err != nil {
			errors.New("Error scan ")
		}
		cust.SetIdEmployee(tmpId)
		cust.SetNohp(tmpNohp)
		cust.SetName(tmpName)
		custs = append(custs, cust)
	}
	// tanya mas jerry
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range custs {
		fmt.Println("")
		fmt.Println("Phone Number \t\t: ", v.noHp)
		fmt.Println("Id Employee\t\t: ", v.idEmployee)
		fmt.Println("Customer Name\t\t: ", v.name)
	}
}

func (ca *CustAuth) DeleteCust(noHp string) (bool, error) {
	deleteQry, err := ca.DB.Prepare("DELETE FROM customers WHERE no_hp = ?")
	if err != nil {
		return false, errors.New("Error delete query")
	}

	res, err := deleteQry.Exec(noHp)
	if err != nil {
		return false, errors.New("Phone Number not exist")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("Error after delete")
	}
	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")

	}
	return true, nil
}
