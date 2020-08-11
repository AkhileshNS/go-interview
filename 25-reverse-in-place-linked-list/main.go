package main

import "fmt"

type node struct {
	value int
	next  *node
}

func printList(head *node) {
	pointer := head
	for pointer != nil {
		fmt.Print(pointer.value, " ")
		pointer = pointer.next
	}
}

func reverse(head *node) *node {
	if head == nil {
		return head
	}

	if head.next == nil {
		return head
	}

	if head.next.next == nil {
		pointer := head
		pointer = pointer.next
		pointer.next = head
		head.next = nil
		return pointer
	}

	leader := head
	pointer := head
	follower := head

	leader = leader.next.next
	pointer = pointer.next

	for leader != nil {
		pointer.next = follower
		follower = pointer
		pointer = leader
		leader = leader.next
	}

	pointer.next = follower
	head.next = nil

	return pointer
}

func main() {
	head := &node{value: 1}
	head.next = &node{value: 2}
	head.next.next = &node{value: 3}
	head.next.next.next = &node{value: 4}

	printList(head)
	head = reverse(head)
	fmt.Print("\n")
	printList(head)
}
