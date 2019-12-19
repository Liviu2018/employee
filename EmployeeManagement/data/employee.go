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

// IsValid checks if the fields of this Employee object meet basic requirements
func (e *Employee) IsValid() bool {
	return e.Name != "" && e.ID >= 0 && e.ManagerID >= 0
}
