package formatting

import (
	"fmt"
	"github.com/Liviu2018/employee/EmployeeManagement/data"
)

// FormatHierarchically takes the list of employee, and produces a list of formated rows
// each row is a list of cells, and corresponds to one Employee
// all cells of a row are empty, except one, containing the name of that Employee
// the first row is the CEO, on its first cell is the CEO name
// all other rows are the employees indented with tabs to their manager
func FormatHierarchically(input []data.Employee) [][]string {
	parentIndexes := computeParentIndexes(input)

	heights := computeHeights(parentIndexes)
	groupedHeights := groupHeights(heights)

	tree := buildTree(input, groupedHeights)

	fmt.Println(tree)

	return nil
}

// groupHeights receives a list of heights, and groups them in a map
// in the resulting map, each key is a key value (from 0 to max height)
// and the value of each key is a slice with the indexes inside height that contain that value
func groupHeights(heights []int) map[int][]int {
	result := make(map[int][]int)

	for i := 0; i < len(heights); i++ {
		if _, ok := result[heights[i]]; !ok {
			result[heights[i]] = make([]int, 0)
		}

		result[heights[i]] = append(result[heights[i]], i)
	}

	return result
}

func computeHeights(parents []int) []int {
	heigths := make([]int, len(parents))

	for i := 0; i < len(parents); i++ {
		startIndex := i

		// as long as we have not reached the CEO
		// and the height of the currentIndex has not already been computed
		for startIndex != parents[startIndex] && heigths[parents[startIndex]] == 0 {
			startIndex = parents[startIndex]

			heigths[i]++
		}

		// we have reached a parent with previous calculated heigth, add that
		heigths[i] += heigths[startIndex]
	}

	return heigths
}

// computeParentIndexes takes a slise of data.Employee as input
// it computes, for each input element, the index of its parent:
// input[result[i]] = the direct manager of input[i]
func computeParentIndexes(input []data.Employee) []int {
	parents := make([]int, len(input))
	idToIndex := mapIDToIndex(input)

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
