package employee

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
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
	// // password validation
	length := len(newEmp.GetPassword())
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Cstr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numb := "0123456789"

	if length < 8 {
		return false, errors.New("\nPassword atleast 8 character")
	} else if !strings.ContainsAny(newEmp.GetPassword(), numb) || !strings.ContainsAny(newEmp.GetPassword(), str) {
		return false, errors.New("\nPassword need combination number and character")
	} else if !strings.ContainsAny(newEmp.GetPassword(), Cstr) {
		return false, errors.New("\nPassword need capital character")
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

func (em *EmployeeAuth) ShowEmps() {
	rows, err := em.DB.Query("SELECT id_employee, username, name FROM employees")
	if err != nil {
		errors.New("Error select query")
	}
	defer rows.Close()

	tmpId := 0
	tmpUname, tmpName := "", ""
	var emp Employee
	var emps []Employee
	for rows.Next() {
		err := rows.Scan(&tmpId, &tmpUname, &tmpName)
		if err != nil {
			errors.New("Error scan ")
		}
		emp.SetId(tmpId)
		emp.SetUsername(tmpUname)
		emp.SetName(tmpName)
		emps = append(emps, emp)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range emps {
		fmt.Println("")
		fmt.Println("ID Employee\t\t: ", v.id)
		fmt.Println("Employee Username\t: ", v.username)
		fmt.Println("Employee Name\t\t: ", v.name)
	}
	if len(emps) == 0 {
		fmt.Println("list employee active has no record")
	}
}

func (em *EmployeeAuth) DeleteEmp(username string) (bool, error) {
	deleteQry, err := em.DB.Prepare("DELETE FROM employees WHERE username = ?")
	if err != nil {
		return false, errors.New("Error delete query")
	}

	res, err := deleteQry.Exec(username)
	if err != nil {
		return false, errors.New("username not match")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, errors.New("Error after delete")
	}
	if affectedRows <= 0 {
		return false, errors.New("0 affected rows")

	}
	return true, nil
}
