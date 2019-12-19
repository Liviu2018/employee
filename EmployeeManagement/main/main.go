package main

import (
	"fmt"
	"github.com/Liviu2018/employee/EmployeeManagement/database"

	"github.com/Liviu2018/employee/EmployeeManagement/data"
)

func main() {
	database.Init()
	defer database.Close()

	err := database.AddEmployee(data.Employee{Name: "bbb", ID: 2, ManagerID: 1})
	if err != nil {
		fmt.Println(err)
	}

	all, err := database.GetAllEmployees()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(all)
}
