package formatting

import "github.com/Liviu2018/employee/EmployeeManagement/data"

func buildTree(input []data.Employee, heights map[int][]int) *Node {
	maxHeight := len(heights) - 1
	ceoIndex := heights[0][0]
	result := Node{Name: input[ceoIndex].Name, heigth: 0, left: nil, right: nil, parent: nil}

	for i := 1; i <= maxHeight; i++ {
		for j := 0; j < len(heights[i]); j++ {
			result.Append(input[heights[i][j]])
		}
	}

	return &result
}

func findCEOIndex(heights []int) int {
	for i := 0; i < len(heights); i++ {
		if heights[i] == 0 {
			return i
		}
	}

	return 0
}
