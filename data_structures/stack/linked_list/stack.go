package main

import "fmt"

// LIFO stack by using linked list.
type stack struct {
	top *node
}

type node struct {
	val  int
	next *node
}

// Remove data from the top of the stack.
func (s *stack) pop() {
	// LIFO
	s.top = s.top.next
}

// Add the item to the top of the stack.
func (s *stack) push(item int) {
	// LIFO
	s.top = &node{
		val:  item,
		next: s.top,
	}
}

// Returns the top item from the stack.
func (s *stack) peek() *node {
	return s.top
}

// Returns true if the stack is empty.
func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) traverse() {
	crt := s.top
	for {
		if crt == nil {
			break
		}
		fmt.Printf("%#v\n", crt)
		crt = crt.next
	}
}

func main() {
	s := &stack{
		top: &node{
			val: 1,
			next: &node{
				val: 2,
				next: &node{
					val: 3,
				},
			},
		},
	}
	s.pop()
	s.traverse()
	fmt.Println("----") // 2 3
	s.pop()
	s.traverse() // 3
	fmt.Println("----")
	s.pop()
	s.traverse() // nil

	s.push(1)
	s.traverse() // 1
	fmt.Println("----")
	s.push(2)
	s.push(3)
	s.traverse() // 3 2 1

	fmt.Println("----")
	fmt.Printf("%#v\n", s.peek()) // 3

	s2 := &stack{}
	fmt.Println(s2.isEmpty()) // true
}
