package item

import (
	"database/sql"
	"errors"
)

type Item struct {
	idItem     int
	idEmployee int
	itemName   string
	quantity   int
}

//method set (untuk mempersiapkan properti yang bisa diakses)

func (i *Item) SetIdItem(newIdItem int) {
	i.idItem = newIdItem
}
func (i *Item) SetIdEmployee(newIdEmployee int) {
	i.idEmployee = newIdEmployee
}

func (i *Item) SetItemName(newItemName string) {
	i.itemName = newItemName
}

func (i *Item) SetQuantity(newQuantity int) {
	i.quantity = newQuantity
}

// method Get ( untuk menginisialisasi atribute dengan method)
func (i *Item) GetIdItem() int {
	return i.idItem
}

func (i *Item) GetIDEmployee() int {
	return i.idEmployee
}

func (i *Item) GetItemName() string {
	return i.itemName
}
func (i *Item) GetNewQuantity() int {
	return i.quantity
}

type ItemAuth struct {
	DB *sql.DB
}

func (ia *ItemAuth) InsertItem(newItem Item) (bool, error) {
	InsertQry, err := ia.DB.Prepare("INSERT INTO items (id_employee,item_name,quantity) values (?,?,?)")
	if err != nil {

		return false, errors.New("Column items not match")
	}

	res, err := InsertQry.Exec(newItem.idEmployee, newItem.itemName, newItem.quantity)
	if err != nil {

		return false, errors.New("Error insert query")
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {

		return false, errors.New("Error After insert query")
	}

	if affectedRows <= 0 {

		return false, errors.New("0 affected rows")
	}
	return true, nil
}
