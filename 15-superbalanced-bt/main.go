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

type nodeWithDepth struct {
	n node
	d int
}

func checkIfSuperBalanced(tree node) bool {
	depths := make(map[int]bool)
	nodes := []nodeWithDepth{nodeWithDepth{n: tree, d: 0}}

	for len(nodes) > 0 {
		curr := nodes[len(nodes)-1].n
		depth := nodes[len(nodes)-1].d
		nodes = nodes[:len(nodes)-1]

		if curr.left == nil && curr.right == nil {
			depths[depth] = true

			count := 0
			max := -999999999
			min := 999999999

			for d := range depths {
				count++
				if d > max {
					max = d
				}
				if d < min {
					min = d
				}

				if count > 2 {
					return false
				}
			}

			if (max - min) > 1 {
				return false
			}
		} else {
			if curr.left != nil {
				nodes = append(nodes, nodeWithDepth{n: *curr.left, d: depth + 1})
			}
			if curr.right != nil {
				nodes = append(nodes, nodeWithDepth{n: *curr.right, d: depth + 1})
			}
		}
	}

	return true
}

func main() {
	tree := node{value: 0}
	tree.addRight(1)
	tree.addLeft(-2)
	tree.left.addLeft(-3)
	tree.left.addRight(-1)
	tree.left.left.addLeft(3)

	fmt.Println(checkIfSuperBalanced(tree))
}
