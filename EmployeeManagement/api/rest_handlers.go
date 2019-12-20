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
		writeErrorMessage(w, err)

		return
	}

	if !newEmployee.IsValid() {
		panic(fmt.Errorf("Valid employee fields"))
	}

	err = database.AddEmployee(newEmployee)
	if err != nil {
		writeErrorMessage(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Employee %s added successfully", newEmployee.String())))
}

// GetEmployees will return a list of all existing employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resthandlers GetEmployees")

	result, err := database.GetAllEmployees()
	if err != nil {
		writeErrorMessage(w, err)

		return
	}

	formatted := formatting.ConvertToSimpleTable(result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(formatted)
}

// ListAllFormattedEmployees will return a formatted list of all existing employees
func ListAllFormattedEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resthandlers ListAllFormattedEmployees")

	result, err := database.GetAllEmployees()
	if err != nil {
		writeErrorMessage(w, err)

		return
	}

	whiteTabs, names := formatting.FormatHierarchically(result)

	formatted := formatting.ComposeTable(whiteTabs, names)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(formatted)
}

// FaviconHandler serves the favorite icon
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/favicon.png")
}

func writeErrorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}
