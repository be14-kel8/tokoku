package item

import "database/sql"

type Item struct {
	idItem     int
	idEmployee int
	itemName   string
	quantity   int
}

//method set (untuk mempersiapkan properti yang bisa diakses)

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




type ConnectSQL struct{
	DB *sql.DB
}

