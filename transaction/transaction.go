package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"tokoku/item"
)

type Transaction struct {
	idTrans    int
	idEmp      int
	noHP       string
	createDate string
}

func (t *Transaction) SetIdTrans(newIdTrans int) {
	t.idTrans = newIdTrans
}

func (t *Transaction) SetidEmp(newIdEmp int) {
	t.idEmp = newIdEmp
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

func (t *Transaction) GetNohp() string {
	return t.noHP
}

func (t *Transaction) GetCreateDate() string {
	return t.createDate
}

type TransAuth struct {
	DB *sql.DB
}

func (ta *TransAuth) insertTransaction(idEmp int, noHp string) (bool, error) {
	InsertQry, err := ta.DB.Prepare("INSERT INTO transactions (id_employee,no_hp,created_date) values (?,?,?)")
	if err != nil {
		return false, errors.New("error insert query")
	}

	res, err := InsertQry.Exec(idEmp, noHp, time.Now())
	if err != nil {
		return false, errors.New("error exec query")
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		return false, errors.New("error After insert query")
	}

	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")
	}
	return true, nil
}

func (ta *TransAuth) insertItemTrans(idEmp int, noHp string, cart map[int]*item.Item) (bool, error) {
	idTrans := 0
	err := ta.DB.QueryRow(
		"SELECT id_transaction FROM transactions WHERE id_employee = ? AND no_hp = ? ORDER BY id_transaction DESC LIMIT 1", idEmp, noHp).
		Scan(&idTrans)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range cart {
		InsertQry, err := ta.DB.Prepare("INSERT INTO item_transaction (id_item,id_transaction,quantity) values (?,?,?)")
		if err != nil {
			return false, errors.New("error insert query")
		}

		res, err := InsertQry.Exec(v.GetIdItem(), idTrans, v.GetQuantity())
		if err != nil {
			return false, errors.New("error exec query")
		}

		affectedRows, err := res.RowsAffected()

		if err != nil {
			return false, errors.New("error After insert query")
		}

		if affectedRows <= 0 {
			return false, errors.New("0 affected rows")
		}
	}
	return true, nil
}

func (ta *TransAuth) Checkout(idEmp int, noHp string, cart map[int]*item.Item) (bool, error) {
	_, err := ta.insertTransaction(idEmp, noHp)
	if err != nil {
		return false, errors.New("error after insert transaction")
	}
	_, err = ta.insertItemTrans(idEmp, noHp, cart)
	if err != nil {
		return false, errors.New("error after insert into item transaction")
	}

	return true, nil
}

// func (ta *TransAuth) PrintReceipt(transId int) (bool, error) {

// }

// func (ta *TransAuth) DeleteTrans(transId int) (bool, error) {

// }

// func (ta *TransAuth) DeleteItemTrans(transId int) (bool, error) {

// }
