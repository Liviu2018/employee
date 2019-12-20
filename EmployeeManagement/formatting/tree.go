package formatting

import "github.com/Liviu2018/employee/EmployeeManagement/data"

func buildTree(input []data.Employee, parentIndexes []int) *Node {
	var result *Node

	for i := 0; i < len(input); i++ {
		path := computePathToRoot(parentIndexes, i)

		if result == nil {
			// root node will be the last node in this path
			rootNodeIndex := path[len(path)-1]
			rootNode := input[path[rootNodeIndex]]

			// we initialize children to an empty slice, not nil
			result = &Node{Name: rootNode.Name, ID: rootNode.ID, Children: []*Node{}}
		}

		addNodesInPath(result, path, input)
	}

	return result
}

func addNodesInPath(root *Node, path []int, input []data.Employee) {
	currentNode := root
	currentIndex := len(path) - 1
	currentIndex-- // root is already added

	for currentIndex >= 0 {
		currentNode = currentNode.Insert(input[path[currentIndex]])

		currentIndex--
	}
}

// currentIndex produces a slice containing current node index,
// the index of his parent, then of his parents parent, and so on,
// ending with the company's CEO
func computePathToRoot(parents []int, current int) []int {
	// starting point is the current node
	result := []int{current}

	// as long as we have not reached the CEO
	for current != parents[current] {
		current = parents[current]

		result = append(result, current)
	}

	return result
}
