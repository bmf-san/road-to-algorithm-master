package main

import "fmt"

type stack struct {
	nodes []*node
}

type node struct {
	val int
}

// Remove data from the top of the stack.
func (s *stack) pop() {
	// LIFO
	s.nodes = s.nodes[1:]
}

// Add the item to the top of the stack.
func (s *stack) push(item int) {
	// LIFO
	s.nodes = append(
		[]*node{
			&node{
				val: item,
			},
		},
		s.nodes...,
	)
}

// Returns the top item from the stack.
func (s *stack) peek() *node {
	return s.nodes[0]
}

// Returns true if the stack is empty.
func (s *stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *stack) traverse() {
	for _, n := range s.nodes {
		fmt.Printf("%#v\n", n)
	}
}

func main() {
	s := &stack{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
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
