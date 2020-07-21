package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) addLeft(value int) {
	n.left = &node{value: value}
}

func (n *node) addRight(value int) {
	n.right = &node{value: value}
}

func inOrderTraverse(values *[]int, tree *node) {
	if tree == nil {
		return
	}

	inOrderTraverse(values, tree.left)
	newValues := append(*values, tree.value)
	*values = newValues
	inOrderTraverse(values, tree.right)
}

func checkBST(tree *node) bool {
	values := []int{}
	inOrderTraverse(&values, tree)
	fmt.Println(values)
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			return false
		}
	}

	return true
}

func main() {
	tree := &node{value: 50}
	tree.addRight(80)
	tree.addLeft(30)
	tree.left.addLeft(20)
	tree.left.addRight(60)
	tree.right.addLeft(70)
	tree.right.addRight(90)

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(checkBST(tree))
}
