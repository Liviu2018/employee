package formatting

import "fmt"

func computeFormat(root *Node, sizeOfTree int) ([]int, []string) {
	whiteTabsCount := make([]int, sizeOfTree)
	displayedNames := make([]string, sizeOfTree)
	f := Format{whiteTabsCount, displayedNames, sizeOfTree - 1}

	computeFormatRecursive(root, &f, 0)

	return whiteTabsCount, displayedNames
}

func computeFormatRecursive(current *Node, f *Format, heigth int) {
	if current == nil {
		return
	}

	if current.Children == nil || len(current.Children) == 0 {
		f.whiteTabsCount[f.index] = heigth
		f.displayedNames[f.index] = current.Name
		f.index--

		return
	}

	// sort children by Name
	current.SortChildren()

	for i := len(current.Children) - 1; i >= 0; i-- {
		computeFormatRecursive(current.Children[i], f, 1+heigth)
	}

	f.whiteTabsCount[f.index] = heigth
	f.displayedNames[f.index] = current.Name
	f.index--
}

// Format is an auxiliary object that we need to append line by line the format of the next line in our result
// since we compute the result starting with the last line, then the previous and so on
// we need format to internally keep the index value on the next successfully computed line
type Format struct {
	whiteTabsCount []int
	displayedNames []string
	index          int
}

func printTreeRecursive(n *Node, height int) {
	s := ""
	for i := 0; i < height; i++ {
		s += "_"
	}

	fmt.Println(s + n.String())

	for i := 0; i < len(n.Children); i++ {
		printTreeRecursive(n.Children[i], 1+height)
	}
}
