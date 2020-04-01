package model

import (
	"go-micro/go-micro-part1-test01/svr/basic/db"
	us "go-micro/go-micro-part1-test01/svr/proto"
)

func GetEmployees() ([]*us.Employee, error) {

	var employees []*us.Employee

	database := db.GetDb()
	rows, err := database.Query(`SELECT emp_no, birth_date, first_name, last_name, gender, hire_date 
										FROM employees WHERE emp_no <= ?`, 10010)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		employee := &us.Employee{}
		if err := rows.Scan(&employee.Empno, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
			return nil, err
		}
		employees = append(employees, employee)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil

}
