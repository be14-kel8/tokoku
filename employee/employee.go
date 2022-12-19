package employee

import "database/sql"

type Employee struct {
	id       int
	username string
	name     string
	password string
}

func (e *Employee) SetId(newEmployeeId int) {
	//cara akses object nya menggunakan method.property
	e.id = newEmployeeId

}

func (e *Employee) SetUsername(newUserName string) {
	e.username = newUserName
}

func (e *Employee) SetName(newName string) {
	e.name = newName
}

func (e *Employee) SetPassword(newPassword string) {
	e.password = newPassword

}

// methodd get

func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetUsername() string {
	return e.username
}
func (e *Employee) GetName() string {
	return e.name
}
func (e *Employee) GePassword() string {
	return e.password
}

type EmployeeAuth struct {
	DB *sql.DB
}

