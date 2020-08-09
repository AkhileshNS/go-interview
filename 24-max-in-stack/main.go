package main

import "fmt"

type node struct {
	value int
	next  *node // points to next node
	prev  *node // points to previous max
}

type stack struct {
	head *node
	max  *node
}

func (s *stack) push(value int) {
	n := &node{value: value}
	if s.max == nil {
		s.max = n
	} else if s.max.value <= value {
		n.prev = s.max
		s.max = n
	}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
}

func (s *stack) peek() int {
	if s.head == nil {
		fmt.Println("Head pointing to null")
		fmt.Println(s.head)
		return -678
	}
	return s.head.value
}

func (s *stack) pop() int {
	if s.head == nil {
		fmt.Println("Head pointing to null")
		fmt.Println(s.head)
		return -678
	}

	value := s.head.value

	if s.max == s.head {
		s.max = s.max.prev
	}

	s.head = s.head.next

	return value
}

func (s *stack) getMax() int {
	if s.head == nil {
		fmt.Println("Head pointing to null")
		fmt.Println(s.head)
		return -678
	}
	return s.max.value
}

func main() {
	s := stack{}
	s.push(0)
	s.push(1)
	s.push(-1)
	s.push(2)
	fmt.Println(s.getMax())

	s.pop()
	fmt.Println(s.getMax())
	s.pop()
	fmt.Println(s.getMax())
	s.pop()
	fmt.Println(s.getMax())
}
