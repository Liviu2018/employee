package formatting

import "github.com/Liviu2018/employee/EmployeeManagement/data"

// buildTree build a tree, each node corresponds to one employee
// tree size = number of employees; root node = CEO
// for each node, its parent node corresponds to its manager
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

// addNodesInPath takes a slice of managers (a chain of command: employee, his manager, his
// manager's maanger, ...., the CEO) and iterates through it backwards, from CEO to the initial employee
// it adds nodes in our tree, one node for every new employee it encounter; it has the advantage that
// it always knows the parent (currentNode) of any new to be inserted node
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
