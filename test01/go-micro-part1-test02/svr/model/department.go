package model

import (
	"go-micro/go-micro-part1-test02/svr/basic/db"
	us "go-micro/go-micro-part1-test02/svr/proto"
)

func GetDepartments() ([]*us.Department, error) {

	var departments []*us.Department

	database := db.GetDb()
	rows, err := database.Query(`SELECT dept_no, dept_name
										FROM departments`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		department := &us.Department{}
		if err := rows.Scan(&department.DeptNo, &department.DeptName); err != nil {
			return nil, err
		}
		departments = append(departments, department)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil

}
