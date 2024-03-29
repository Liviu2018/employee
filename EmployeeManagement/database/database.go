package database

import (
	"database/sql"
	"fmt"

	"github.com/Liviu2018/employee/EmployeeManagement/data"
	// we need mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB is our connetion to the database
var DB *sql.DB

// Init initializes the mySql database connection
func Init(datasource string) error {
	var err error
	DB, err = sql.Open("mysql", datasource)

	createTableIfNotExists()

	return err
}

var createTableQuery = "CREATE TABLE IF NOT EXISTS employee.Employee ( " +
	"id INT PRIMARY KEY, " +
	"name VARCHAR(255) NOT NULL, " +
	"manager_id INT);"

// this query will only insert if the table is empty
var initialEmployeesQuery = "INSERT INTO employee.Employee (name, id, manager_id) VALUES " +
	"('Jamie', 150, 150)," +
	"('Alan', 100, 150)," +
	"('Steve', 400, 150)," +
	"('Martin', 220, 100)," +
	"('Alex', 275, 100)," +
	"('David', 190, 400); "

var isEmptyTableQuery = "SELECT EXISTS (SELECT 1 FROM employee.Employee);"

// creates a new employee.Employee table, only if the table does not exist
func createTableIfNotExists() error {
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		return err
	}

	containsElements := false
	err = DB.QueryRow(isEmptyTableQuery).Scan(&containsElements)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if containsElements {
		return nil
	}

	_, err = DB.Exec(initialEmployeesQuery)
	return err
}

var insertQuery = "INSERT INTO employee.Employee (name, id, manager_id) VALUES ( ?, ?, ?)"

// AddEmployee adds a new employee to the employee.Employee table
func AddEmployee(e data.Employee) error {
	// if employee is not the CEO, first check that his manager exists
	if e.ID != e.ManagerID {
		manager, err := containsID(e.ManagerID)
		if err != nil || !manager {
			return fmt.Errorf("Manager %d does not exist", e.ManagerID)
		}
	}

	exists, err := containsID(e.ID)
	if err != nil || exists {
		return fmt.Errorf("ID %d already exists", e.ManagerID)
	}

	insert, err := DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	_, err = insert.Exec(e.Name, e.ID, e.ManagerID)

	return err
}

var containsIDQuery = "SELECT EXISTS (SELECT * FROM employee.Employee WHERE id=?)"

// containsID checks if a row with that ID is already in employee.Employee table
func containsID(ID int) (bool, error) {
	result := false

	err := DB.QueryRow(containsIDQuery, ID).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return result, nil
}

var getEmployeesQuery = "SELECT name, id, manager_id FROM employee.Employee"

// GetAllEmployees retrives all lines in employee.Employee table
func GetAllEmployees() ([]data.Employee, error) {
	rows, err := DB.Query(getEmployeesQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]data.Employee, 0)
	current := &data.Employee{}

	for rows.Next() {
		err = rows.Scan(&current.Name, &current.ID, &current.ManagerID)
		if err != nil {
			return nil, err
		}

		result = append(result, *current)
	}

	return result, nil
}

// Close will close the database connection
func Close() {
	DB.Close()
}
