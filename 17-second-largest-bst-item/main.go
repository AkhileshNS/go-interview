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

func get2Largest(tree *node) (int, int) {
	follower := tree
	start := tree.right

	for start.right != nil {
		follower = start
		start = start.right
	}

	return start.value, follower.value
}

func find2ndLargest(tree *node) int {
	if tree == nil || (tree.left == nil && tree.right == nil) {
		return -1
	}

	if tree.right == nil {
		start, _ := get2Largest(tree.left)
		return start
	} else {
		_, follower := get2Largest(tree)
		return follower
	}
}

func main() {
	tree := &node{value: 50}
	// tree.addRight(80)
	tree.addLeft(30)
	tree.left.addLeft(20)
	tree.left.addRight(40)
	// tree.right.addLeft(70)
	// tree.right.addRight(90)

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(find2ndLargest(tree))
}
