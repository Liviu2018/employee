package data

import "fmt"

// Employee models an employee entity and a row in the employee.Employee table
type Employee struct {
	Name      string `json:"name"`
	ID        int    `json:"id"`
	ManagerID int    `json:"manager_id"`
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee[name: %s, ID: %d, ManagerID: %d]", e.Name, e.ID, e.ManagerID)
}
