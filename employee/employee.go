package employee

import (
	"database/sql"
	"errors"
)

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
func (e *Employee) GetPassword() string {
	return e.password
}

type EmployeeAuth struct {
	DB *sql.DB
}

func (em *EmployeeAuth) Duplicate(username string) bool {
	res := em.DB.QueryRow("SELECT id_employee FROM employees WHERE username = ?", username)
	tmp := 0
	err := res.Scan(&tmp)
	if err != nil {
		return false
	}
	return true
}

func (em *EmployeeAuth) RegisterEmp(newEmp Employee) (bool, error) {
	registerQry, err := em.DB.Prepare("INSERT INTO employees (username,name,password) values (?,?,?)")
	if err != nil {
		return false, errors.New("Column not match")
	}
	//duplicate
	if em.Duplicate(newEmp.GetUsername()) {
		return false, errors.New("\nUsername already exist")
	}
	//execute
	res, err := registerQry.Exec(newEmp.GetUsername(), newEmp.GetName(), newEmp.GetPassword())
	if err != nil {
		return false, errors.New("Insert Employees Errors")
	}

	// Chek rows affected
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("Error after insert")
	}
	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")
	}
	return true, nil
}

func (em *EmployeeAuth) Login(username, password string) (Employee, error) {
	emp := Employee{}
	id := 0
	name := ""
	err := em.DB.QueryRow(
		"SELECT id_employee, name FROM employees WHERE username = ? AND password = ?", username, password).
		Scan(&id, &name)
	emp.SetId(id)
	emp.SetName(name)

	if err != nil {
		return emp, errors.New("\nUser not found or wrong password")
	}
	return emp, nil
}
