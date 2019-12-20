package formatting

import "fmt"

import "github.com/Liviu2018/employee/EmployeeManagement/data"

import "strconv"

// ComposeTable takes the employee names (in the order we want them displayed)
// and how many white tabs before each name
// produces a table, each row has one employee name and empty strings
func ComposeTable(whiteTabs []int, names []string) [][]string {
	if whiteTabs == nil || names == nil || len(whiteTabs) != len(names) {
		fmt.Println("ComposeTable: invalid input !")
		return nil
	}

	maxTabs := getMax(whiteTabs)

	result := make([][]string, len(whiteTabs))

	for i := 0; i < len(whiteTabs); i++ {
		result[i] = make([]string, 1+maxTabs)

		result[i][whiteTabs[i]] = names[i]
	}

	return result
}

// ConvertToSimpleTable takes a slice of employees and returns a table (name, ID, managerID)
func ConvertToSimpleTable(input []data.Employee) [][]string {
	if input == nil || len(input) == 0 {
		return nil
	}

	result := make([][]string, len(input))
	for i := 0; i < len(result); i++ {
		result[i] = []string{input[i].Name, strconv.Itoa(input[i].ID), strconv.Itoa(input[i].ManagerID)}
	}

	return result
}

// getMax returns the max of an int slice
func getMax(input []int) int {
	result := input[0]

	for i := 1; i < len(input); i++ {
		if result < input[i] {
			result = input[i]
		}
	}

	return result
}
