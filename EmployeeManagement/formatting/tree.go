package formatting

import "github.com/Liviu2018/employee/EmployeeManagement/data"

func buildTree(input []data.Employee, parentIndexes []int) *Node {
	var result *Node

	for i := 0; i < len(input); i++ {
		path := computePathToRoot(parentIndexes, i)

		if result == nil {
			rootNodeIndex := path[len(path)-1]

			result = &Node{Name: input[rootNodeIndex].Name}
		}

		addNodesInPath(result, path, input)
	}

	return result
}

func addNodesInPath(root *Node, path []int, input []data.Employee) {
	currentNode := root
	currentIndex := path[len(path)-1]

	for currentIndex >= 0 {
		currentNode = currentNode.Insert(input[path[currentIndex]])

		currentIndex--
	}
}

func computePathToRoot(parents []int, current int) []int {
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
