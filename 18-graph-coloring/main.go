package main

import "fmt"

var colors = []string{"violet", "indigo", "blue", "green", "yellow", "orange", "red"}

type node struct {
	value      int
	color      string
	neighbours []*node
}

func colorGraph(graph []*node) {
	maxDegree := -1

	for _, n := range graph {
		if len(n.neighbours) > maxDegree {
			maxDegree = len(n.neighbours)
		}
	}

	options := colors[:maxDegree+1]

	for i := range graph {
		isLoop := false
		illegalColors := make(map[string]bool)

		for _, n := range graph[i].neighbours {
			if n == graph[i] {
				isLoop = true
			}

			if n.color != "" {
				illegalColors[n.color] = true
			}
		}

		if isLoop {
			break
		}

		for _, color := range options {
			if _, ok := illegalColors[color]; !ok {
				graph[i].color = color
				break
			}
		}
	}
}

func main() {
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	node4 := &node{value: 4}
	node5 := &node{value: 5}
	node6 := &node{value: 6}
	node7 := &node{value: 7}

	node1.neighbours = []*node{node2, node3}
	node2.neighbours = []*node{node4, node5, node1}
	node3.neighbours = []*node{node6, node7, node1}
	node4.neighbours = []*node{node2}
	node5.neighbours = []*node{node2, node6}
	node6.neighbours = []*node{node3, node5}
	node7.neighbours = []*node{node7, node3}

	graph := []*node{node1, node2, node3, node4, node5, node6, node7}

	colorGraph(graph)

	for _, n := range graph {
		fmt.Println(n.color, n.value)
	}
}
