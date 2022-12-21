package item

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
func (i *Item) GetQuantity() int {
	return i.quantity
}

type ItemAuth struct {
	DB *sql.DB
}

func (ia *ItemAuth) InsertItem(newItem Item) (bool, error) {
	InsertQry, err := ia.DB.Prepare("INSERT INTO items (id_employee,item_name,quantity) values (?,?,?)")
	if err != nil {

		return false, errors.New("column items not match")
	}

	res, err := InsertQry.Exec(newItem.GetIDEmployee(), newItem.GetItemName(), newItem.GetQuantity())
	if err != nil {

		return false, errors.New("error insert query")
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

func (ia *ItemAuth) ShowItems() {
	rows, err := ia.DB.Query("SELECT * FROM items")
	if err != nil {
		errors.New("error select query")
	}
	defer rows.Close()

	tmpId, tmpIdE, tmpQ := 0, 0, 0
	tmpName := ""
	var item Item
	var items []Item
	for rows.Next() {
		err := rows.Scan(&tmpId, &tmpIdE, &tmpName, &tmpQ)
		if err != nil {
			errors.New("Error scan ")
		}
		item.SetIdItem(tmpId)
		item.SetIdEmployee(tmpIdE)
		item.SetItemName(tmpName)
		item.SetQuantity(tmpQ)
		items = append(items, item)
	}
	// tanya mas jerry
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range items {
		fmt.Println("")
		fmt.Println("ID Item\t\t: ", v.idItem)
		fmt.Println("ID Employee\t: ", v.idEmployee)
		fmt.Println("Item Name \t: ", v.itemName)
		fmt.Println("Item Quantity\t: ", v.quantity)
	}
}

func (ia *ItemAuth) DeleteItem(idItem int) (bool, error) {
	deleteQry, err := ia.DB.Prepare("DELETE FROM items WHERE id_item = ?")
	if err != nil {
		return false, errors.New("error delete query")
	}

	res, err := deleteQry.Exec(idItem)
	if err != nil {
		return false, errors.New("idItem not match")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("error after delete")
	}
	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")

	}
	return true, nil
}

func (ia *ItemAuth) UpdateQty(idItem, qty int) (bool, error) {
	UpdateQry, err := ia.DB.Prepare("UPDATE items SET quantity = ? WHERE id_item = ?")
	if err != nil {
		return false, errors.New("error update query")
	}
	res, err := UpdateQry.Exec(qty, idItem)
	if err != nil {
		return false, errors.New("id_item not exist")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("error after update query")
	}

	if affectedRows <= 0 {

		return false, errors.New("0 affected rows")
	}

	return true, nil
}

func (ia *ItemAuth) EditItems(idItem int, newName string) (bool, error) {
	EditQry, err := ia.DB.Prepare("UPDATE  items  SET  item_name = ?   WHERE id_item = ?")
	if err != nil {
		return false, errors.New("error update query")
	}

	res, err := EditQry.Exec(newName, idItem)
	if err != nil {
		return false, errors.New("id_item not exist")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("error after update")
	}

	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")
	}
	return true, nil

}

func (ia *ItemAuth) ItemList() map[int]*Item {
	rows, err := ia.DB.Query("SELECT * FROM items")
	if err != nil {
		errors.New("error select query")
	}
	defer rows.Close()

	// m := make(map[int]Item)
	tmpId, tmpIdE, tmpQ := 0, 0, 0
	tmpName := ""
	// var item Item
	// var items []Item
	items := make(map[int]*Item)
	for rows.Next() {
		err := rows.Scan(&tmpId, &tmpIdE, &tmpName, &tmpQ)
		if err != nil {
			errors.New("error scan ")
		}
		// item.SetIdItem(tmpId)
		// item.SetIdEmployee(tmpIdE)
		// item.SetItemName(tmpName)
		// item.SetQuantity(tmpQ)
		items[tmpId] = &Item{tmpId, tmpIdE, tmpName, tmpQ}
		// items[tmpId].idItem = tmpId
		// items[tmpId].idEmployee = tmpIdE
		// items[tmpId].itemName = tmpName
		// items[tmpId].quantity = tmpQ
	}
	// tanya mas jerry
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return items
}
