package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"tokoku/item"
)

type Transaction struct {
	idTrans    int
	empName    string
	custName   string
	createDate string
}

func (t *Transaction) SetIdTrans(newIdTrans int) {
	t.idTrans = newIdTrans
}

func (t *Transaction) SetEmpName(newEmpName string) {
	t.empName = newEmpName
}

func (t *Transaction) SetCustName(newCustName string) {
	t.custName = newCustName
}

func (t *Transaction) SetCreateDate(newCreateDate string) {
	t.createDate = newCreateDate
}

// method

func (t *Transaction) GetIdTrans() int {
	return t.idTrans
}

func (t *Transaction) GetEmpName() string {
	return t.empName
}

func (t *Transaction) GetCustName() string {
	return t.custName
}

func (t *Transaction) GetCreateDate() string {
	return t.createDate
}

type TransAuth struct {
	DB *sql.DB
}

type ItemTransaction struct {
	idItem   int
	itemName string
	qty      int
}

func (ia *ItemTransaction) GetIdItem() int {
	return ia.idItem
}
func (ia *ItemTransaction) GetItemName() string {
	return ia.itemName
}
func (ia *ItemTransaction) GetQty() int {
	return ia.qty
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

// func (ta *TransAuth) DeleteItemTrans(transId int) (bool, error) {

// }

// func (ta *TransAuth) DeleteTrans(transId int) (bool, error) {

// }

func (ta *TransAuth) GetTransaction() map[int]*Transaction {
	rows, err := ta.DB.Query("SELECT id_transaction, name, customer_name, created_date FROM transactions t JOIN customers c ON t.no_hp = c.no_hp JOIN employees e ON t.id_employee = e.id_employee")
	if err != nil {
		fmt.Println(errors.New("error select query"))
	}
	defer rows.Close()

	transList := make(map[int]*Transaction)
	tmpIdT := 0
	tmpEmpName, tmpCustName := "", ""
	var tmpDate string

	for rows.Next() {
		err := rows.Scan(&tmpIdT, &tmpEmpName, &tmpCustName, &tmpDate)
		if err != nil {
			fmt.Println(errors.New("error scan"))
		}
		transList[tmpIdT] = &Transaction{tmpIdT, tmpEmpName, tmpCustName, tmpDate}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return transList
}

func (ta *TransAuth) GetItemsTransaction(idTrans int) []ItemTransaction {
	rows, err := ta.DB.Query("SELECT it.id_item, i.item_name, it.quantity FROM item_transaction it JOIN items i ON it.id_item = i.id_item WHERE it.id_transaction  = ?", idTrans)
	if err != nil {
		fmt.Println(errors.New("error select query"))
	}
	defer rows.Close()

	listItemTrans := []ItemTransaction{}
	itemTrans := ItemTransaction{}

	for rows.Next() {
		err := rows.Scan(&itemTrans.idItem, &itemTrans.itemName, &itemTrans.qty)
		if err != nil {
			errors.New("error scan ")
		}

		listItemTrans = append(listItemTrans, itemTrans)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return listItemTrans
}
