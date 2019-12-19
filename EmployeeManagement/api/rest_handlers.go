package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Liviu2018/employee/EmployeeManagement/data"
	"github.com/Liviu2018/employee/EmployeeManagement/formatting"

	"github.com/Liviu2018/employee/EmployeeManagement/database"
)

// CreateEmployee will add a new employee into the database
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resthandlers CreateEmployee")

	var newEmployee data.Employee

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEmployee)
	if err != nil {
		panic(err)
	}

	if !newEmployee.IsValid() {
		panic(fmt.Errorf("Valid employee fields"))
	}

	err = database.AddEmployee(newEmployee)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Employee %s added successfully", newEmployee.String())))
}

// ListAllEmployees will return a list of all existing employees
func ListAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resthandlers ListAllEmployees")

	result, err := database.GetAllEmployees()
	if err != nil {
		panic(err)
	}

	whiteTabs, names := formatting.FormatHierarchically(result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Found %d employees: %v, %v", len(whiteTabs), whiteTabs, names)))
}
