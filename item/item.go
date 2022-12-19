package item

import "database/sql"

type Item struct {
	idItem     int
	idEmployee int
	itemName   string
	quantity   int
}

//method item

func (i *Item) SetIdItem(newIdItem int){
	i.idItem = newIdItem
}
func (i *Item) SetIdEmployee(newIdEmployee int){
	i.idEmployee = newIdEmployee
}

func (i *Item) SetItemName(newItemName string){
	i.itemName = newItemName
}

func (i *Item) SetQuantity(newQuantity int){
	i.quantity = newQuantity
}

func (i *Item) getIdItem() int {
	return i.idItem
}

func (i *Item) getIDEmployee() int {
	return i.id
}

func (i *Item) getItemName() int {
	return i.idItemName
}





type ConnectSQL struct{
	DB *sql.DB
}

