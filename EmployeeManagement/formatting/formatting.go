package formatting

import (
	"github.com/Liviu2018/employee/EmployeeManagement/data"
)

// FormatHierarchically takes the list of employee, and produces a list of formated rows
// each row is a list of cells, and corresponds to one Employee
// all cells of a row are empty, except one, containing the name of that Employee
// the first row is the CEO, on its first cell is the CEO name
// all other rows are the employees indented with tabs to their manager
func FormatHierarchically(input []data.Employee) ([]int, []string) {
	parentIndexes := computeParentIndexes(input)

	tree := buildTree(input, parentIndexes)

	return computeFormat(tree, len(input))
}

// computeParentIndexes takes a slise of data.Employee as input
// it computes, for each input element, the index of its parent:
// input[result[i]] = the direct manager of input[i]
func computeParentIndexes(input []data.Employee) []int {
	idToIndex := mapIDToIndex(input)

	parents := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		parents[i] = idToIndex[input[i].ManagerID]
	}

	return parents
}

// mapIDToIndex produces a map, when for each ID its key will be the
// index, in the input slice, of the employee with that ID
func mapIDToIndex(input []data.Employee) map[int]int {
	result := make(map[int]int)

	for i := 0; i < len(input); i++ {
		result[input[i].ID] = i
	}

	return result
}
