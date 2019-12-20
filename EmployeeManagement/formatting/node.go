package formatting

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Liviu2018/employee/EmployeeManagement/data"
)

// Node models a node in the employee tree
// that is a tree where the CEO is the root, its direct managers are its chidren
// and each employee node is a child of its immediate manager node
type Node struct {
	Name     string
	ID       int
	Parent   *Node
	Children []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node[%s, %d]", n.Name, n.ID)
}

// SortChildren sorts the children of the current node, by name alphabetically
func (n *Node) SortChildren() {
	if n == nil || n.Children == nil || len(n.Children) == 0 {
		return
	}

	sort.SliceStable(n.Children, func(i, j int) bool {
		return strings.Compare(n.Children[i].Name, n.Children[j].Name) > 0
	})
}

// Insert adds a new child to the current node, with the values of its employee paramter
// if child already exists, it does nothing
// returns the newly inserted child node
func (n *Node) Insert(e data.Employee) *Node {
	child := n.findChildWithValue(e)

	if child != nil { // it already exists
		return child
	}

	// we initialize new nodes to an empty slice, not nil
	newNode := Node{Name: e.Name, ID: e.ID, Parent: n, Children: []*Node{}}
	n.Children = append(n.Children, &newNode)

	return &newNode
}

func (n *Node) findChildWithValue(e data.Employee) *Node {
	if n == nil || n.Children == nil || len(n.Children) == 0 {
		return nil
	}

	for i := 0; i < len(n.Children); i++ {
		if n.Children[i].ID == e.ID {
			return n.Children[i]
		}
	}

	return nil
}
