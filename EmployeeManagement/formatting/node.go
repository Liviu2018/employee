package formatting

import "fmt"

// Node models a node in the employee tree
// that is a tree where the CEO is the root, its direct managers are its chidren
// and each employee node is a child of its immediate manager node
type Node struct {
	Name                string
	heigth              int
	left, right, parent *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node[%s, %d, %v, %v, %v]", n.Name, n.heigth, n.left, n.right, n.parent)
}
