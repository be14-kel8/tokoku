package customer

import (
	"database/sql"
	"errors"
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

type CustAuth struct{
	DB *sql.DB
}

func (ca *CustAuth) InsertCust(newCustomer Customer) (bool,error){
	InsertQry, err := ca.DB.Prepare("INSERT INTO customers values (?,?,?)")
	if err != nil{
		return false, errors.New("Insert query customers error")
	}

	res, err := InsertQry.Exec(newCustomer.GetNohp(),newCustomer.GetIdEmployee(),newCustomer.GetName())
	if err != nil{
		return false, errors.New("Insert query not match")
	}

	affectedRows,err := res.RowsAffected()
	if err != nil{
		return false, errors.New("Error after insert")
	}

	if affectedRows <=0 {
		return false, errors.New(" 0 affected rows")
	}
	return true, nil
}
