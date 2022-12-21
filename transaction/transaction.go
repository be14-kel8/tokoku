package transaction

import (
	"database/sql"
)

type Transaction struct {
	idTrans    int
	idEmp      int
	idItem     int
	noHP       string
	createDate string
}

func (t *Transaction) SetIdTrans(newIdTrans int) {
	t.idTrans = newIdTrans
}

func (t *Transaction) SetidEmp(newIdEmp int) {
	t.idEmp = newIdEmp
}

func (t *Transaction) SetIdItem(newIdItem int) {
	t.idItem = newIdItem
}

func (t *Transaction) SetNohp(newNohp string) {
	t.noHP = newNohp
}

func (t *Transaction) SetCreateDate(newCreateDate string) {
	t.createDate = newCreateDate
}

// method

func (t *Transaction) GetIdTrans() int {
	return t.idTrans
}

func (t *Transaction) GetIdEmployee() int {
	return t.idEmp
}

func (t *Transaction) GetIdItem() int {
	return t.idItem
}

func (t *Transaction) GetNohp() string {
	return t.noHP
}

func (t *Transaction) GetCreateDate() string {
	return t.createDate
}

type TransAuth struct {
	DB *sql.DB
}
