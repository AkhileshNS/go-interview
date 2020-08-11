package main

import "fmt"

type node struct {
	value int
	next  *node
}

func kthLastNode(head *node, last int) int {
	tracker := 0
	pointer := head
	follower := head

	for pointer != nil {
		if tracker < last {
			tracker++
		} else {
			follower = follower.next
		}
		pointer = pointer.next
	}

	if tracker < last {
		return -1
	}
	return follower.value
}

func main() {
	head := &node{value: 1}
	head.next = &node{value: 2}
	head.next.next = &node{value: 3}
	head.next.next.next = &node{value: 4}

	fmt.Println(kthLastNode(head, 4))
}
